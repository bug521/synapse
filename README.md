# Synapse

![CI/CD Status](https://img.shields.io/github/actions/workflow/status/YOUR_USERNAME/synapse/go.yml?branch=main&style=for-the-badge)
![Go Report Card](https://goreportcard.com/badge/github.com/YOUR_USERNAME/synapse?style=for-the-badge)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/YOUR_USERNAME/synapse?style=for-the-badge)
![GitHub](https://img.shields.io/github/license/YOUR_USERNAME/synapse?style=for-the-badge)

一个灵活、可自托管的Webhook到通知转发服务。Synapse接收来自任何来源的webhook，并智能地将它们转发到您配置的通道，如Telegram、Slack等。

---

## 🤔 什么是 Synapse？

在现代开发中，我们依赖无数通过webhook通信的服务（例如Git推送、用户注册、监控警报）。问题是每个服务都有自己的JSON格式。Synapse作为一个强大的中间件，捕获这些webhook，允许您提取需要的数据，将其格式化为人类可读的消息，并路由到一个或多个通知通道。

它是您的个人、集中化通知中心。

## ✨ 主要功能

* **基于项目的Webhook**: 为每个消息源创建隔离的"项目"，每个项目都有自己唯一的webhook URL。
* **多通道支持**: 同时向多个通道转发消息。易于扩展以支持新服务（Telegram、Slack、Webhooks等）。
* **高级消息模板**: 不仅仅是转发丑陋的JSON！使用Go的`text/template`语法创建美观、自定义的消息格式。
* **动态变量提取**: 使用`gjson`路径语法从传入JSON负载的任何部分提取数据。无需为每个webhook格式编写自定义代码。
* **灵活的路由策略**:
    * **发送给所有**: 向所有配置的通道广播消息。
    * **故障转移**: 按优先级顺序发送消息，一旦一个通道成功就停止。确保可传递性的完美选择。
* **可配置的执行模式**:
    * **异步（默认）**: 立即响应webhook源并在后台处理消息以获得最大性能。
    * **同步**: 等待消息发送完成并将最终传递结果返回给调用者。
* **消息历史和日志**: 接收消息及其传递状态的完整历史，便于调试和审计。
* **现代Web UI**: 使用Vue.js和Naive UI构建的干净直观的界面，用于管理您的项目、通道和路由。

## 架构图

```mermaid
graph TD
    A[外部服务<br/>(例如GitHub、Stripe)] -- JSON负载 --> B{Synapse Webhook端点<br/>(Gin)};
    B -- 1. 保存消息（待处理） --> C[MySQL数据库];
    B -- 2. 推送消息ID --> D{异步队列<br/>(Go Channel)};
    E[分发器工作器] -- 3. 从队列消费 --> D;
    E -- 4. 获取规则和消息 --> C;
    E -- 5. 渲染模板并发送 --> F[通知器];
    F -- Telegram --> G[Telegram API];
    F -- Slack等 --> H[...];
    
    U[用户] <--> X[Web UI<br/>(Vue.js + Naive UI)];
    X <--> Y{Synapse API<br/>(Gin)};
    Y <--> C;
```

## 🛠️ 技术栈

  * **后端**: Go, Gin, GORM, Viper, JWT
  * **前端**: Vue.js (v3), Pinia, Vue Router, Naive UI, Axios
  * **数据库**: MySQL
  * **容器化**: Docker, Docker Compose

## 🚀 快速开始

按照以下步骤在本地启动您自己的Synapse实例。

### 前置要求

  * [Git](https://git-scm.com/)
  * [Go](https://go.dev/) (v1.21+)
  * [Node.js](https://nodejs.org/) (v18+)
  * [Docker](https://www.docker.com/) & [Docker Compose](https://docs.docker.com/compose/)

### 1. 安装

克隆仓库：

```bash
git clone [https://github.com/YOUR_USERNAME/synapse.git](https://github.com/YOUR_USERNAME/synapse.git)
cd synapse
```

### 2. 配置

复制示例配置文件。

```bash
cp config/config.yaml.example config/config.yaml
```

现在编辑`config/config.yaml`并填写您的详细信息，特别是`database`连接设置和强`jwt_secret`。

### 3. 数据库设置

确保您的MySQL服务器正在运行。然后连接到它并执行SQL脚本来创建所有必要的表。

```bash
# 使用MySQL客户端，执行此文件的内容：
# db/schema.sql
```

### 4. 运行应用程序

运行整个堆栈（后端、前端和数据库）的最简单方法是使用Docker Compose。

```bash
docker-compose up -d --build
```

就是这样！

  * 前端将在`http://localhost:5173`可用
  * 后端API将在`http://localhost:8080`可用

### 本地开发（不使用Docker）

如果您更喜欢直接在主机机器上运行服务：

```bash
# 终端1: 运行后端
go run main.go

# 终端2: 运行前端
cd web
npm install
npm run dev
```

## 📖 API文档

### 认证

所有受保护的API都需要JWT认证。在请求头中包含：

```
Authorization: Bearer <your-jwt-token>
```

### 用户管理

#### 注册用户
```http
POST /api/register
Content-Type: application/json

{
  "username": "testuser",
  "password": "password123",
  "email": "test@example.com"
}
```

#### 用户登录
```http
POST /api/login
Content-Type: application/json

{
  "username": "testuser",
  "password": "password123"
}
```

### 通道管理

#### 创建通道
```http
POST /api/channels
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "我的Telegram通道",
  "type": "telegram",
  "credentials": {
    "bot_token": "your-bot-token",
    "chat_id": "your-chat-id",
    "parse_mode": "HTML"
  }
}
```

#### 获取通道列表
```http
GET /api/channels
Authorization: Bearer <token>
```

#### 更新通道
```http
PUT /api/channels/{id}
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "更新的通道名称",
  "type": "telegram",
  "credentials": {
    "bot_token": "new-bot-token",
    "chat_id": "new-chat-id"
  }
}
```

### 主题管理

#### 创建主题
```http
POST /api/topics
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "GitHub通知",
  "sending_strategy": "all",
  "execution_mode": "async",
  "description": "接收GitHub事件通知"
}
```

#### 获取主题列表
```http
GET /api/topics
Authorization: Bearer <token>
```

#### 重新生成Webhook Key
```http
POST /api/topics/{id}/regenerate-key
Authorization: Bearer <token>
```

### 路由管理

#### 创建路由
```http
POST /api/routings
Authorization: Bearer <token>
Content-Type: application/json

{
  "topic_id": 1,
  "channel_id": 1,
  "priority": 1,
  "variable_mappings": {
    "title": "repository.name",
    "action": "action"
  },
  "message_template": "仓库 {{.title}} 有新活动: {{.action}}"
}
```

#### 获取主题的路由
```http
GET /api/topics/{topic_id}/routings
Authorization: Bearer <token>
```

### Webhook接收

#### 发送Webhook
```http
POST /webhook/{webhook_key}
Content-Type: application/json

{
  "repository": {
    "name": "my-repo"
  },
  "action": "push",
  "user": "john_doe"
}
```

#### 获取Webhook信息
```http
GET /webhook/{webhook_key}/info
```

## 使用示例

1. 导航到`http://localhost:5173`。

2. **注册**一个新用户并登录。

3. 转到**通道**页面并添加一个新通道（例如，您的Telegram Bot Token和Chat ID）。

4. 转到**项目**页面并创建一个新项目。将为您生成一个唯一的webhook URL。

5. 在项目的管理页面上，创建一个**路由规则**来链接您的项目到您刚刚创建的通道。

6. 使用像`curl`或Postman这样的工具向您的webhook URL发送测试负载：

    ```bash
    curl -X POST \
      -H "Content-Type: application/json" \
      -d '{ "user": "John Doe", "action": "signed_up", "level": "critical" }' \
      http://localhost:8080/webhook/YOUR_UNIQUE_WEBHOOK_KEY
    ```

7. 检查您的通知通道。您应该看到消息！要自定义它，编辑您的路由规则并添加消息模板。

## 🔧 开发

### 项目结构

```
synapse/
├── config/                 # 配置文件
├── db/                     # 数据库脚本
├── internal/               # 内部包
│   ├── config/            # 配置管理
│   ├── controller/        # HTTP控制器
│   ├── middleware/        # 中间件
│   ├── model/             # 数据模型
│   ├── repository/        # 数据访问层
│   ├── router/            # 路由配置
│   ├── service/           # 业务逻辑层
│   └── utils/             # 工具函数
├── pkg/                   # 公共包
├── storage/               # 存储目录
├── web/                   # 前端代码
├── main.go               # 主程序入口
└── README.md             # 项目文档
```

### 添加新的通知通道

1. 在`internal/model/channel.go`中添加新的配置结构
2. 在`internal/service/channel.go`中添加验证逻辑
3. 在`internal/service/message.go`中实现发送逻辑
4. 更新前端UI以支持新通道类型

### 运行测试

```bash
go test ./...
```

## 🤝 贡献

欢迎贡献、问题和功能请求！请随时查看[issues页面](https://www.google.com/search?q=https://github.com/YOUR_USERNAME/synapse/issues)。

## 📄 许可证

本项目采用MIT许可证。有关详细信息，请参阅[LICENSE](https://www.google.com/search?q=LICENSE)文件。