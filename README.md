# 雷神加速器自动暂停工具

![GitHub Actions](https://img.shields.io/badge/GitHub%20Actions-Ready-brightgreen)
![Go Version](https://img.shields.io/badge/Go-1.24-blue)
![License](https://img.shields.io/badge/License-MIT-yellow)

这是一个用Go语言编写的雷神加速器（NN加速器）自动暂停工具，通过GitHub Actions自动运行，无需本地环境配置。

## 🚀 快速开始

1. **Fork本项目** → 2. **获取TOKEN** → 3. **配置Secrets** → 4. **运行Actions** ✅

### 1. Fork项目

点击右上角的 **Fork** 按钮，将项目fork到你的GitHub账户

### 2. 获取雷神加速器TOKEN
![image](config/get-token.png)

**快速步骤**：
1. **登录雷神加速器官网**: https://vip.leigod.com/user.html
2. **打开浏览器开发者工具**（按F12键）
4. **在网站上进行任意操作**（如点击暂停、开始等按钮）
5. **查找包含 `account_token` 的请求**:
   - 在Network面板中找到发送到 `webapi.leigod.com` 的请求
   - 点击请求，查看 **Request Payload** 或 **Form Data**
   - 复制 `account_token` 的值

### 3. 配置GitHub Secrets

1. **进入你fork的仓库**
2. **点击 Settings(设置) 标签页**
3. **在左侧菜单中选择 Secrets and variables > Actions**
4. **点击 New repository secret**
5. **添加密钥**:
   - Name: `TOKEN`
   - Secret: 粘贴你获取的 `account_token` 值
6. **点击 Add secret 保存**

### 4. 运行GitHub Action

1. **进入 Actions 标签页**
2. **选择 "Auto Pause Leishen" 工作流**
3. **点击 "Run workflow" 按钮**
4. **选择分支（通常是main）并点击绿色的 "Run workflow" 按钮**

### 5. 自动定时运行

配置完成后，GitHub Actions会在每天凌晨1点（北京时间）自动运行暂停程序。

你也可以修改 `.github/workflows/auto-pause.yml` 文件中的 cron 表达式来调整运行时间：

```yaml
schedule:
  # 每天凌晨1点自动暂停 (UTC时间17:00 = 北京时间1:00)
  - cron: '0 17 * * *'
```


## API说明


### 错误码说明

- `0`: 操作成功
- `400803`: 账号已经停止加速，请不要重复操作
- 其他错误码: 请参考雷神加速器官方API文档



## 本地开发

如果你想在本地运行或修改代码：

### 1. 克隆项目

```bash
git clone https://github.com/your-username/leishen-auto.git
cd leishen-auto
```

### 2. 设置环境变量

```bash
# Windows (PowerShell)
$env:TOKEN="your_account_token_here"

# Linux/macOS
export TOKEN="your_account_token_here"
```

### 3. 运行程序

```bash
go run main.go
```


## 功能特性

- 🌀 自动暂停雷神加速器
- 🔄 支持重复操作检测（错误码400803）
- 🎯 简洁的命令行输出
- 📦 模块化代码结构
- ⚡ 高性能Go语言实现

## 项目结构

```
leishen-auto/
├── api/           # API客户端包
│   └── client.go  # 雷神API客户端实现
├── config/        # 配置包
│   └── config.go  # 配置加载和管理
├── main.go        # 主程序入口
├── go.mod         # Go模块文件
└── README.md      # 项目说明文档
```

## 环境要求

- Go 1.19 或更高版本
- 有效的雷神加速器账户令牌

## 通过GitHub Actions使用（推荐）

## 常见问题

### Q: 如何修改自动运行时间？
A: 编辑 `.github/workflows/auto-pause.yml` 文件中的 cron 表达式。例如：
- `0 17 * * *` = 每天凌晨1点（北京时间）
- `0 9 * * *` = 每天下午5点（北京时间）

### Q: TOKEN在哪里获取？
A: 登录雷神加速器官网，打开浏览器开发者工具，在Network面板中查看任意请求的 `account_token` 参数。

### Q: 如何查看运行日志？
A: 在GitHub仓库的Actions标签页中，点击对应的工作流运行记录即可查看详细日志。

### Q: 支持哪些操作？
A: 目前只支持暂停操作。如需其他功能（如开始加速），可以基于现有代码进行扩展。

## 许可证

本项目采用MIT许可证，详情请参见LICENSE文件。


<!-- Security scan triggered at 2026-09-05 07:36:50 -->