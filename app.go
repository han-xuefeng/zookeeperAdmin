package zookeeperAdmin

import (
	_ "han-xuefeng/zookeeperAdmin/apis/web"
	"han-xuefeng/zookeeperAdmin/infra"
	"han-xuefeng/zookeeperAdmin/infra/base"
)

func init() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.GinServerStarter{})
	infra.Register(&base.DbxGorm{})
	infra.Register(&infra.WebApiStarter{})
}
