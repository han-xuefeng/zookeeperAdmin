package base

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"han-xuefeng/zookeeperAdmin/infra"
	"os"
	"time"
)

var ginApplication *gin.Engine

type GinServerStarter struct {
	infra.BaseStarter
}

func Gin() *gin.Engine {
	Check(ginApplication)
	return ginApplication
}

func (g *GinServerStarter) Init(cxt infra.StarterContext) {
	ginApplication = initGin()
}

func (g *GinServerStarter) Setup(cxt infra.StarterContext) {
	ginApplication.GET("/", func(context *gin.Context) {
		context.JSON(
			200,
			gin.H{
				"message": "pong",
			})
	})
}

func (g *GinServerStarter) Start(ctx infra.StarterContext) {
	//gin默认会把路由打印到控制台
	port := ctx.Props().GetDefault("app.server.port", "18080")
	Gin().Run(":"+port)
}

func (g *GinServerStarter) StartBlocking() bool {
	return true
}

func initGin() *gin.Engine {
	app := gin.New()

	logfile, err := os.Create("./logs/req_" + time.Now().Format("2006-01-02-15") + ".log")
	if err != nil {
		fmt.Println("Could not create log file")
	}
	app.Use(gin.LoggerWithWriter(logfile))
	app.Use(gin.Recovery()) // Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500
	return app
}
