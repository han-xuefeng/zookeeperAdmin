package starter

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"han-xuefeng/zookeeperAdmin/app/dao"
)

var ginEngine *gin.Engine

type GinStarter struct {
	BaseStarter
}

func NewGin() *gin.Engine{
	return ginEngine
}

func (ginStart *GinStarter)Init(ctx StarterContext){
	ginEngine = initGinEngine()
}

func (ginStart *GinStarter)Setup(ctx StarterContext){
	// 加载路由
	fmt.Println("注册路由")
	ginEngine.GET("/", func(context *gin.Context) {
		db := NewDbGorm()
		user := &dao.User{}
		db.First(user)
		fmt.Println(user)
	})
}

// Start 启动
func (ginStart *GinStarter) Start(ctx StarterContext){
	for _, r := range ginEngine.Routes() {
		fmt.Println(r.Path)
	}
	ginEngine.Run(":8080")
}

func initGinEngine() *gin.Engine{
	// todo 注册中间件 日志等组件
	return gin.New()
}