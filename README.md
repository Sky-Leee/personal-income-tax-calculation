# 个人所得税计算服务

该项目是一个使用 Go 和 Gin 框架实现的个人所得税计算服务，包含后端 RESTful API 和前端页面。

## 目录结构

```
.
├── README.md             # 项目说明文档
├── cmd                   # 主程序
│   └── main.go          # 服务入口
├── go.mod                # Go模块依赖管理
├── go.sum                # Go模块依赖校验
├── internal              # 内部逻辑和处理程序
│   ├── handlers          # HTTP处理程序
│   │   └── tax.go       # 处理税务相关请求
│   └── logic            # 税务计算逻辑
│       └── tax.go       # 税务计算实现
├── pkg                  # 公共API
│   └── apis             # API定义
│       └── tax.go       # 税务相关API
└── static               # 前端静态资源
    └── index.html       # 前端页面
```

## 安装依赖

在项目根目录下运行以下命令来安装依赖：

```bash
go mod tidy
```

## 启动服务

在 `cmd` 目录下运行以下命令启动服务：

```bash
go run main.go
```

服务启动后，访问 `http://localhost:8080` 以查看前端页面。

## 使用说明

1. 在输入框中输入您的收入。
2. 点击“计算税款”按钮。
3. 应缴税款将显示在页面上。
