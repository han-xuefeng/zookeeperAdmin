package infra

import (
	log "github.com/sirupsen/logrus"
	"github.com/tietang/props/kvs"
	"reflect"
	"sort"
)

const (
	KeyProps = "_conf"
)

const (
	SystemGroup         PriorityGroup = 30
	BasicResourcesGroup PriorityGroup = 20
	AppGroup            PriorityGroup = 10

	INT_MAX          = int(^uint(0) >> 1)
	DEFAULT_PRIORITY = 10000
)

// StarterContext 基础资源上下结构体
type StarterContext map[string]interface{}

// Props 这里加载配置文件
func (s StarterContext) Props() kvs.ConfigSource {
	p := s[KeyProps]
	if p == nil {
		panic("配置还没被初始化")
	}
	return p.(kvs.ConfigSource)
}

type PriorityGroup int

// Starter 基础资源启动器接口
type Starter interface {
	// Init 1.系统启动，初始化一些基础资源
	Init(StarterContext)
	// Setup 2. 系统基础资源的安装
	Setup(StarterContext)
	// Start 3. 启动基础资源
	Start(StarterContext)
	// StartBlocking 启动器是否可阻塞
	StartBlocking() bool
	// Stop 4. 资源停止和销毁
	Stop(StarterContext)

	PriorityGroup() PriorityGroup
	Priority() int
}

var _ Starter = new(BaseStarter)

//基础空启动器实现,为了方便资源启动器的代码实现
type BaseStarter struct {
}



func (b *BaseStarter) Init(ctx StarterContext)  {}
func (b *BaseStarter) Setup(ctx StarterContext) {}
func (b *BaseStarter) Start(ctx StarterContext) {}
func (b *BaseStarter) StartBlocking() bool      { return false }
func (b *BaseStarter) Stop(ctx StarterContext)  {}

func (b *BaseStarter) PriorityGroup() PriorityGroup { return BasicResourcesGroup }
func (b *BaseStarter) Priority() int                { return DEFAULT_PRIORITY }

type Starters []Starter

func (s Starters) Len() int      { return len(s) }
func (s Starters) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Starters) Less(i, j int) bool {
	return s[i].PriorityGroup() > s[j].PriorityGroup() && s[i].Priority() > s[j].Priority()
}

// 全局只有一个
var StarterRegister *starterRegister = new(starterRegister)
// 启动器注册器
type starterRegister struct {
	nonBlockingStarters []Starter
	blockingStarters    []Starter
}

func SortStarters() {
	sort.Sort(Starters(StarterRegister.AllStarters()))
}

//获取所有注册的starter
func GetStarters() []Starter {
	return StarterRegister.AllStarters()
}

// 启动器注册器   把阻塞的启动器放到最后启动
func (r *starterRegister) Register(starter Starter) {
	if starter.StartBlocking() {
		r.blockingStarters = append(r.blockingStarters, starter)
	} else {
		r.nonBlockingStarters = append(r.nonBlockingStarters, starter)
	}
	typ := reflect.TypeOf(starter)
	log.Infof("Register starter: %s", typ.String())
}

func (r *starterRegister) AllStarters() []Starter {
	starters := make([]Starter, 0)
	starters = append(starters, r.nonBlockingStarters...)
	starters = append(starters, r.blockingStarters...)
	return starters
}

func Register(s Starter) {
	StarterRegister.Register(s)
}