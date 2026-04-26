.PHONY: all build-windows build-unix clean

PLUGIN_DIR=plugin-e
TARGET_DIR=$(PLUGIN_DIR)/target
MAIN_SOURCES=main.go plugin_loader.go

# 默认目标
all:
ifeq ($(OS),Windows_NT)
	$(MAKE) build-windows
else
	$(MAKE) build-unix
endif

# Windows构建
build-windows:
	@echo "Building for Windows..."
	mkdir -p $(TARGET_DIR)
	cd $(PLUGIN_DIR) && go build -buildmode=c-shared -o $(TARGET_DIR)/plugin-e.dll .
	go build -o run.exe $(MAIN_SOURCES)
	@echo "Build completed: run.exe and $(TARGET_DIR)/plugin-e.dll"

# Unix-like系统构建
build-unix:
	@echo "Building for Unix-like system..."
	mkdir -p $(TARGET_DIR)
	cd $(PLUGIN_DIR) && go build -buildmode=plugin -o $(TARGET_DIR)/plugin-e.so .
	go build -o run $(MAIN_SOURCES)
	@echo "Build completed: run and $(TARGET_DIR)/plugin-e.so"

# 清理
clean:
	rm -f run run.exe
	rm -rf $(TARGET_DIR)
