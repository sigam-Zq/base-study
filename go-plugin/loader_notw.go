//go:build !windows

package main

import (
	"errors"
	"fmt"
	"plugin"
)

// 条件编译：Unix-like系统实现

// UnixPluginLoader Unix-like系统插件加载器
type UnixPluginLoader struct {
	plugin *plugin.Plugin
}

// NewPluginLoader 创建Unix插件加载器
func NewPluginLoader() PluginLoader {
	return &UnixPluginLoader{}
}

// Load 加载.so插件
func (u *UnixPluginLoader) Load(pluginPath string, functionName string) (PluginFunction, error) {
	p, err := plugin.Open(pluginPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open plugin: %w", err)
	}

	symbol, err := p.Lookup(functionName)
	if err != nil {
		return nil, fmt.Errorf("failed to lookup symbol %s: %w", functionName, err)
	}

	u.plugin = p

	// 类型断言
	fn, ok := symbol.(func(string))
	if !ok {
		return nil, errors.New("symbol is not of type func(string)")
	}

	return fn, nil
}

// Close 关闭插件
func (u *UnixPluginLoader) Close() error {
	// Go的plugin包没有显式的关闭方法
	// 依赖垃圾回收器来清理
	return nil
}
