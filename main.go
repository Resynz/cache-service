/**
 * @Author: Resynz
 * @Date: 2022/6/15 15:26
 */
package main

import (
	"cache-service/conf"
	"cache-service/lib/redis"
	"cache-service/queue"
	"cache-service/server"
	"log"
)

func main() {
	var err error
	log.Println("initializing config ...")
	if err = conf.InitConf();err!=nil{
		log.Fatalf("init config failed! error:%s\n",err.Error())
	}
	log.Println("initializing redis ...")
	if err = redis.InitRedis();err!=nil{
		log.Fatalf("init redis failed! error:%s\n",err.Error())
	}
	log.Println("initializing cfn ...")
	if err = conf.InitCnfMaps();err!=nil{
		log.Fatalf("init cfn failed! error:%s\n",err.Error())
	}
	log.Println("starting service exit listener...")
	queue.SetSignalHandler()
	log.Println("starting service ...")
	server.StartServer()
}
