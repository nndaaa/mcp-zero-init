# MCP-Zero 项目初始化工具

自动检测并配置 MCP-Zero 环境，让你零配置开始使用 go-zero 开发。

## 功能

- ✅ 自动检测 Go 环境
- ✅ 自动检测/安装 goctl
- ✅ 自动检测/安装 mcp-zero
- ✅ 自动创建 Claude Code MCP 配置
- ✅ 支持项目级配置（推荐）
- ✅ 支持全局配置
- ✅ 支持更新现有配置

## 安装

```bash
go install github.com/nndaaa/mcp-zero-init@latest
```

## 使用

### 在项目目录中初始化（推荐）

```bash
cd your-go-zero-project
mcp-zero-init
```

### 全局安装

```bash
mcp-zero-init -global
```

## License

MIT
