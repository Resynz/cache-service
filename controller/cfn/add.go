/**
 * @Author: Resynz
 * @Date: 2022/6/15 17:23
 */
package cfn

import (
	"cache-service/code"
	"cache-service/common"
	"cache-service/conf"
	"cache-service/lib/fn"
	"cache-service/lib/redis"
	"cache-service/structs"
	"github.com/gin-gonic/gin"
	"time"
)

func Add(ctx *gin.Context) {
	var form structs.Cfn
	if err := ctx.ShouldBind(&form); err != nil {
		common.HandleResponse(ctx, code.InvalidParams, nil, err.Error())
		return
	}
	value, err := fn.CallFn(form.Fn)
	if err != nil {
		common.HandleResponse(ctx, code.BadRequest, nil, err.Error())
		return
	}

	if err = redis.RedisClient.Set(form.Name, value, time.Second*time.Duration(form.Expire)).Err(); err != nil {
		common.HandleResponse(ctx, code.BadRequest, nil, err.Error())
		return
	}

	conf.CfnMaps.Set(form.Name, &form)
	common.HandleResponse(ctx, code.SuccessCode, nil)
}
