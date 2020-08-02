package main

import (
	"flag"
	"fmt"
	"github.com/cshengqun/wechat_public_account/global"
	"github.com/cshengqun/wechat_public_account/processor"
	"github.com/gin-gonic/gin"
)

var (
	path = flag.String("path", "./conf/live.toml", "config file path")
)

func initBaseUrl(engine *gin.Engine)  {
	engine.GET("/wechat", processor.ValidateWechatReq)
}

func initMessageUrl(engine *gin.Engine) {
	engine.POST("/wechat", processor.DealWithMsg)
}

func main()  {
	err := global.Init(*path)
	if err != nil {
		panic(err)
	}
	engine := gin.Default()
	initBaseUrl(engine)
	initMessageUrl(engine)
	engine.Run(fmt.Sprintf("%s:%d", global.GetConfig().Host.Ip, global.GetConfig().Host.Port))
}
