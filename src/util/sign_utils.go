package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"lingye-gin/src/config"
	"net/url"
	"sort"
	"strconv"
	"time"
)

// MD5 方法
func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

// 获取当前时间戳
func GetTimeUnix() int64 {
	return time.Now().Unix()
}

// 生成签名
func CreateSign(params url.Values) string {
	var key []string
	var str = ""
	for k := range params {
		if k != "sn" {
			key = append(key, k)
		}
	}
	sort.Strings(key)
	for i := 0; i < len(key); i++ {
		if i == 0 {
			str = fmt.Sprintf("%v=%v", key[i], params.Get(key[i]))
		} else {
			str = str + fmt.Sprintf("&%v=%v", key[i], params.Get(key[i]))
		}
	}
	// 自定义签名算法
	sign := MD5(MD5(str) + MD5(config.AppProps.App.Name+config.AppProps.Jwt.Secret))
	return sign
}

// 验证签名
func VerifySign(c *gin.Context) {
	var method = c.Request.Method
	var ts int64
	var sn string
	var req url.Values

	if method == "GET" {
		req = c.Request.URL.Query()
		sn = c.Query("sn")
		ts, _ = strconv.ParseInt(c.Query("ts"), 10, 64)

	} else if method == "POST" {
		_ = c.Request.ParseForm()
		req = c.Request.PostForm
		sn = c.PostForm("sn")
		ts, _ = strconv.ParseInt(c.PostForm("ts"), 10, 64)
	} else {
		RErrorJson(c, "Illegal requests", "")
		return
	}

	exp, _ := strconv.ParseInt(string(rune(config.AppProps.Jwt.Expiry)), 10, 64)

	// 验证过期时间
	if ts > GetTimeUnix() || GetTimeUnix()-ts >= exp {
		RErrorJson(c, "Ts Error", "")
		return
	}

	// 验证签名
	if sn == "" || sn != CreateSign(req) {
		RErrorJson(c, "Sn Error", "")
		return
	}
}
