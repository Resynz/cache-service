/**
 * @Author: Resynz
 * @Date: 2022/6/15 16:22
 */
package server

import (
	"cache-service/conf"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"log"
)
func StartServer() {
	gin.SetMode(gin.ReleaseMode)
	app:=gin.New()
	app.MaxMultipartMemory = 8 << 20 // 8MB
	_ = app.SetTrustedProxies(nil)
	app.Use(gzip.Gzip(gzip.DefaultCompression))
	// 加入中间件
	corConf := cors.DefaultConfig()
	corConf.AllowAllOrigins = true
	corConf.AllowHeaders = []string{
		"Content-Type",
		"token",
	}
	app.TrustedPlatform = gin.PlatformGoogleAppEngine
	app.Use(cors.New(corConf))
	app.Use(gin.Recovery())
	RegisterApiRoute(app.Group("/api"))
	go func() {
		log.Printf("start success at port %d\n",conf.Conf.Port)
	}()
	if err := app.Run(fmt.Sprintf(":%d", conf.Conf.Port)); err != nil {
		log.Fatalf("start server failed! error:%s\n", err.Error())
	}
}
