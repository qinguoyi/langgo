package plugins

import (
	"fmt"
	"langgo/bootstrap"
)

// Plugin 插件接口
type Plugin interface {
	// Name 插件名称
	Name() string
	// New 初始化插件资源
	New() interface{}
	// Health 插件健康检查
	Health()
	// Close 释放插件资源
	Close()
}

// Plugins 插件注册集合
var Plugins = make(map[string]Plugin)

// RegisteredPlugin 插件注册
func RegisteredPlugin(plugin Plugin) {
	Plugins[plugin.Name()] = plugin
}

// NewPlugins 初始化插件资源
func NewPlugins() {
	for _, p := range Plugins {
		bootstrap.NewLogger().Logger.Info(fmt.Sprintf("%s Init ... ", p.Name()))
		p.New()
		bootstrap.NewLogger().Logger.Info(fmt.Sprintf("%s HealthCheck ... ", p.Name()))
		p.Health()
		bootstrap.NewLogger().Logger.Info(fmt.Sprintf("%s Success Init. ", p.Name()))
	}
}

// ClosePlugins 释放插件资源
func ClosePlugins() {
	for _, p := range Plugins {
		p.Close()
		bootstrap.NewLogger().Logger.Info(fmt.Sprintf("%s Success Close ... ", p.Name()))
	}
}
