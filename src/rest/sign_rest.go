package rest

import (
	"github.com/gin-gonic/gin"
	"lingye-gin/src/util"
	"net/url"
	"strconv"
)

func DescribeSign(c *gin.Context) {
	ts := strconv.FormatInt(util.GetTimeUnix(), 10)
	res := map[string]interface{}{}
	params := url.Values{
		"username": []string{"lingye"},
		"age":      []string{"10"},
		"ts":       []string{ts},
	}
	res["sn"] = util.CreateSign(params)
	res["ts"] = ts
	util.RetJson("200", "", res, c)
}
