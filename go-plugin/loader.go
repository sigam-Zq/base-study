package main

// PluginFunction 定义插件函数类型
type PluginFunction func(string)

// PluginLoader 插件加载器接口
type PluginLoader interface {
	Load(pluginPath string, functionName string) (PluginFunction, error)
	Close() error
}

// // LoadPlugin 跨平台插件加载函数
// func LoadPlugin(pluginPath string, functionName string) (PluginFunction, error) {
// 	var loader PluginLoader

// 	// 根据操作系统选择加载器
// 	if os.Getenv("OS") == "Windows_NT" || (len(string(os.PathListSeparator)) == 1 && os.PathListSeparator == ';') {
// 		loader = &WindowsPluginLoader{}
// 	} else {
// 		loader = &UnixPluginLoader{}
// 	}

// 	return loader.Load(pluginPath, functionName)
// }
func LoadPlugin(pluginPath string, functionName string) (PluginFunction, error) {
	loader := NewPluginLoader()
	return loader.Load(pluginPath, functionName)
}
