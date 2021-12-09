package boot

// Application 应用程序管理器
type Application struct {
	starterContext StarterContext
	starterManage *StarterManage
}


func NewApplication() *Application {
	app := &Application{starterContext:StarterContext{}, starterManage: NewStarterManage()}
	app.registerBaseStarter()
	return app
}

// 注册默认的starter 系统自己需要用的 后续需要的 可以使用Register
func (app *Application) registerBaseStarter(){
	// todo:
}

func (app *Application)Register(s Starter){
	app.starterManage.Register(s)
}

// Start 启动
func (app *Application) Start(){
	app.init()
}

func (app *Application)init(){
	starters := app.starterManage.GetAllStarters()
	for _, starter := range starters {
		starter.Init(app.starterContext)
	}
}

func (app *Application)setup(){
	starters := app.starterManage.GetAllStarters()
	for _, starter := range starters {
		starter.Setup(app.starterContext)
	}
}

func (app *Application)start(){
	starters := app.starterManage.GetAllStarters()
	for _, starter := range starters {
		if starter.StartBlocking() {
			go starter.Start(app.starterContext)
		} else {
			starter.Start(app.starterContext)
		}
	}
}