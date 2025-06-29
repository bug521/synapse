# Synapse Web 前端

这是 Synapse 消息转发服务的 Web 管理界面，基于 Vue 3 + TypeScript + Naive UI 构建。

## 功能特性

- 🎨 现代化的用户界面，基于 Naive UI 组件库
- 📱 响应式设计，支持桌面和移动设备
- 🔐 用户认证和授权管理
- 📡 通知通道管理（Telegram、Email、Slack、Webhook）
- 🎯 消息主题管理和路由配置
- 🔑 Webhook Key 管理和复制
- ⚡ 实时数据更新和状态反馈

## 技术栈

- **框架**: Vue 3 + TypeScript
- **UI 组件库**: Naive UI
- **状态管理**: Pinia
- **路由**: Vue Router 4
- **HTTP 客户端**: Axios
- **构建工具**: Vite
- **包管理器**: pnpm

## 快速开始

### 环境要求

- Node.js >= 16
- pnpm >= 7

### 安装依赖

```bash
pnpm install
```

### 环境配置

复制环境配置文件：

```bash
cp env.example .env
```

根据需要修改 `.env` 文件中的配置：

```env
# API基础URL
VITE_API_BASE_URL=http://localhost:8080

# 应用标题
VITE_APP_TITLE=Synapse

# 开发模式
VITE_DEV_MODE=true
```

### 开发模式

```bash
pnpm dev
```

访问 http://localhost:5173 查看应用。

### 构建生产版本

```bash
pnpm build
```

构建产物将输出到 `dist` 目录。

### 预览生产版本

```bash
pnpm preview
```

## 项目结构

```
src/
├── api/                 # API 接口定义
│   └── index.ts        # API 客户端和类型定义
├── assets/             # 静态资源
├── components/         # 公共组件
│   ├── Header.vue     # 顶部导航栏
│   ├── Layout.vue     # 主布局组件
│   └── Sidebar.vue    # 侧边栏导航
├── router/            # 路由配置
│   └── index.ts       # 路由定义
├── store/             # 状态管理
│   └── userStore.ts   # 用户状态
├── views/             # 页面组件
│   ├── auth/          # 认证相关页面
│   │   ├── Login.vue  # 登录页面
│   │   └── Register.vue # 注册页面
│   ├── channels/      # 通道管理
│   │   └── index.vue  # 通道列表页面
│   └── topics/        # 主题管理
│       ├── index.vue  # 主题列表页面
│       └── detail.vue # 主题详情页面
├── App.vue            # 根组件
└── main.ts           # 应用入口
```

## 主要功能

### 1. 用户认证

- 用户注册和登录
- JWT Token 管理
- 自动登录状态恢复

### 2. 通知通道管理

- 创建和管理不同类型的通知通道
- 支持 Telegram、Email、Slack、Webhook
- 通道配置和凭证管理
- 通道列表和搜索

### 3. 消息主题管理

- 创建和管理消息主题
- Webhook Key 生成和管理
- 发送策略和执行模式配置
- 主题详情查看

### 4. 路由配置

- 为主题配置通知路由
- 通道优先级设置
- 变量映射配置
- 消息模板定制

## API 集成

前端通过 RESTful API 与后端服务通信，主要接口包括：

- **用户管理**: `/api/register`, `/api/login`, `/api/profile`
- **通道管理**: `/api/channels`
- **主题管理**: `/api/topics`
- **路由管理**: `/api/routings`
- **Webhook**: `/webhook/:webhook_key`

## 开发指南

### 添加新页面

1. 在 `src/views/` 下创建页面组件
2. 在 `src/router/index.ts` 中添加路由配置
3. 在 `src/components/Sidebar.vue` 中添加菜单项

### 添加新 API

1. 在 `src/api/index.ts` 中定义接口类型
2. 添加对应的 API 方法
3. 在页面组件中调用 API

### 组件开发

- 使用 Naive UI 组件库
- 遵循 Vue 3 Composition API 规范
- 使用 TypeScript 进行类型检查

## 部署

### Docker 部署

项目包含 Dockerfile，可以使用 Docker 进行部署：

```bash
# 构建镜像
docker build -t synapse-web .

# 运行容器
docker run -p 80:80 synapse-web
```

### Nginx 部署

构建后的静态文件可以部署到 Nginx：

```nginx
server {
    listen 80;
    server_name your-domain.com;
    root /path/to/dist;
    index index.html;
    
    location / {
        try_files $uri $uri/ /index.html;
    }
    
    location /api {
        proxy_pass http://backend:8080;
    }
}
```

## 贡献指南

1. Fork 项目
2. 创建功能分支
3. 提交更改
4. 推送到分支
5. 创建 Pull Request

## 许可证

MIT License
