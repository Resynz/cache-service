/**
 * @Author: Resynz
 * @Date: 2022/6/15 15:29
 */
package conf

type Config struct {
	Port int `json:"port"`
	Redis *RedisConfig `json:"redis"`
}

type RedisConfig struct {
	Host string `json:"host"`
	Port int `json:"port"`
	Password string `json:"password"`
	DB int `json:"db"`
}

type Fn struct {
	Method string `form:"method" binding:"required" json:"method"`
	Uri string `form:"uri" binding:"required" json:"uri"`
	Param interface{} `form:"param" binding:"" json:"param,omitempty"`
}

type Cfn struct {
	Name string `form:"name" binding:"required" json:"name"`
	Expire int64 `form:"expire" binding:"" json:"expire"`
	Fn *Fn `form:"fn" binding:"required" json:"fn"`
}
