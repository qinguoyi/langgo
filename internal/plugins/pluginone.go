package plugins

import (
	"github.com/corex-io/codec"
)

/*当实现接口时，如果方法的接收类型为指针和普通类型混用，则返回与注册的类型需要为指针*/

type PluginOne struct {
	Id string
}

// Name 接口名称
func (p PluginOne) Name() string {
	return "pluginOne"
}

// Load 数据转换为插件信息
func (p PluginOne) Load(i interface{}) (Plugin, error) {
	return &p, codec.Format(&p, i)
}

// Store 存储数据
func (p *PluginOne) Store(id string) {
	p.Id = id
}

func (p PluginOne) GetInfo() string {
	return p.Id
}

// init注册插件信息
func init() {
	p := PluginOne{}
	RegistPlugin(&p)
}
