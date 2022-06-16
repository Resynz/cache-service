/**
 * @Author: Resynz
 * @Date: 2022/6/16 10:13
 */
package cache

import (
	"cache-service/code"
	"cache-service/common"
	"cache-service/conf"
	"cache-service/lib/fn"
	"cache-service/lib/redis"
	"github.com/gin-gonic/gin"
	"time"
)

func Get(ctx *gin.Context) {
	name := ctx.Param("name")
	cfn := conf.CfnMaps.Get(name)
	if cfn == nil {
		common.HandleResponse(ctx, code.InvalidParams, nil)
		return
	}
	val := redis.RedisClient.Get(name).Val()
	if val == "" {
		res, err := fn.CallFn(cfn.Fn)
		if err != nil {
			common.HandleResponse(ctx, code.BadRequest, nil, err.Error())
			return
		}
		val = string(res)
		go func() {
			redis.RedisClient.Set(cfn.Name, res, time.Second*time.Duration(cfn.Expire))
		}()
	}
	common.HandleResponse(ctx, code.SuccessCode, val)
}
