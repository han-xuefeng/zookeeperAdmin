package boot

import "han-xuefeng/zookeeperAdmin/boot/starter"

// Application 应用程序管理器
type Application struct {
	starterContext starter.StarterContext
	starterManage  *starter.StarterManage
}


func NewApplication() *Application {
	app := &Application{starterContext: starter.StarterContext{}, starterManage: starter.NewStarterManage()}
	app.registerBaseStarter()
	return app
}

// 注册默认的starter 系统自己需要用的 后续需要的 可以使用Register
func (app *Application) registerBaseStarter(){
	// todo:
	// 注册gin
	app.Register(&starter.GinStarter{})
}

func (app *Application)Register(s starter.Starter){
	app.starterManage.Register(s)
}

// Start 启动
func (app *Application) Start(){
	app.init()
	app.setup()
	app.start()
}

func (app *Application)init(){
	starters := app.starterManage.GetAllStarters()
	for _, s := range starters {
		s.Init(app.starterContext)
	}
}

func (app *Application)setup(){
	starters := app.starterManage.GetAllStarters()
	for _, s := range starters {
		s.Setup(app.starterContext)
	}
}

func (app *Application)start(){
	starters := app.starterManage.GetAllStarters()
	for _, s := range starters {
		if s.StartBlocking() {
			go s.Start(app.starterContext)
		} else {
			s.Start(app.starterContext)
		}
	}
}