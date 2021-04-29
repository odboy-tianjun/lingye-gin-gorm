package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"lingye-gin/src/config"
	"lingye-gin/src/modules/system/entity"
	"net/http"
	"strings"
	"time"
)

// Jwt引擎
type JwtEngine struct {
}

// 认证请求头
const HeadAuthorization = "Authorization"
const HeadLingYe = "LingYe"

// Payload 载荷
// 继承jwt提供给载荷，扩展自己所需字段
type JwtClaims struct {
	jwt.StandardClaims
	UserId   uint64   `json:"id"`        // 用户ID
	UserName string   `json:"username"`  // 用户名
	NickName string   `json:"nickname"`  // 用户昵称
	Email    string   `json:"email"`     // 邮箱账号
	UserRole []string `json:"userRoles"` // 用户角色编码合集
}

// 生成令牌
func (JwtEngine) CreateJwtToken(user entity.User, timeout int) (string, bool) {
	// Jwt Secret 私钥
	jwtSecret := config.AppProps.Jwt.Secret
	// 过期时间
	expiredAt := time.Now().Add(time.Hour * time.Duration(timeout)).Unix()
	// 设置载荷
	claims := JwtClaims{}
	claims.UserId = user.ID
	claims.UserName = user.UserName
	claims.NickName = user.NickName
	claims.Email = user.Email
	claims.ExpiresAt = expiredAt
	// claims.Issuer = "ginv" // 非必须，也可以填充用户名
	// 生成令牌 采用HMAC SHA256算法加密
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	// 令牌签名
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", false
	}
	return tokenString, true
}

// 验证令牌
func (JwtEngine) validateJwt(tokenString string) (*JwtClaims, bool) {
	// Jwt Secret 私钥
	jwtSecret := config.AppProps.Jwt.Secret
	// 解析令牌字符串
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		config.Logger.Println(err)
		return nil, false
	}
	// 获取载荷
	claims, ok := token.Claims.(*JwtClaims)
	if ok && token.Valid {
		return claims, true
	}
	return nil, false
}

// 更新Token
func (engine JwtEngine) RefreshToken(tokenString string) (string, bool) {
	// Jwt Secret 私钥
	jwtSecret := config.AppProps.Jwt.Secret
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	// 解析令牌字符串
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		config.Logger.Println(err)
		return "", false
	}
	// 获取载荷
	var claims *JwtClaims
	var ok bool
	if claims, ok = token.Claims.(*JwtClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		// 过期时间
		expiredAt := time.Now().Add(1 * time.Hour).Unix()
		claims.StandardClaims.ExpiresAt = expiredAt
		// 拷贝载体
		currentUserInfo := entity.User{}
		currentUserInfo.ID = claims.UserId
		currentUserInfo.UserName = claims.UserName
		currentUserInfo.NickName = claims.NickName
		currentUserInfo.Email = claims.Email
		return engine.CreateJwtToken(currentUserInfo, config.AppProps.Jwt.Expiry)
	}
	return "", false
}

// 初始化
func (engine JwtEngine) Init() {
	config.Logger.Info("JwtEngine Init...")
	config.JwtHandle = func(c *gin.Context) {
		// 获取请求头中的Authorization
		authorization := c.Request.Header.Get(HeadAuthorization)
		if strings.Compare(authorization, "") == 0 {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "Authorization Header Is None, No Permission",
			})
			c.Abort()
			return
		}
		// 拆分Authorization字段获取token字符串
		strSlice := strings.SplitN(authorization, " ", 2)
		if len(strSlice) != 2 || strings.Compare(HeadLingYe, strSlice[0]) != 0 {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "Authorization Header Format Is Error, No Permission",
			})
			c.Abort()
			return
		}
		// 验证token字符串
		claim, ok := engine.validateJwt(strSlice[1])
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusBadRequest,
				"msg":  "Token Validate Error",
			})
			c.Abort()
			return
		}
		// 过期判断
		if time.Now().Unix() > claim.ExpiresAt {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusBadRequest,
				"msg":  "Token Is Expired",
			})
			c.Abort()
			return
		}
		// 设置用户名
		c.Set("username", claim.UserName)
		c.Next()
	}
	config.Logger.Info("JwtEngine Ok...")
}
