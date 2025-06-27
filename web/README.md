# Synapse Web 前端

基于 Vue 3 + TypeScript + Naive UI 构建的现代化管理系统前端。

## 功能特性

### 🎨 现代化界面
- 使用 Naive UI 组件库，提供美观的界面设计
- 响应式布局，支持多种屏幕尺寸
- 深色/浅色主题切换

### 🔐 用户认证
- 用户登录/注册功能
- JWT 令牌认证
- 用户状态管理（Pinia）
- 路由守卫保护

### 📱 完整布局
- 顶部导航栏：网站名称 + 用户信息
- 左侧菜单栏：可折叠的导航菜单
- 主内容区域：路由页面展示
- 面包屑导航

### 🏠 页面功能
- **首页**：欢迎页面，系统概览
- **仪表板**：数据统计和图表展示
- **个人资料**：用户信息管理和安全设置
- **系统设置**：基本设置、通知设置、安全设置

## 技术栈

- **框架**：Vue 3 + TypeScript
- **UI 组件库**：Naive UI
- **状态管理**：Pinia
- **路由**：Vue Router 4
- **构建工具**：Vite
- **包管理**：pnpm

## 项目结构

```
web/
├── src/
│   ├── components/          # 公共组件
│   │   ├── auth/           # 认证相关页面
│   │   ├── home/           # 首页
│   │   ├── dashboard/      # 仪表板
│   │   ├── profile/        # 个人资料
│   │   └── settings/       # 系统设置
│   ├── store/              # 状态管理
│   ├── router/             # 路由配置
│   ├── assets/             # 静态资源
│   ├── App.vue             # 根组件
│   └── main.ts             # 入口文件
├── public/                 # 公共资源
├── package.json            # 依赖配置
├── vite.config.ts          # Vite 配置
└── tsconfig.json           # TypeScript 配置
```

## 快速开始

### 安装依赖

```bash
cd web
pnpm install
```

### 开发模式

```bash
pnpm dev
```

### 构建生产版本

```bash
pnpm build
```

### 预览生产版本

```bash
pnpm preview
```

## 主要功能说明

### 1. 用户认证系统
- 支持用户注册和登录
- 使用 localStorage 持久化用户状态
- 登录状态自动恢复

### 2. 响应式布局
- 顶部固定导航栏
- 左侧可折叠菜单栏
- 主内容区域自适应

### 3. 路由系统
- 基于 Vue Router 4
- 支持路由懒加载
- 认证页面独立布局

### 4. 状态管理
- 使用 Pinia 进行状态管理
- 用户信息全局状态
- 登录状态持久化

## 开发指南

### 添加新页面

1. 在 `src/views/` 下创建新的页面组件
2. 在 `src/router/index.ts` 中添加路由配置
3. 在 `src/App.vue` 中添加菜单项（如需要）

### 使用 Naive UI 组件

项目已配置 Naive UI 的自动导入，可以直接使用组件：

```vue
<template>
  <n-button type="primary">按钮</n-button>
  <n-input placeholder="输入框" />
</template>
```

### 状态管理

使用 Pinia 进行状态管理：

```typescript
import { useUserStore } from '@/store/userStore'

const userStore = useUserStore()
userStore.login(userData)
```

## 配置说明

### Vite 配置
- 配置了 Naive UI 的自动导入
- 支持 TypeScript
- 配置了路径别名

### TypeScript 配置
- 严格的类型检查
- 支持 Vue 3 语法
- 配置了路径映射

## 部署

### 构建
```bash
pnpm build
```

### 部署到静态服务器
将 `dist` 目录下的文件部署到 Web 服务器即可。

## 贡献指南

1. Fork 项目
2. 创建功能分支
3. 提交更改
4. 推送到分支
5. 创建 Pull Request

## 许可证

MIT License
