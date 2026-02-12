package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// MCPConfig MCP服务器配置
type MCPConfig struct {
	MCPServers map[string]ServerConfig `json:"mcpServers"`
}

type ServerConfig struct {
	Command string            `json:"command"`
	Env     map[string]string `json:"env"`
}

var (
	globalFlag = flag.Bool("global", false, "创建全局配置 (~/.claude/servers.json)")
	forceFlag  = flag.Bool("force", false, "强制覆盖现有配置")
	checkFlag  = flag.Bool("check", false, "仅检查环境，不创建配置")
)

func main() {
	flag.Parse()

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "❌ 错误: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	fmt.Println("🚀 MCP-Zero 项目初始化工具")
	fmt.Println("============================")

	// 1. 检测 Go 环境
	goBin, err := getGoBin()
	if err != nil {
		return fmt.Errorf("无法找到Go环境: %w", err)
	}
	fmt.Printf("✓ 找到Go bin目录: %s\n", goBin)

	// 2. 检测 goctl
	goctlPath, goctlExists, err := checkGoctl(goBin)
	if err != nil {
		return err
	}

	// 3. 检测 mcp-zero
	mcpZeroPath, mcpZeroExists, err := checkMcpZero(goBin)
	if err != nil {
		return err
	}

	// 如果仅检查模式，到此结束
	if *checkFlag {
		printStatus(goctlExists, mcpZeroExists, goctlPath, mcpZeroPath)
		return nil
	}

	// 4. 安装缺失的工具
	if !goctlExists {
		goctlPath, err = installGoctl(goBin)
		if err != nil {
			return fmt.Errorf("安装 goctl 失败: %w", err)
		}
	}
	fmt.Printf("✓ goctl 已就绪: %s\n", goctlPath)

	if !mcpZeroExists {
		mcpZeroPath, err = installMcpZero(goBin)
		if err != nil {
			return fmt.Errorf("安装 mcp-zero 失败: %w", err)
		}
	}
	fmt.Printf("✓ mcp-zero 已就绪: %s\n", mcpZeroPath)

	// 5. 创建配置文件
	configPath, err := createConfig(goctlPath, mcpZeroPath)
	if err != nil {
		return err
	}

	fmt.Printf("\n✅ 初始化完成!\n")
	fmt.Printf("📄 配置文件: %s\n", configPath)
	fmt.Println("\n下一步:")
	fmt.Println("  1. 重启 Claude Code 以加载 MCP 配置")
	fmt.Println("  2. 输入 /mcp 查看MCP服务器状态")
	fmt.Println("\n使用示例:")
	fmt.Println("  - 创建一个用户服务，端口8080")
	fmt.Println("  - 生成数据库模型")
	fmt.Println("  - 分析项目结构")

	return nil
}

// printStatus 打印环境状态
func printStatus(goctlExists, mcpZeroExists bool, goctlPath, mcpZeroPath string) {
	fmt.Println("\n📋 环境检查报告")
	fmt.Println("================")

	if goctlExists {
		fmt.Printf("✅ goctl: %s\n", goctlPath)
		// 显示版本
		if out, err := exec.Command(goctlPath, "--version").Output(); err == nil {
			fmt.Printf("   版本: %s\n", strings.TrimSpace(string(out)))
		}
	} else {
		fmt.Println("❌ goctl: 未安装")
	}

	if mcpZeroExists {
		fmt.Printf("✅ mcp-zero: %s\n", mcpZeroPath)
	} else {
		fmt.Println("❌ mcp-zero: 未安装")
	}
}

// getGoBin 获取Go可执行文件目录
func getGoBin() (string, error) {
	// 首先尝试 GOPATH/bin
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		cmd := exec.Command("go", "env", "GOPATH")
		out, err := cmd.Output()
		if err == nil {
			goPath = strings.TrimSpace(string(out))
		}
	}

	if goPath != "" {
		return filepath.Join(goPath, "bin"), nil
	}

	// 其次尝试 GOROOT/bin
	goRoot := os.Getenv("GOROOT")
	if goRoot == "" {
		cmd := exec.Command("go", "env", "GOROOT")
		out, err := cmd.Output()
		if err == nil {
			goRoot = strings.TrimSpace(string(out))
		}
	}

	if goRoot != "" {
		return filepath.Join(goRoot, "bin"), nil
	}

	return "", fmt.Errorf("无法找到 GOPATH 或 GOROOT，请确保 Go 已正确安装")
}

// checkGoctl 检查 goctl 是否存在
func checkGoctl(goBin string) (string, bool, error) {
	goctlPath := filepath.Join(goBin, "goctl")
	if runtime.GOOS == "windows" {
		goctlPath += ".exe"
	}

	if _, err := os.Stat(goctlPath); err == nil {
		return goctlPath, true, nil
	}

	// 也检查 PATH
	if path, err := exec.LookPath("goctl"); err == nil {
		return path, true, nil
	}

	return goctlPath, false, nil
}

// checkMcpZero 检查 mcp-zero 是否存在
func checkMcpZero(goBin string) (string, bool, error) {
	mcpZeroPath := filepath.Join(goBin, "mcp-zero")
	if runtime.GOOS == "windows" {
		mcpZeroPath += ".exe"
	}

	if _, err := os.Stat(mcpZeroPath); err == nil {
		return mcpZeroPath, true, nil
	}

	// 也检查 PATH
	if path, err := exec.LookPath("mcp-zero"); err == nil {
		return path, true, nil
	}

	return mcpZeroPath, false, nil
}

// installGoctl 安装 goctl
func installGoctl(goBin string) (string, error) {
	fmt.Println("📦 正在安装 goctl...")
	cmd := exec.Command("go", "install", "github.com/zeromicro/go-zero/tools/goctl@latest")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return "", err
	}

	goctlPath := filepath.Join(goBin, "goctl")
	if runtime.GOOS == "windows" {
		goctlPath += ".exe"
	}

	// 验证安装
	if _, err := os.Stat(goctlPath); err != nil {
		return "", fmt.Errorf("安装后找不到 goctl")
	}

	return goctlPath, nil
}

// installMcpZero 安装 mcp-zero
func installMcpZero(goBin string) (string, error) {
	fmt.Println("📦 正在安装 mcp-zero...")

	// 尝试从远程安装
	cmd := exec.Command("go", "install", "github.com/zeromicro/mcp-zero@latest")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("⚠️ 远程安装失败，尝试本地构建...")
		return buildLocalMcpZero(goBin)
	}

	mcpZeroPath := filepath.Join(goBin, "mcp-zero")
	if runtime.GOOS == "windows" {
		mcpZeroPath += ".exe"
	}

	// 验证安装
	if _, err := os.Stat(mcpZeroPath); err != nil {
		return "", fmt.Errorf("安装后找不到 mcp-zero")
	}

	return mcpZeroPath, nil
}

// buildLocalMcpZero 从本地源码构建 mcp-zero
func buildLocalMcpZero(goBin string) (string, error) {
	// 获取当前可执行文件路径
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}

	// 找到项目根目录 (假设当前在 cmd/mcp-zero-init/ 下)
	currentDir := filepath.Dir(exePath)
	projectRoot := filepath.Dir(currentDir) // 到 mcp-zero 根目录

	// 检查是否是正确的目录
	mainGo := filepath.Join(projectRoot, "main.go")
	if _, err := os.Stat(mainGo); err != nil {
		// 可能当前是在开发环境，尝试从当前工作目录找
		cwd, _ := os.Getwd()
		projectRoot = cwd
		mainGo = filepath.Join(projectRoot, "main.go")
		if _, err := os.Stat(mainGo); err != nil {
			return "", fmt.Errorf("找不到 mcp-zero 源码，请手动安装")
		}
	}

	fmt.Printf("📦 从本地源码构建: %s\n", projectRoot)

	mcpZeroPath := filepath.Join(goBin, "mcp-zero")
	if runtime.GOOS == "windows" {
		mcpZeroPath += ".exe"
	}

	cmd := exec.Command("go", "build", "-o", mcpZeroPath, mainGo)
	cmd.Dir = projectRoot
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("本地构建失败: %w", err)
	}

	return mcpZeroPath, nil
}

// createConfig 创建配置文件
func createConfig(goctlPath, mcpZeroPath string) (string, error) {
	// 确定配置目录
	var configDir string
	if *globalFlag {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("无法获取用户目录: %w", err)
		}
		configDir = filepath.Join(homeDir, ".claude")
		fmt.Println("🌐 创建全局配置...")
	} else {
		cwd, err := os.Getwd()
		if err != nil {
			return "", fmt.Errorf("无法获取当前目录: %w", err)
		}
		configDir = filepath.Join(cwd, ".claude")
		fmt.Println("📁 创建项目配置...")
	}

	// 检查是否已存在
	configPath := filepath.Join(configDir, "servers.json")
	if _, err := os.Stat(configPath); err == nil && !*forceFlag {
		return "", fmt.Errorf("配置文件已存在: %s (使用 -force 覆盖)", configPath)
	}

	// 创建目录
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "", fmt.Errorf("创建目录失败: %w", err)
	}

	// 构建配置
	config := MCPConfig{
		MCPServers: map[string]ServerConfig{
			"mcp-zero": {
				Command: mcpZeroPath,
				Env: map[string]string{
					"GOCTL_PATH": goctlPath,
				},
			},
		},
	}

	// 序列化为 JSON
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return "", fmt.Errorf("序列化配置失败: %w", err)
	}

	// 写入文件
	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return "", fmt.Errorf("写入配置文件失败: %w", err)
	}

	return configPath, nil
}
