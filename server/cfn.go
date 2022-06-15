/**
 * @Author: Resynz
 * @Date: 2022/6/15 17:22
 */
package server

import (
	"cache-service/controller/cfn"
	"github.com/gin-gonic/gin"
)

func RegisterCfnRoute(route *gin.RouterGroup) {
	route.POST("/",cfn.Add)
}
