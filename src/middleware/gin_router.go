package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lingye-gin/src"
	"lingye-gin/src/config"
	"strings"
)

type GinRouter struct {
}

func (v GinRouter) Init(r *gin.Engine) {
	config.Logger.Info("GinRouter Init")
	// 全局异常拦截器
	r.Use(GinPanic)

	// 一般路由映射关系
	groupMap := make(map[string]*gin.RouterGroup)
	// api路由映射关系
	apiUrls := make([]main.RequestApi, 0)
	// api组名称
	apiGroupName := "api"

	for _, normalRa := range main.Urls {
		// 判断是否在某一组下
		if strings.Compare(normalRa.GroupName, "") == 0 {
			// 批量处理
			if handle(normalRa, r) {
				continue
			}
		} else {
			// 分组名称
			groupName := fmt.Sprintf("/%s", normalRa.GroupName)

			// api组, 基于jwt验证
			if strings.Compare(normalRa.GroupName, apiGroupName) == 0 {
				apiUrls = append(apiUrls, normalRa)
				continue
			}

			// 分组不存在
			if groupMap[normalRa.GroupName] == nil {
				if normalRa.GroupHandleFunction == nil {
					// 不存在分组过滤器
					groupMap[normalRa.GroupName] = r.Group(groupName)
				} else {
					groupMap[normalRa.GroupName] = r.Group(groupName, normalRa.GroupHandleFunction)
				}
			}
			// 批量处理
			if handleGroup(normalRa, groupMap[normalRa.GroupName]) {
				continue
			}
		}
	}

	for _, apiRa := range apiUrls {
		jwtR := r.Group(apiGroupName)
		jwtR.Use(config.JwtHandle)
		// 批量处理
		if handleGroup(apiRa, jwtR) {
			continue
		}
	}

	config.Logger.Info("GinRouter Ok")
}

func handle(request main.RequestApi, engine *gin.Engine) bool {
	// get
	if strings.Compare(request.Mode, "get") == 0 {
		engine.GET(request.RelativePath, request.HandleFunction)
		return true
	}
	// post
	if strings.Compare(request.Mode, "post") == 0 {
		engine.POST(request.RelativePath, request.HandleFunction)
		return true
	}
	// delete
	if strings.Compare(request.Mode, "delete") == 0 {
		engine.DELETE(request.RelativePath, request.HandleFunction)
		return true
	}
	// put
	if strings.Compare(request.Mode, "put") == 0 {
		engine.PUT(request.RelativePath, request.HandleFunction)
		return true
	}
	return false
}

func handleGroup(request main.RequestApi, group *gin.RouterGroup) bool {
	// get
	if strings.Compare(request.Mode, "get") == 0 {
		group.GET(request.RelativePath, request.HandleFunction)
		return true
	}
	// post
	if strings.Compare(request.Mode, "post") == 0 {
		group.POST(request.RelativePath, request.HandleFunction)
		return true
	}
	// delete
	if strings.Compare(request.Mode, "delete") == 0 {
		group.DELETE(request.RelativePath, request.HandleFunction)
		return true
	}
	// put
	if strings.Compare(request.Mode, "put") == 0 {
		group.PUT(request.RelativePath, request.HandleFunction)
		return true
	}
	return false
}
