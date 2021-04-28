package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"lingye-gin/src/config"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"time"
)

// 打印
func Print(i interface{}) {
	fmt.Println("---")
	fmt.Println(i)
	fmt.Println("---")
}

// 返回JSON
func RetJson(code, msg string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
	c.Abort()
}

// 获取当前时间戳
func GetTimeUnix() int64 {
	return time.Now().Unix()
}

// MD5 方法
func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

// 修正分页参数
func FixPage(page uint, pageSize uint) (p uint, ps uint) {
	if page == 0 {
		p = 1
	}
	if pageSize == 0 {
		ps = 15
	}
	return p, ps
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
		RetJson("500", "Illegal requests", "", c)
		return
	}

	exp, _ := strconv.ParseInt(string(rune(config.AppProps.Jwt.Expiry)), 10, 64)

	// 验证过期时间
	if ts > GetTimeUnix() || GetTimeUnix()-ts >= exp {
		RetJson("500", "Ts Error", "", c)
		return
	}

	// 验证签名
	if sn == "" || sn != CreateSign(req) {
		RetJson("500", "Sn Error", "", c)
		return
	}
}

// source  有数据的结构体
// target  空的结构体
func StructCopy(source interface{}, target interface{}) {
	sourceVal := reflect.ValueOf(source).Elem() // 获取reflect.Type类型
	targetVal := reflect.ValueOf(target).Elem() // 获取reflect.Type类型
	vTypeOfT := sourceVal.Type()
	for i := 0; i < sourceVal.NumField(); i++ {
		// 在要空的结构体中查询有数据结构体中相同属性的字段，有则修改其值
		name := vTypeOfT.Field(i).Name
		if ok := targetVal.FieldByName(name).IsValid(); ok {
			targetVal.FieldByName(name).Set(reflect.ValueOf(sourceVal.Field(i).Interface()))
		}
	}
}
