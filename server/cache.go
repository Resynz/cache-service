/**
 * @Author: Resynz
 * @Date: 2022/6/16 11:17
 */
package server

import (
	"cache-service/controller/cache"
	"github.com/gin-gonic/gin"
)

func RegisterCacheRoute(route *gin.RouterGroup) {
	route.GET("/:name", cache.Get)
}
