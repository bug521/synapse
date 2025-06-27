<script setup lang="ts">
import { computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useUserStore } from './store/userStore'
import Layout from './components/Layout.vue'

const route = useRoute()
const userStore = useUserStore()

// 判断是否为认证页面
const isAuthPage = computed(() => {
  return ['Login', 'Register'].includes(route.name as string)
})

// 初始化用户状态
onMounted(() => {
  userStore.initUser()
})
</script>

<template>
  <n-config-provider>
    <n-loading-bar-provider>
      <n-dialog-provider>
        <n-notification-provider>
          <n-message-provider>
            <div class="app-container">
              <!-- 登录/注册页面不显示布局 -->
              <template v-if="isAuthPage">
                <router-view />
              </template>
              
              <!-- 主布局 -->
              <template v-else>
                <Layout>
                  <router-view />
                </Layout>
              </template>
            </div>
          </n-message-provider>
        </n-notification-provider>
      </n-dialog-provider>
    </n-loading-bar-provider>
  </n-config-provider>
</template>

<style scoped>
.app-container {
  height: 100vh;
}
</style>
