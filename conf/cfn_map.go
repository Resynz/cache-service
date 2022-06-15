/**
 * @Author: Resynz
 * @Date: 2022/6/15 18:02
 */
package conf

import "sync"

type CfnMap struct {
	*sync.Mutex
	maps map[string]*Cfn
}

func (s *CfnMap) Get(name string) *Cfn {
	v,ok:=s.maps[name]
	if !ok {
		return nil
	}
	return v
}

func (s *CfnMap) Set(name string,val *Cfn)  {
	s.Lock()
	defer s.Unlock()
	s.maps[name] = val
}
