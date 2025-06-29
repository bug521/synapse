<template>
  <div class="login-container">
    <n-card title="登录" class="login-card">
      <n-form
        ref="formRef"
        :model="formValue"
        :rules="rules"
        label-placement="left"
        label-width="auto"
        require-mark-placement="right-hanging"
        size="medium"
      >
        <n-form-item label="用户名" path="username">
          <n-input v-model:value="formValue.username" placeholder="请输入用户名" />
        </n-form-item>
        <n-form-item label="密码" path="password">
          <n-input
            v-model:value="formValue.password"
            type="password"
            placeholder="请输入密码"
            show-password-on="click"
          />
        </n-form-item>
        <n-form-item>
          <n-button type="primary" block @click="handleLogin" :loading="loading">
            登录
          </n-button>
        </n-form-item>
        <div class="register-link">
          <span>还没有账号？</span>
          <n-button text type="primary" @click="$router.push('/register')">
            立即注册
          </n-button>
        </div>
      </n-form>
    </n-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { useMessage } from 'naive-ui'
import { useUserStore } from '../../store/userStore'
import api from '../../api'

const router = useRouter()
const message = useMessage()
const userStore = useUserStore()

const formRef = ref()
const loading = ref(false)

const formValue = reactive({
  username: '',
  password: ''
})

const rules = {
  username: {
    required: true,
    message: '请输入用户名',
    trigger: 'blur'
  },
  password: {
    required: true,
    message: '请输入密码',
    trigger: 'blur'
  }
}

const handleLogin = () => {
  formRef.value?.validate(async (errors: any) => {
    if (!errors) {
      loading.value = true
      try {
        const response = await api.post('/api/login', formValue)
        const { user, token } = response.data || response
        
        userStore.login(user, token)
        message.success('登录成功')
        router.push('/')
      } catch (error: any) {
        message.error('登录失败: ' + (error.response?.data?.message || error.message || '未知错误'))
      } finally {
        loading.value = false
      }
    }
  })
}
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.login-card {
  width: 400px;
  max-width: 90vw;
}

.register-link {
  text-align: center;
  margin-top: 16px;
}
</style> 