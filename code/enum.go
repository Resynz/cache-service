/**
 * @Author: Resynz
 * @Date: 2022/6/15 16:54
 */
package code

type ResponseCode int
const (
	SuccessCode ResponseCode = 200
	BadRequest ResponseCode = 1000
	InvalidParams ResponseCode = 1001
	InvalidRequest ResponseCode = 1002
)

var ResponseCodeMap = map[ResponseCode]string {
	SuccessCode: "请求成功",
	BadRequest: "系统错误",
	InvalidParams: "请求参数无效",
	InvalidRequest: "无效的请求",
}
func GetCodeMsg(code ResponseCode) string {
	h, ok := ResponseCodeMap[code]
	if !ok {
		return "未知错误"
	}
	return h
}
