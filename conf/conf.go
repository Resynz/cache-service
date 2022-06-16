/**
 * @Author: Resynz
 * @Date: 2022/6/15 15:45
 */
package conf

import (
	"cache-service/structs"
	"encoding/json"
	"os"
	"sync"
)

var (
	Conf    Config
	CfnMaps *CfnMap
)

const DataPath = "./data/cfn.json"
const ConfigPath = "./configs/config.json"

func InitConf() error {
	c, err := os.ReadFile(ConfigPath)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(c, &Conf); err != nil {
		return err
	}
	CfnMaps = &CfnMap{
		Mutex: &sync.Mutex{},
		maps:  make(map[string]*structs.Cfn),
	}
	return nil
}

func InitCnfMaps() error {
	c, err := os.ReadFile(DataPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	var maps []*structs.Cfn
	if err = json.Unmarshal(c, &maps); err != nil {
		return err
	}

	for _, v := range maps {
		CfnMaps.Set(v.Name, v)
	}
	return nil
}

func StoreCfnMaps() error {
	mps := make([]*structs.Cfn, 0)
	for _, v := range CfnMaps.maps {
		mps = append(mps, v)
	}
	f, err := os.OpenFile(DataPath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	ec := json.NewEncoder(f)
	if err = ec.Encode(&mps); err != nil {
		return err
	}
	return nil
}
