#!/bin/bash
# 构建脚本

PLUGIN_DIR="plugin-e"
TARGET_DIR="$PLUGIN_DIR/target"

# 创建目标目录
mkdir -p "$TARGET_DIR"

# 检测操作系统
if [[ "$OSTYPE" == "msys" || "$OSTYPE" == "cygwin" || "$OSTYPE" == "win32" ]]; then
    echo "Building for Windows..."
    
    # 构建Windows DLL
    cd "$PLUGIN_DIR"
    go build -buildmode=c-shared -o "$TARGET_DIR/plugin-e.dll" .
    cd ..
    
    # 构建主程序
    go build -o run.exe main.go plugin_loader.go
    
    echo "Build completed: run.exe and plugin-e/target/plugin-e.dll"
else
    echo "Building for Unix-like system..."
    
    # 构建Unix插件
    cd "$PLUGIN_DIR"
    go build -buildmode=plugin -o "$TARGET_DIR/plugin-e.so" .
    cd ..
    
    # 构建主程序
    go build -o run main.go plugin_loader.go
    
    echo "Build completed: run and plugin-e/target/plugin-e.so"
fi
