//go:build windows

package main

import (
	"log"
	"syscall"
	"unsafe"
)

// 条件编译：Windows实现

// WindowsPluginLoader Windows DLL加载器
type WindowsPluginLoader struct {
	dllHandle syscall.Handle
}

// NewPluginLoader 创建Windows插件加载器
func NewPluginLoader() PluginLoader {
	return &WindowsPluginLoader{}
}

// Load 加载DLL并获取函数地址
// func (w *WindowsPluginLoader) Load(pluginPath string, functionName string) (PluginFunction, error) {

// 	dll, err := syscall.LoadLibrary(pluginPath)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to load DLL: %w", err)
// 	}

// 	proc, err := syscall.GetProcAddress(dll, functionName)
// 	if err != nil {
// 		syscall.FreeLibrary(dll)
// 		return nil, fmt.Errorf("failed to find function %s: %w", functionName, err)
// 	}

// 	w.dllHandle = dll

// 	// 返回闭包函数
// 	return func(name string) {
// 		// 转换字符串为C字符串
// 		cName, err := syscall.BytePtrFromString(name)
// 		if err != nil {
// 			log.Printf("Failed to convert string: %v", err)
// 			return
// 		}

// 		// 调用DLL函数
// 		_, _, _ = syscall.Syscall(uintptr(proc), 1, uintptr(unsafe.Pointer(cName)), 0, 0)
// 	}, nil
// }

// Load 加载DLL并获取函数地址
func (w *WindowsPluginLoader) Load(pluginPath string, functionName string) (PluginFunction, error) {

	dll := syscall.NewLazyDLL(pluginPath)
	proc := dll.NewProc(functionName)

	// 返回闭包函数
	return func(name string) {
		cName, err := syscall.BytePtrFromString(name)
		_, _, err = proc.Call(uintptr(unsafe.Pointer(cName)))
		if err != syscall.Errno(0) {
			log.Println("call error:", err)
		}

	}, nil
}

// Close 释放DLL
func (w *WindowsPluginLoader) Close() error {
	if w.dllHandle != 0 {
		return syscall.FreeLibrary(w.dllHandle)
	}
	return nil
}
