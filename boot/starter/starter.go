package starter

type StarterContext map[string]interface{}

// BaseStarter 基础启动器
type BaseStarter struct {
}

// Starter 启动器的接口 所有需要启动的服务  需要实现这个接口 由启动器管理器统一管理
type Starter interface {
	Init(StarterContext)
	Setup(StarterContext)
	Start(StarterContext)
	StartBlocking() bool
	Stop(StarterContext)
}

// 基础启动器实现基础的方法  后续的启动器不需要每个方法都实现一遍

func (s *BaseStarter) Init(ctx StarterContext)  {}
func (s *BaseStarter) Setup(ctx StarterContext) {}
func (s *BaseStarter) Start(ctx StarterContext) {}
func (s *BaseStarter) Stop(ctx StarterContext)  {}
func (s *BaseStarter) StartBlocking() bool      { return false }

// StarterManage 启动器管理器 管理所有的启动器
type StarterManage struct {
	starters []Starter
}

func NewStarterManage() *StarterManage {
	// 因为切片必须要先初始化
	starters := make([]Starter,0)
	return &StarterManage{
		starters: starters,
	}
}

// Register 启动器管理器的注册方法 注册所有的启动器
func (m *StarterManage) Register(s Starter){
	m.starters = append(m.starters, s)
}

// GetAllStarters 启动器管理器  启动的方法
func (m *StarterManage) GetAllStarters()[]Starter {
	return m.starters
}
