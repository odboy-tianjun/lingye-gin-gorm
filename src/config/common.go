package config

import "github.com/gin-gonic/gin"

// 应用配置
var AppProps ApplicationProperties

// jwt拦截器
var JwtHandle gin.HandlerFunc
