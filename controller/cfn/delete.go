/**
 * @Author: Resynz
 * @Date: 2022/6/16 09:44
 */
package cfn

import (
	"cache-service/code"
	"cache-service/common"
	"cache-service/conf"
	"cache-service/lib/redis"
	"github.com/gin-gonic/gin"
)

func Delete(ctx *gin.Context) {
	name := ctx.Param("name")
	cfn := conf.CfnMaps.Get(name)
	if cfn != nil {
		err := redis.RedisClient.Del(cfn.Name).Err()
		if err != nil {
			common.HandleResponse(ctx, code.BadRequest, nil, err.Error())
			return
		}
		conf.CfnMaps.Del(name)
	}
	common.HandleResponse(ctx, code.SuccessCode, nil)
}
