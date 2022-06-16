/**
 * @Author: Resynz
 * @Date: 2022/6/16 09:14
 */
package structs

type Fn struct {
	Method string      `form:"method" binding:"required" json:"method"`
	Uri    string      `form:"uri" binding:"required" json:"uri"`
	Param  interface{} `form:"param" binding:"" json:"param,omitempty"`
}

type Cfn struct {
	Name   string `form:"name" binding:"required" json:"name"`
	Expire int64  `form:"expire" binding:"" json:"expire"`
	Fn     *Fn    `form:"fn" binding:"required" json:"fn"`
}
