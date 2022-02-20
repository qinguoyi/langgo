package plugins

import (
	"github.com/corex-io/codec"
)

type PluginTwo struct {
	Id string
}

// Name 接口名称
func (p PluginTwo) Name() string {
	return "pluginTwo"
}

// Load 数据转换为插件信息
func (p PluginTwo) Load(i interface{}) (Plugin, error) {
	return &p, codec.Format(&p, i)
}

// Store 存储数据
func (p *PluginTwo) Store(id string) {
	p.Id = id
}

// GetInfo 获取插件信息
func (p PluginTwo) GetInfo() string {
	return p.Id
}

// init 注册插件信息
func init() {
	p := PluginTwo{}
	RegistPlugin(&p)
}
