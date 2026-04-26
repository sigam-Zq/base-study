package main

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
)

func main() {
	// p, err := plugin.Open("./plugin-e/target/plugin-e.dll")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// pluP, err := p.Lookup("PluginP")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// pluP.(func(string))("Go")

	// 根据操作系统选择插件文件扩展名
	var pluginFile string
	if runtime.GOOS == "windows" {
		pluginFile = "plugin-e.dll"
	} else {
		pluginFile = "plugin-e.so"
	}

	// 构建插件路径
	pluginPath := filepath.Join(".", "plugin-e", "target", pluginFile)

	// 加载插件函数
	pluginFunc, err := LoadPlugin(pluginPath, "PluginP")
	if err != nil {
		log.Fatalf("Failed to load plugin: %v", err)
	}

	// 调用插件函数
	fmt.Println("Calling plugin function...")
	pluginFunc("Go")
	fmt.Println("Plugin call completed successfully!")
}
