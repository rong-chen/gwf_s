# Go 编译器 用于一键打包 
# make build 打包
# make clean 清除打包文件

GO := go

# 项目名称
APP_NAME := gwf
#dist路径
DIST_DIR := ./dist

# 编译输出目录
BUILD_DIR := ./build
#config.yaml路径
CONFIG_FILE := ./config.yaml
# 主程序入口文件
MAIN_FILE := ./main.go

# 编译命令
build:
	@echo 开始构建 $(APP_NAME)...
	@make clean
	@GOOS=linux GOARCH=amd64 $(GO) build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE) 
	@if [ -d $(DIST_DIR) ]; then \
        cp -r $(DIST_DIR)/* $(BUILD_DIR)/dist; \
        echo "复制静态资源完成"; \
    else \
        echo "没有找到静态资源文件，跳过复制"; \
    fi
	@cp $(CONFIG_FILE) $(BUILD_DIR)
	@echo 构建完成，输出文件为 $(BUILD_DIR)/$(APP_NAME)
# 清理命令
clean:
	@echo 清理旧版本
	@rm -rf $(BUILD_DIR)
	@echo 清理完成

# 默认目标
.DEFAULT_GOAL := build
