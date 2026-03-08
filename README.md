# 跨学科启发式教学系统后端 (Teaching Backend)

本项目是一个基于 `go-zero` 微服务框架开发的跨学科启发式教学系统后端项目。系统采用微服务架构设计，将功能拆分为多个相互协同工作的 RPC 和 API 服务，旨在提供稳定、高性能的教学平台后台支持。

## 🎯 项目功能模块

项目主要包含以下几个核心微服务模块（位于 `application` 目录下）：

- **User (用户服务)**：负责用户的注册、登录、权限鉴权（JWT）、个人信息管理等核心用户链路。
- **Course (课程服务)**：负责课程的管理、课程内容的展示、章节进度跟踪等课程核心业务。
- **Exam (考试系统)**：支持试题的获取、考试作答、自动判题和分数结算。
- **AI (智能助手服务)**：对接大语言模型（LLM），提供苏格拉底式启发教学、跨学科内容串联及智能问答等功能。
- **Applet (API 网关/小程序接口)**：作为所有 RPC 服务的统一业务聚合端，提供 RESTful API 供前端应用或小程序调用。

## 🛠 技术栈

本项目使用了以下核心技术与依赖：

### 核心框架
- **语言**: Go (1.25+)
- **微服务框架**: [go-zero](https://github.com/zeromicro/go-zero) v1.10.0 (支持高并发、内置缓存、限流等)
- **RPC 框架**: gRPC & Protobuf

### 存储与中间件
- **关系型数据库**: MySQL
- **缓存**: Redis (提升接口读性能及 session 状态维护)
- **对象存储**: MinIO (负责图片、课件和文件上传的分布式存储)
- **全文检索**: Elasticsearch v8 (实现文章/课程/试题等大文本内容的高效搜索)
- **消息队列**: Kafka (用于服务间异步解耦，如考试提交后异步批改、行为日志记录等)
- **服务发现与配置中心**: ETCD / Consul

## 📂 目录结构

```text
teaching-backend/
├── application/       # 所有微服务代码所在目录
│   ├── ai/            # AI 问答及大模型相关微服务
│   ├── applet/        # 面向客户端的 API Gateway (HTTP)
│   ├── course/        # 课程核心业务(RPC)
│   ├── exam/          # 考试及题库业务(RPC)
│   └── user/          # 用户及权限管理业务(RPC)
├── db/                # 数据库初始化脚本及相关配置
├── docs/              # 随项目存放的设计文档与学习资料
├── pkg/               # 全局公共组件、工具类及常量定义
├── go.mod             # Go Modules 依赖管理
└── README.md          # 项目说明文档
```

## 🚀 快速启动

### 1. 环境准备
在本地运行之前，请确保已安装以下环境：
- Go 1.25.7 或更高版本
- Docker & Docker-Compose (用于快速启动中间件)
- `goctl` 代码生成工具 (go-zero 必备)
  ```bash
  go install github.com/zeromicro/go-zero/tools/goctl@latest
  ```

### 2. 启动基础中间件
启动所需的 MySQL, Redis, ETCD, MinIO 等依赖中间件

### 3. 配置数据库与缓存
将 `etc` 或项目内部的配置文件中关于 MySQL 连接串、Redis 地址修改为你的本地配置，并导入 `db/` 目录下的相关 SQL 数据。

### 4. 运行服务
因为是微服务架构，请分别启动各个 RPC 服务和最后启动 API Gateway 服务。
```bash
# 例如分别进入各个微服务目录启动
cd application/user/rpc
go run user.go -f etc/user.yaml

cd ../../../application/applet/api
go run applet.go -f etc/applet.yaml
```

## 📝 开发规范
- **API 定义**: 使用 `.api` 文件编写 HTTP 接口规范，通过 `goctl api go` 生成路由及 handler。
- **RPC 定义**: 使用 `.proto` 文件定义 gRPC 接口，通过 `goctl rpc protoc` 生成代码。
- **错误码**: **必须使用自定义错误码（Custom Err Code）返回错误**，不能直接返回底层的 error，统一前后端错误信息体。

## 📜 许可证

本项目为个人/内部开发系统。如需开源请在此处补充 License 说明。
