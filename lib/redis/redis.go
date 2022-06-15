/**
 * @Author: Resynz
 * @Date: 2022/6/15 15:51
 */
package redis

import (
	"cache-service/conf"
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

var RedisClient *redis.Client

func InitRedis() error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:               fmt.Sprintf("%s:%d",conf.Conf.Redis.Host,conf.Conf.Redis.Port),
		OnConnect: func(conn *redis.Conn) error {
			log.Println("redis connected.")
			return nil
		},
		Password:           conf.Conf.Redis.Password,
		DB:                 conf.Conf.Redis.DB,
	})
	_,err:=RedisClient.Ping().Result()
	return err
}
