/**
 * @Author: Resynz
 * @Date: 2022/6/15 18:02
 */
package conf

import (
	"cache-service/structs"
	"sync"
)

type CfnMap struct {
	*sync.Mutex
	maps map[string]*structs.Cfn
}

func (s *CfnMap) Get(name string) *structs.Cfn {
	v, ok := s.maps[name]
	if !ok {
		return nil
	}
	return v
}

func (s *CfnMap) Set(name string, val *structs.Cfn) {
	s.Lock()
	defer s.Unlock()
	s.maps[name] = val
}

func (s *CfnMap) Del(name string) {
	s.Lock()
	defer s.Unlock()
	delete(s.maps, name)
}
