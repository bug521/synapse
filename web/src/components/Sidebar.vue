<template>
  <n-layout-sider
    bordered
    collapse-mode="width"
    :collapsed-width="64"
    :width="240"
    :collapsed="collapsed"
    show-trigger
    @collapse="handleCollapse"
    @expand="handleExpand"
  >
    <div class="sidebar-header">
      <n-icon size="32" color="#18a058">
        <svg viewBox="0 0 24 24">
          <path fill="currentColor" d="M12 2L2 7l10 5 10-5-10-5zM2 17l10 5 10-5M2 12l10 5 10-5"/>
        </svg>
      </n-icon>
      <span v-show="!collapsed" class="logo-text">Synapse</span>
    </div>
    
    <n-menu
      :collapsed="collapsed"
      :collapsed-width="64"
      :collapsed-icon-size="22"
      :options="menuOptions"
      :value="activeKey"
      @update:value="handleMenuClick"
    />
  </n-layout-sider>
</template>

<script setup lang="ts">
import { computed, h } from 'vue'
import { useRoute, useRouter } from 'vue-router'

interface Props {
  collapsed: boolean
}

interface Emits {
  (e: 'update:collapsed', value: boolean): void
}

const props = defineProps<Props>()
const emit = defineEmits<Emits>()

const route = useRoute()
const router = useRouter()

// 菜单配置
const menuOptions = [
  {
    label: '首页',
    key: 'home',
    icon: () => h('svg', { viewBox: '0 0 24 24' }, [
      h('path', { fill: 'currentColor', d: 'M10 20v-6h4v6h5v-8h3L12 3 2 12h3v8z' })
    ])
  },
  {
    label: '仪表板',
    key: 'dashboard',
    icon: () => h('svg', { viewBox: '0 0 24 24' }, [
      h('path', { fill: 'currentColor', d: 'M3 13h8V3H3v10zm0 8h8v-6H3v6zm10 0h8V11h-8v10zm0-18v6h8V3h-8z' })
    ])
  },
  {
    label: '通知通道',
    key: 'channels',
    icon: () => h('svg', { viewBox: '0 0 24 24' }, [
      h('path', { fill: 'currentColor', d: 'M20 4H4c-1.1 0-1.99.9-1.99 2L2 18c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V6c0-1.1-.9-2-2-2zm-5 14H4v-4h11v4zm0-5H4V9h11v4zm5 5h-4V9h4v9z' })
    ])
  },
  {
    label: '消息主题',
    key: 'topics',
    icon: () => h('svg', { viewBox: '0 0 24 24' }, [
      h('path', { fill: 'currentColor', d: 'M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z' })
    ])
  },
  {
    label: '个人资料',
    key: 'profile',
    icon: () => h('svg', { viewBox: '0 0 24 24' }, [
      h('path', { fill: 'currentColor', d: 'M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm0 2c-2.67 0-8 1.34-8 4v2h16v-2c0-2.66-5.33-4-8-4z' })
    ])
  }
]

// 当前激活的菜单项
const activeKey = computed(() => {
  const routeMap: Record<string, string> = {
    'Home': 'home',
    'Dashboard': 'dashboard',
    'Channels': 'channels',
    'Topics': 'topics',
    'TopicDetail': 'topics',
    'Profile': 'profile'
  }
  return routeMap[route.name as string] || 'home'
})

// 处理菜单点击
const handleMenuClick = (key: string) => {
  const routeMap: Record<string, string> = {
    'home': '/',
    'dashboard': '/dashboard',
    'channels': '/channels',
    'topics': '/topics',
    'profile': '/profile'
  }
  router.push(routeMap[key])
}

// 处理折叠
const handleCollapse = () => {
  emit('update:collapsed', true)
}

// 处理展开
const handleExpand = () => {
  emit('update:collapsed', false)
}
</script>

<style scoped>
.sidebar-header {
  height: 64px;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 12px;
  border-bottom: 1px solid #f0f0f0;
  padding: 0 16px;
}

.logo-text {
  font-size: 18px;
  font-weight: bold;
  color: #18a058;
}
</style> 