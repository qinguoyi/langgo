package main

import (
	"fmt"
	"langlangGo/internal/plugins"
)

type pluginTestData struct {
	Id string
}

func main() {
	// 从插件中获取插件一信息
	var plugin plugins.Plugin
	var ok bool
	plugin, ok = plugins.Plugins["pluginOne"]
	if ok {
		fmt.Println("pluginOne Id is %v", plugin.GetInfo())
	}

	// 设置插件数据信息
	plugin.Store("3")
	fmt.Println("pluginOne Id is %v", plugin.GetInfo())

	// 设置插件数据信息
	var pluginData interface{}
	pluginData = pluginTestData{
		Id: "1",
	}
	plugin, err := plugin.Load(pluginData)
	if err != nil {
		fmt.Println("set pluginOne fail")
	} else {
		fmt.Println("pluginOne Id is %v", plugin.GetInfo())
	}
}
