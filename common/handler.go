/**
 * @Author: Resynz
 * @Date: 2022/6/15 16:50
 */
package common

import (
	"cache-service/code"
	"github.com/gin-gonic/gin"
	"net/http"
)

type responseObj struct {
	Code code.ResponseCode `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data,omitempty"`
}


func HandleResponse(ctx *gin.Context,c code.ResponseCode,d interface{},msg ...string) {
	m := code.GetCodeMsg(c)
	if len(msg) > 0 {
		m = msg[0]
	}
	ro:=&responseObj{
		Code:    c,
		Message: m,
		Data:    d,
	}
	ctx.JSON(http.StatusOK, ro)
}
