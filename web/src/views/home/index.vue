<template>
  <div class="home-container">
    <n-card title="欢迎来到 Synapse">
      <div class="welcome-content">
        <n-icon size="64" color="#18a058">
          <svg viewBox="0 0 24 24">
            <path fill="currentColor" d="M12 2L2 7l10 5 10-5-10-5zM2 17l10 5 10-5M2 12l10 5 10-5"/>
          </svg>
        </n-icon>
        <h2>欢迎使用 Synapse 系统</h2>
        <p>这是一个基于 Vue 3 + Naive UI 构建的现代化管理系统</p>
        <n-space>
          <n-button type="primary" @click="$router.push('/dashboard')">
            进入仪表板
          </n-button>
          <n-button @click="$router.push('/profile')">
            个人资料
          </n-button>
        </n-space>
      </div>
    </n-card>
  </div>
</template>

<script setup lang="ts">
// 首页组件
</script>

<style scoped>
.home-container {
  padding: 24px;
}

.welcome-content {
  text-align: center;
  padding: 40px 0;
}

.welcome-content h2 {
  margin: 24px 0 16px;
  color: #18a058;
}

.welcome-content p {
  margin-bottom: 32px;
  color: #666;
  font-size: 16px;
}
</style>
