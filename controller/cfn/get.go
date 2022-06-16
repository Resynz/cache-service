/**
 * @Author: Resynz
 * @Date: 2022/6/16 10:13
 */
package cfn

import (
	"cache-service/code"
	"cache-service/common"
	"cache-service/conf"
	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	name := ctx.Param("name")
	cfn := conf.CfnMaps.Get(name)
	if cfn == nil {
		common.HandleResponse(ctx, code.InvalidParams, nil)
		return
	}
	common.HandleResponse(ctx, code.SuccessCode, cfn)
}
