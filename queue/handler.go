/**
 * @Author: Resynz
 * @Date: 2020/8/4 10:36
 */
package queue

import (
	"cache-service/conf"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func SetSignalHandler() {
	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	go func() {
		defer func() {
			log.Println("程序退出.")
			os.Exit(0)
		}()
		s := <-sign
		log.Printf("接收到退出信号:%v\n", s)
		log.Println("正在保存信息...")

		if err := conf.StoreCfnMaps(); err != nil {
			log.Printf("保存失败！error:%v\n", err)
		}
	}()
}
