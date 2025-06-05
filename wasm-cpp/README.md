# CPP WASM 构建说明

这个目录包含了用于编译 `register_types.cpp` 的构建配置，专门解决了 WebAssembly 内存共享状态不匹配的问题。

## 文件说明

- `register_types.cpp` - Godot GDExtension 的注册代码，包含 `summator_library_init` 函数
- `gdextension_interface.h` - Godot GDExtension 接口头文件
- `Makefile` - Make 构建脚本（推荐使用）
- `README.md` - 本说明文件

## 问题解决

### WebAssembly 内存共享状态不匹配错误

之前遇到的错误：
```
WebAssembly.instantiate(): Import #24 "env" "memory": mismatch in shared state of memory, declared = 0, imported = 1
```

**解决方案**：添加了正确的 Emscripten 编译参数：
- `-sIMPORTED_MEMORY=1` - 使用导入的内存而不是自己的内存
- `-sUSE_PTHREADS=1` - 启用线程支持，与共享内存配合使用
- `-sSIDE_MODULE=1` - 编译为侧模块
- `-sSUPPORT_LONGJMP='wasm'` - 支持 WASM longjmp

## 构建方法

### 使用 Make（推荐）

```bash
# 自动检测并构建
make

# 强制编译为 WASM
make wasm

# 编译为 WASM (带线程支持，现在默认包含)
make wasm-threads

# 使用本地编译器
make native

# 清理
make clean

# 查看帮助
make help
```

### 手动编译 WASM

```bash
em++ -O2 -I. \
     -DWEB_ENABLED -DUNIX_ENABLED \
     -sSIDE_MODULE=1 \
     -sSUPPORT_LONGJMP='wasm' \
     -sIMPORTED_MEMORY=1 \
     -sUSE_PTHREADS=1 \
     register_types.cpp \
     -o libgdsummator.web.template_debug.wasm32.wasm
```

### 手动编译本地库

```bash
# macOS
g++ -O2 -I. -fPIC -dynamiclib register_types.cpp -o libsummator.dylib

# Linux
g++ -O2 -I. -fPIC -shared register_types.cpp -o libsummator.so
```

## 输出文件

根据编译方式不同，会生成以下文件：

- **WASM 编译**: `libgdsummator.web.template_debug.wasm32.wasm` - 与 Godot 兼容的 WASM 模块
- **本地编译**: 
  - macOS: `libsummator.dylib`
  - Linux: `libsummator.so`
  - Windows: `summator.dll`

## 集成说明

### 文件名约定

生成的 WASM 文件命名为 `libgdsummator.web.template_debug.wasm32.wasm`，这与项目的构建脚本 `buildcpp.sh` 预期的文件名匹配。

### 函数导出

代码导出了以下函数：
- `summator_library_init` - GDExtension 初始化入口点
- `initialize_summator_types` - 类型初始化函数
- `uninitialize_summator_types` - 类型清理函数

## 依赖要求

### WASM 编译
- [Emscripten](https://emscripten.org/) 工具链
- 支持 SharedArrayBuffer 的现代浏览器（用于线程支持）

### 本地编译
- GCC 或 Clang 编译器
- 支持 C++11 或更高版本

## 使用流程

1. **编译 WASM 模块**:
   ```bash
   cd cppwasm
   make wasm
   ```

2. **运行项目构建脚本**:
   ```bash
   cd ..
   ./buildcpp.sh
   ```

3. **启动服务器测试**:
   ```bash
   ./run.sh
   ```

## 注意事项

1. **内存兼容性**: 使用 `-sIMPORTED_MEMORY=1` 确保与主程序的内存状态兼容
2. **线程支持**: 包含 `-sUSE_PTHREADS=1` 以支持多线程环境
3. **文件命名**: 确保生成的文件名与项目预期一致
4. **浏览器兼容性**: 需要支持 WebAssembly 和 SharedArrayBuffer 的现代浏览器

## 故障排除

### 常见错误

1. **内存状态不匹配**: 确保使用了 `-sIMPORTED_MEMORY=1` 参数
2. **编译失败**: 检查 Emscripten 工具链是否正确安装
3. **运行时错误**: 确保浏览器支持所需的 WebAssembly 特性

### 验证编译结果

```bash
# 检查文件大小（应该是几KB，不是几百KB）
ls -la libgdsummator.web.template_debug.wasm32.wasm

# 检查文件类型
file libgdsummator.web.template_debug.wasm32.wasm
```

预期输出应该显示这是一个 WebAssembly 二进制模块。
