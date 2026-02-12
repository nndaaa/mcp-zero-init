# MCP-Zero 项目初始化工具

[![Go Version](https://img.shields.io/badge/go-1.23+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

自动检测并配置 MCP-Zero 环境，让你**零配置**开始使用 go-zero 开发。

## ✨ 功能特性

- 🔍 自动检测 Go 环境（GOPATH/GOROOT）
- 📦 自动检测/安装 [goctl](https://github.com/zeromicro/go-zero/tree/master/tools/goctl)
- 📦 自动检测/安装 [mcp-zero](https://github.com/zeromicro/mcp-zero)
- ⚙️ 自动创建 Claude Code MCP 配置
- 🌍 支持项目级和全局配置

## 📦 安装

```bash
go install github.com/nndaaa/mcp-zero-init@latest
```

确保 `$GOPATH/bin` 在你的 PATH 中。

## 🚀 快速开始

### 方式一：项目级配置（推荐）

在你的 go-zero 项目目录中运行：

```bash
cd your-go-zero-project
mcp-zero-init
```

这会在当前目录创建 `.claude/servers.json` 配置。

### 方式二：全局配置

```bash
mcp-zero-init -global
```

这会在 `~/.claude/servers.json` 创建全局配置。

## 📖 使用指南

### 检查环境

```bash
mcp-zero-init -check
```

输出示例：

```
🚀 MCP-Zero 项目初始化工具
============================
✓ 找到Go bin目录: C:\Users\Administrator\go\bin

📋 环境检查报告
================
✅ goctl: C:\Users\Administrator\go\bin\goctl.exe
   版本: goctl version 1.8.4 windows/amd64
✅ mcp-zero: C:\Users\Administrator\go\bin\mcp-zero.exe
```

### 强制更新配置

```bash
mcp-zero-init -force
```

## 📋 命令参数

| 参数 | 说明 |
|------|------|
| `-global` | 创建全局配置 |
| `-force` | 强制覆盖现有配置 |
| `-check` | 仅检查环境，不创建配置 |

## 🔧 完整示例

### 场景：新项目初始化

```bash
# 1. 创建新项目目录
mkdir my-api-service
cd my-api-service

# 2. 初始化 go mod
go mod init my-api-service

# 3. 初始化 MCP-Zero
mcp-zero-init

# 输出：
# 🚀 MCP-Zero 项目初始化工具
# ============================
# ✓ 找到Go bin目录: C:\Users\Administrator\go\bin
# ✓ goctl 已就绪: C:\Users\Administrator\go\bin\goctl.exe
# ✓ mcp-zero 已就绪: C:\Users\Administrator\go\bin\mcp-zero.exe
# 📁 创建项目配置...
#
# ✅ 初始化完成!
# 📄 配置文件: C:\projects\my-api-service\.claude\servers.json
#
# 下一步:
#   1. 重启 Claude Code 以加载 MCP 配置
#   2. 输入 /mcp 查看MCP服务器状态
```

## ❓ 常见问题

### "无法找到Go环境"

确保 Go 已正确安装：

```bash
go version
go env GOPATH
```

### "安装 goctl 失败"

手动安装：

```bash
go install github.com/zeromicro/go-zero/tools/goctl@latest
```

### 配置不生效

1. 重启 Claude Code
2. 检查配置文件路径
3. 在 Claude Code 中输入 `/mcp` 查看服务器状态

## 📚 相关资源

- [MCP-Zero](https://github.com/zeromicro/mcp-zero) - MCP 服务器主项目
- [go-zero](https://github.com/zeromicro/go-zero) - 云原生微服务框架
- [go-zero 文档](https://go-zero.dev/)
- [Claude Code 文档](https://docs.anthropic.com/en/docs/agents-and-tools/claude-code/overview)

## 📄 详细文档

- [PROJECT_SUMMARY.md](./PROJECT_SUMMARY.md) - 项目完整文档
- [CHANGELOG.md](./CHANGELOG.md) - 更新日志

## 🤝 贡献

欢迎提交 Issue 和 PR！

## 📜 许可证

MIT License

---

**Happy coding with go-zero! 🚀**
