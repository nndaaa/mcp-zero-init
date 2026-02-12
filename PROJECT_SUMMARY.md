# MCP-Zero-Init 项目总结文档

> 📅 创建日期: 2025-02-12
> 👤 维护者: nndaaa
> 🏷️ 版本: v1.0.0
> 📦 仓库: https://github.com/nndaaa/mcp-zero-init

---

## 🎯 项目概述

**MCP-Zero-Init** 是一个自动化配置工具，帮助开发者一键设置 MCP-Zero 开发环境，实现零配置开始使用 go-zero 框架。

### 解决的问题

1. **环境配置复杂**: 新手需要手动安装 goctl、配置 MCP 服务器
2. **路径问题**: GOPATH/GOROOT 环境变量配置容易出错
3. **配置繁琐**: Claude Code 的 MCP 配置需要手动编辑 JSON 文件

### 核心功能

- ✅ 自动检测 Go 环境（GOPATH/GOROOT）
- ✅ 自动检测/安装 goctl
- ✅ 自动检测/安装 mcp-zero
- ✅ 自动创建 Claude Code MCP 配置
- ✅ 支持项目级和全局配置

---

## 📂 项目结构

```
mcp-zero-init/
├── main.go           # 主程序入口
├── go.mod            # Go 模块定义
├── README.md         # 用户文档
└── PROJECT_SUMMARY.md # 本文档
```

### 源码说明

**main.go** 包含以下核心模块：

| 函数 | 功能 |
|------|------|
| `getGoBin()` | 检测 Go 安装目录（GOPATH/bin 或 GOROOT/bin） |
| `checkGoctl()` | 检查 goctl 是否已安装 |
| `checkMcpZero()` | 检查 mcp-zero 是否已安装 |
| `installGoctl()` | 自动安装 goctl |
| `installMcpZero()` | 自动安装 mcp-zero |
| `createConfig()` | 创建 Claude Code MCP 配置文件 |

---

## 🚀 快速开始

### 安装

```bash
go install github.com/nndaaa/mcp-zero-init@latest
```

### 使用

#### 1. 项目级配置（推荐）

```bash
cd your-go-zero-project
mcp-zero-init
```

在当前目录创建 `.claude/servers.json`

#### 2. 全局配置

```bash
mcp-zero-init -global
```

在 `~/.claude/servers.json` 创建全局配置

#### 3. 检查环境

```bash
mcp-zero-init -check
```

仅检查环境，不创建配置

#### 4. 强制更新

```bash
mcp-zero-init -force
```

覆盖已存在的配置

---

## 📖 使用示例

### 示例 1: 首次使用

```bash
$ cd my-new-project
$ mcp-zero-init

🚀 MCP-Zero 项目初始化工具
============================
✓ 找到Go bin目录: C:\Users\Administrator\go\bin
📦 正在安装 goctl...
go: downloading github.com/zeromicro/go-zero v1.8.4
✓ goctl 已就绪: C:\Users\Administrator\go\bin\goctl.exe
📦 正在安装 mcp-zero...
go: downloading github.com/zeromicro/mcp-zero v1.0.0
✓ mcp-zero 已就绪: C:\Users\Administrator\go\bin\mcp-zero.exe
📁 创建项目配置...

✅ 初始化完成!
📄 配置文件: C:\projects\my-new-project\.claude\servers.json

下一步:
  1. 重启 Claude Code 以加载 MCP 配置
  2. 输入 /mcp 查看MCP服务器状态
```

### 示例 2: 环境已就绪

```bash
$ mcp-zero-init -check

🚀 MCP-Zero 项目初始化工具
============================
✓ 找到Go bin目录: C:\Users\Administrator\go\bin

📋 环境检查报告
================
✅ goctl: C:\Users\Administrator\go\bin\goctl.exe
   版本: goctl version 1.8.4 windows/amd64
✅ mcp-zero: C:\Users\Administrator\go\bin\mcp-zero.exe
```

---

## 🔧 配置说明

### 项目级配置

文件: `your-project/.claude/servers.json`

```json
{
  "mcpServers": {
    "mcp-zero": {
      "command": "C:\\Users\\Administrator\\go\\bin\\mcp-zero.exe",
      "env": {
        "GOCTL_PATH": "C:\\Users\\Administrator\\go\\bin\\goctl.exe"
      }
    }
  }
}
```

### 全局配置

- **Windows**: `C:\Users\{用户名}\.claude\servers.json`
- **macOS/Linux**: `~/.claude/servers.json`

---

## 📋 命令行参数

| 参数 | 说明 | 示例 |
|------|------|------|
| `-global` | 创建全局配置 | `mcp-zero-init -global` |
| `-force` | 强制覆盖现有配置 | `mcp-zero-init -force` |
| `-check` | 仅检查环境 | `mcp-zero-init -check` |

---

## 🔍 工作原理

```
┌─────────────────┐
│  mcp-zero-init  │
└────────┬────────┘
         │
    ┌────┴────┐
    ▼         ▼
┌───────┐  ┌────────┐
│检测Go │  │检查工具│
│环境   │  │安装状态│
└───────┘  └────┬───┘
                │
        ┌───────┼───────┐
        ▼       ▼       ▼
    ┌──────┐┌──────┐┌──────┐
    │goctl ││mcp-  ││创建  │
    │安装  ││zero  ││配置  │
    │      ││安装  ││文件  │
    └──────┘└──────┘└──────┘
```

---

## 🛠️ 开发与调试

### 本地开发

```bash
# 克隆仓库
git clone https://github.com/nndaaa/mcp-zero-init.git
cd mcp-zero-init

# 本地测试
go run main.go -check

# 构建
go build -o mcp-zero-init.exe

# 安装到 GOPATH/bin
go install
```

### 调试技巧

1. **检查 Go 环境**
   ```bash
   go version
   go env GOPATH
   go env GOROOT
   ```

2. **验证工具安装**
   ```bash
   goctl --version
   which mcp-zero  # Linux/macOS
   where mcp-zero  # Windows
   ```

3. **查看配置文件**
   ```bash
   cat .claude/servers.json
   ```

---

## ❓ 常见问题

### Q: "无法找到Go环境"

**A**: 确保 Go 已正确安装：
```bash
go version
go env GOPATH
```

### Q: "安装 goctl 失败"

**A**: 手动安装：
```bash
go install github.com/zeromicro/go-zero/tools/goctl@latest
```

### Q: 配置不生效

**A**:
1. 重启 Claude Code
2. 检查配置文件路径
3. 在 Claude Code 中输入 `/mcp` 查看状态

### Q: Windows 上找不到 mcp-zero-init

**A**: 确保 `%GOPATH%\bin` 在系统 PATH 中：
```powershell
# 临时添加
$env:PATH += ";$(go env GOPATH)\bin"
```

---

## 🗺️ 路线图

### 已完成功能 ✅

- [x] 自动检测 Go 环境
- [x] 自动安装 goctl
- [x] 自动安装 mcp-zero
- [x] 项目级配置
- [x] 全局配置
- [x] 环境检查模式
- [x] 强制覆盖配置

### 计划功能 📅

- [ ] 配置备份/恢复
- [ ] 多版本管理
- [ ] 配置文件验证
- [ ] 交互式配置向导
- [ ] Docker 支持
- [ ] CI/CD 集成

### 待讨论 💬

- [ ] 支持其他 AI 助手（Copilot、Cursor 等）
- [ ] 配置文件模板自定义
- [ ] 插件机制

---

## 🔗 相关资源

### 官方链接

- **MCP-Zero**: https://github.com/zeromicro/mcp-zero
- **go-zero**: https://github.com/zeromicro/go-zero
- **go-zero 文档**: https://go-zero.dev/

### 本项目

- **GitHub**: https://github.com/nndaaa/mcp-zero-init
- **Issues**: https://github.com/nndaaa/mcp-zero-init/issues

### 社区

- **go-zero Discord**: https://discord.gg/go-zero
- **Claude Code 文档**: https://docs.anthropic.com/en/docs/agents-and-tools/claude-code/overview

---

## 📜 更新日志

### v1.0.0 (2025-02-12)

- 🎉 初始版本发布
- ✅ 支持自动检测和安装
- ✅ 支持项目级和全局配置
- ✅ 支持环境检查模式

---

## 🤝 贡献指南

欢迎提交 Issue 和 PR！

### 提交 Issue

1. 描述问题
2. 提供环境信息（OS、Go 版本）
3. 提供复现步骤

### 提交 PR

1. Fork 仓库
2. 创建功能分支
3. 提交更改
4. 创建 Pull Request

---

## 📄 许可证

MIT License - 详见 [LICENSE](./LICENSE)

---

## 💡 反馈与建议

如果你有任何问题或建议，欢迎通过以下方式联系：

- GitHub Issues: https://github.com/nndaaa/mcp-zero-init/issues
- Email: [你的邮箱]

---

**Happy coding with MCP-Zero! 🚀**
