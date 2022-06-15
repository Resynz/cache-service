/**
 * @Author: Resynz
 * @Date: 2022/6/15 17:32
 */
package fn

import (
	"cache-service/conf"
	"fmt"
	"github.com/rosbit/gnet"
	"net/http"
)

func CallFn(fn *conf.Fn) ([]byte, error) {
	status,res,_,err:=gnet.Http(fn.Uri,gnet.M(fn.Method),gnet.Params(fn.Param))
	if err!=nil{
		return nil,err
	}
	if status != http.StatusOK {
		return nil,fmt.Errorf("invalid http status:%d resp:%s",status,string(res))
	}
	return res,err
}
