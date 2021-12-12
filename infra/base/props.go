package base

import (
	"github.com/tietang/props/kvs"
	"han-xuefeng/zookeeperAdmin/infra"
)

// props 全局变量
var props kvs.ConfigSource

func Props() kvs.ConfigSource {
	return props
}

type PropsStarter struct {
	infra.BaseStarter
}

func (p *PropsStarter)Init(ctx infra.StarterContext) {
	props = ctx.Props()
}