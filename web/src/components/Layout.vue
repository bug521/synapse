<template>
  <n-layout has-sider>
    <!-- 侧边栏 -->
    <Sidebar v-model:collapsed="collapsed" />
    
    <n-layout>
      <!-- 顶部导航栏 -->
      <Header v-model:collapsed="collapsed" />
      
      <!-- 主内容区域 -->
      <n-layout-content class="main-content">
        <slot />
      </n-layout-content>
    </n-layout>
  </n-layout>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Header from './Header.vue'
import Sidebar from './Sidebar.vue'

// 侧边栏状态
const collapsed = ref(false)
</script>

<style scoped>
.main-content {
  background: #f5f5f5;
  min-height: calc(100vh - 64px);
}
</style> 