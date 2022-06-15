/**
 * @Author: Resynz
 * @Date: 2022/6/15 16:35
 */
package server

import (
	"cache-service/controller"
	"github.com/gin-gonic/gin"
)

func RegisterApiRoute(route *gin.RouterGroup) {
	route.GET("/ping",controller.Ping)

	RegisterCfnRoute(route.Group("/cfn"))
}
