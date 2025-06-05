# Makefile for compiling register_types.cpp
# 支持 WASM (Emscripten) 和本地编译

# 变量定义
CXX_NATIVE = g++
CXX_WASM = em++
CFLAGS = -O2 -I.
DEFINES_WASM = -DWEB_ENABLED -DUNIX_ENABLED
WASM_FLAGS = -sSIDE_MODULE=1 -sSUPPORT_LONGJMP='wasm' -sIMPORTED_MEMORY=1 -sUSE_PTHREADS=1
THREAD_FLAGS = -sUSE_PTHREADS=1

# 检测操作系统
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Darwin)
    SHARED_EXT = dylib
    SHARED_FLAG = -dynamiclib
else ifeq ($(UNAME_S),Linux)
    SHARED_EXT = so
    SHARED_FLAG = -shared
else
    SHARED_EXT = so
    SHARED_FLAG = -shared
endif

# 源文件
SOURCES = register_types.cpp

# 目标文件
TARGET_NATIVE = libsummator.$(SHARED_EXT)
TARGET_WASM = libgdsummator.web.template_debug.wasm32.wasm

# 默认目标
.PHONY: all wasm native clean help

all: check_emcc

# 检查 Emscripten 是否可用
check_emcc:
	@if command -v emcc >/dev/null 2>&1; then \
		echo "检测到 Emscripten，编译为 WASM"; \
		$(MAKE) wasm; \
	else \
		echo "未找到 Emscripten，使用本地编译器"; \
		$(MAKE) native; \
	fi

# WASM 编译 - 包含线程和内存共享支持
wasm:
	@echo "使用 Emscripten 编译为 WASM..."
	$(CXX_WASM) $(CFLAGS) $(DEFINES_WASM) $(WASM_FLAGS) $(SOURCES) -o $(TARGET_WASM)
	@echo "WASM 编译完成: $(TARGET_WASM)"

# WASM 编译 (带线程支持) - 现在和普通wasm相同
wasm-threads:
	@echo "使用 Emscripten 编译为 WASM (线程支持)..."
	$(CXX_WASM) $(CFLAGS) $(DEFINES_WASM) $(WASM_FLAGS) $(SOURCES) -o $(TARGET_WASM)
	@echo "WASM 编译完成 (线程支持): $(TARGET_WASM)"

# 本地编译
native:
	@echo "使用本地编译器编译..."
	$(CXX_NATIVE) $(CFLAGS) -fPIC $(SHARED_FLAG) $(SOURCES) -o $(TARGET_NATIVE)
	@echo "本地编译完成: $(TARGET_NATIVE)"

# 清理
clean:
	@echo "清理构建文件..."
	rm -f *.o *.wasm *.so *.dylib *.dll *.a
	@echo "清理完成"

# 帮助
help:
	@echo "可用目标:"
	@echo "  all          - 自动检测并选择编译方式"
	@echo "  wasm         - 使用 Emscripten 编译为 WASM"
	@echo "  wasm-threads - 使用 Emscripten 编译为 WASM (线程支持)"
	@echo "  native       - 使用本地编译器编译"
	@echo "  clean        - 清理构建文件"
	@echo "  help         - 显示此帮助信息" 