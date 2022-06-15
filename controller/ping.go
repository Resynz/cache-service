/**
 * @Author: Resynz
 * @Date: 2022/6/15 16:44
 */
package controller

import (
	"cache-service/code"
	"cache-service/common"
	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	common.HandleResponse(ctx,code.SuccessCode,nil)
}
