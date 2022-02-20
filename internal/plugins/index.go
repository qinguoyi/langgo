package plugins

// Plugin 插件接口声明
type Plugin interface {
	Name() string
	Load(interface{}) (Plugin, error)
	Store(string)
	GetInfo() string
}

// Plugins 插件注册集合
var Plugins = make(map[string]Plugin)

// RegistPlugin 插件注册函数
func RegistPlugin(plugin Plugin) {
	Plugins[plugin.Name()] = plugin
}
