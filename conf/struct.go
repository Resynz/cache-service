/**
 * @Author: Resynz
 * @Date: 2022/6/15 15:29
 */
package conf

type Config struct {
	Port  int          `json:"port"`
	Redis *RedisConfig `json:"redis"`
}

type RedisConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}
