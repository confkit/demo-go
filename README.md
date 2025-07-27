# Demo Go - Hello World 服务

这是一个简单的Go HTTP服务，提供Hello World响应。

## 快速开始

### 使用Task（推荐）

首先安装工具：
```bash
# macOS
brew install go-task/tap/go-task

# 或者使用go install
go install github.com/go-task/task/v3/cmd/task@latest

# 安装 air
go install github.com/cosmtrek/air@latest
```

然后使用以下命令：

```bash
# 查看所有可用任务
task

# 运行服务
task run

# 构建项目
task build

# 构建并运行
task start

# 开发模式（自动重启）
task dev
```

### 传统方式

```bash
# 运行服务
go run main.go

# 构建项目
go build -o bin/demo-go main.go
./bin/demo-go
```

## 访问

服务启动后，访问 http://localhost:8090 即可看到 "Hello World!" 响应。

## 可用任务

- `task run` - 启动HTTP服务
- `task build` - 构建可执行文件
- `task start` - 构建并运行程序
- `task dev` - 开发模式（自动重启）
- `task clean` - 清理构建文件
- `task test` - 运行测试
- `task fmt` - 格式化代码
- `task check` - 检查代码
- `task deps` - 获取依赖
- `task all` - 完整构建流程

## 功能

- 提供基本的HTTP服务
- 访问根路径 `/` 返回 "Hello World!"
- 默认监听8090端口
- 支持Task任务管理 