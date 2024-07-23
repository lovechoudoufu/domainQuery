# 可执行文件名称前缀
BINARY_NAME=domainQuery

# 版本号
VERSION=1.0.0

# 主程序路径
MAIN_PATH=./cmd/domainQuery

# 目标平台
TARGETS=windows-386 windows-amd64 linux-386 linux-amd64 linux-arm linux-arm64 darwin-amd64 darwin-arm64

.PHONY: all clean $(TARGETS)

all: $(TARGETS)

# 清理
clean:
	@rm -rf release/

# 为每个目标平台编译
$(TARGETS):
	$(eval OS := $(word 1,$(subst -, ,$@)))
	$(eval ARCH := $(word 2,$(subst -, ,$@)))
	@echo "[+] 编译 $(OS)/$(ARCH)..."
	@mkdir -p release/$(OS)-$(ARCH)
	@GOOS=$(OS) GOARCH=$(ARCH) go build -trimpath -ldflags "-s -w" -o release/$(OS)-$(ARCH)/$(BINARY_NAME)_$(VERSION)$(if $(findstring windows,$(OS)),.exe,) $(MAIN_PATH)

# 编译所有平台
build-all: $(TARGETS)
	@echo "所有平台编译完成!"

# 运行测试
test:
	@go test -v ./...

# 显示帮助
help:
	@echo "可用的 make 命令:"
	@echo "  all        - 编译所有平台"
	@echo "  clean      - 清理构建目录"
	@echo "  test       - 运行测试"
	@echo "  build-all  - 编译所有平台 (同 all)"
	@echo "  windows-386   - 仅编译 Windows 32位版本"
	@echo "  windows-amd64 - 仅编译 Windows 64位版本"
	@echo "  linux-amd64   - 仅编译 Linux 64位版本"
	@echo "  linux-arm64   - 仅编译 Linux ARM64 版本"
	@echo "  darwin-amd64  - 仅编译 macOS 64位版本"