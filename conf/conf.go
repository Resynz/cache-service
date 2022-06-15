/**
 * @Author: Resynz
 * @Date: 2022/6/15 15:45
 */
package conf

import (
	"cache-service/lib/fn"
	"cache-service/lib/redis"
	"encoding/json"
	"os"
	"sync"
	"time"
)

var (
	Conf Config
	CfnMaps *CfnMap
)

func InitConf () error {
	c,err:=os.ReadFile("./configs/config.json")
	if err!=nil{
		return err
	}
	if err = json.Unmarshal(c,&Conf);err!=nil{
		return err
	}
	CfnMaps = &CfnMap{
		Mutex: &sync.Mutex{},
		maps:  make(map[string]*Cfn),
	}
	return nil
}

func InitCnfMaps() error {
	c,err:=os.ReadFile("./data/cfn.json")
	if err!=nil{
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	var maps map[string]*Cfn
	if err = json.Unmarshal(c,&maps);err!=nil{
		return err
	}

	for k,v:=range maps {
		value,err:=fn.CallFn(v.Fn)
		if err!=nil{
			return err
		}

		if err = redis.RedisClient.Set(v.Name,value,time.Second * time.Duration(v.Expire)).Err();err!=nil{
			return err
		}
		CfnMaps.Set(k,v)
	}
	return nil
}

func StoreCfnMaps() error {

}