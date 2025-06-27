<template>
  <div class="profile-container">
    <n-card title="个人资料">
      <n-grid :cols="24" :x-gap="24">
        <!-- 左侧：头像和基本信息 -->
        <n-grid-item :span="8">
          <div class="profile-avatar">
            <n-avatar
              round
              size="large"
              :src="userStore.user?.avatar"
              :fallback-src="defaultAvatar"
            />
            <n-button class="upload-btn" @click="handleUploadAvatar">
              更换头像
            </n-button>
          </div>
          
          <n-divider />
          
          <div class="user-info">
            <h3>{{ userStore.user?.username }}</h3>
            <p>{{ userStore.user?.email }}</p>
            <n-tag type="success" size="small">已认证</n-tag>
          </div>
        </n-grid-item>
        
        <!-- 右侧：详细信息表单 -->
        <n-grid-item :span="16">
          <n-form
            ref="formRef"
            :model="formValue"
            :rules="rules"
            label-placement="left"
            label-width="auto"
            require-mark-placement="right-hanging"
          >
            <n-form-item label="用户名" path="username">
              <n-input v-model:value="formValue.username" placeholder="请输入用户名" />
            </n-form-item>
            
            <n-form-item label="邮箱" path="email">
              <n-input v-model:value="formValue.email" placeholder="请输入邮箱" />
            </n-form-item>
            
            <n-form-item label="昵称" path="nickname">
              <n-input v-model:value="formValue.nickname" placeholder="请输入昵称" />
            </n-form-item>
            
            <n-form-item label="手机号" path="phone">
              <n-input v-model:value="formValue.phone" placeholder="请输入手机号" />
            </n-form-item>
            
            <n-form-item label="个人简介" path="bio">
              <n-input
                v-model:value="formValue.bio"
                type="textarea"
                placeholder="请输入个人简介"
                :autosize="{ minRows: 3, maxRows: 5 }"
              />
            </n-form-item>
            
            <n-form-item>
              <n-space>
                <n-button type="primary" @click="handleSave" :loading="loading">
                  保存修改
                </n-button>
                <n-button @click="handleReset">
                  重置
                </n-button>
              </n-space>
            </n-form-item>
          </n-form>
        </n-grid-item>
      </n-grid>
    </n-card>
    
    <!-- 安全设置 -->
    <n-card title="安全设置" class="security-card">
      <n-list>
        <n-list-item>
          <n-thing title="修改密码" description="定期更换密码以确保账户安全">
            <template #action>
              <n-button size="small" @click="showChangePassword = true">
                修改
              </n-button>
            </template>
          </n-thing>
        </n-list-item>
        
        <n-list-item>
          <n-thing title="两步验证" description="启用两步验证以提高账户安全性">
            <template #action>
              <n-switch v-model:value="twoFactorEnabled" />
            </template>
          </n-thing>
        </n-list-item>
        
        <n-list-item>
          <n-thing title="登录设备" description="查看和管理已登录的设备">
            <template #action>
              <n-button size="small">
                管理
              </n-button>
            </template>
          </n-thing>
        </n-list-item>
      </n-list>
    </n-card>
    
    <!-- 修改密码对话框 -->
    <n-modal v-model:show="showChangePassword" preset="card" title="修改密码" style="width: 400px">
      <n-form
        ref="passwordFormRef"
        :model="passwordForm"
        :rules="passwordRules"
        label-placement="left"
        label-width="auto"
      >
        <n-form-item label="当前密码" path="currentPassword">
          <n-input
            v-model:value="passwordForm.currentPassword"
            type="password"
            placeholder="请输入当前密码"
            show-password-on="click"
          />
        </n-form-item>
        
        <n-form-item label="新密码" path="newPassword">
          <n-input
            v-model:value="passwordForm.newPassword"
            type="password"
            placeholder="请输入新密码"
            show-password-on="click"
          />
        </n-form-item>
        
        <n-form-item label="确认新密码" path="confirmPassword">
          <n-input
            v-model:value="passwordForm.confirmPassword"
            type="password"
            placeholder="请再次输入新密码"
            show-password-on="click"
          />
        </n-form-item>
      </n-form>
      
      <template #action>
        <n-space justify="end">
          <n-button @click="showChangePassword = false">
            取消
          </n-button>
          <n-button type="primary" @click="handleChangePassword" :loading="passwordLoading">
            确认修改
          </n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { useMessage } from 'naive-ui'
import { useUserStore } from '@/store/userStore'

const message = useMessage()
const userStore = useUserStore()

const formRef = ref()
const passwordFormRef = ref()
const loading = ref(false)
const passwordLoading = ref(false)
const showChangePassword = ref(false)
const twoFactorEnabled = ref(false)

const defaultAvatar = 'https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg'

const formValue = reactive({
  username: '',
  email: '',
  nickname: '',
  phone: '',
  bio: ''
})

const passwordForm = reactive({
  currentPassword: '',
  newPassword: '',
  confirmPassword: ''
})

const rules = {
  username: {
    required: true,
    message: '请输入用户名',
    trigger: 'blur'
  },
  email: {
    required: true,
    message: '请输入邮箱',
    trigger: 'blur',
    pattern: /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  }
}

const passwordRules = {
  currentPassword: {
    required: true,
    message: '请输入当前密码',
    trigger: 'blur'
  },
  newPassword: {
    required: true,
    message: '请输入新密码',
    trigger: 'blur',
    min: 6
  },
  confirmPassword: {
    required: true,
    message: '请确认新密码',
    trigger: 'blur',
    validator: (rule: any, value: string) => {
      if (value !== passwordForm.newPassword) {
        return new Error('两次输入的密码不一致')
      }
    }
  }
}

const handleUploadAvatar = () => {
  message.info('头像上传功能开发中...')
}

const handleSave = () => {
  formRef.value?.validate(async (errors: any) => {
    if (!errors) {
      loading.value = true
      try {
        // 模拟保存请求
        await new Promise(resolve => setTimeout(resolve, 1000))
        message.success('保存成功')
      } catch (error) {
        message.error('保存失败')
      } finally {
        loading.value = false
      }
    }
  })
}

const handleReset = () => {
  // 重置表单数据
  Object.assign(formValue, {
    username: userStore.user?.username || '',
    email: userStore.user?.email || '',
    nickname: '',
    phone: '',
    bio: ''
  })
  message.info('表单已重置')
}

const handleChangePassword = () => {
  passwordFormRef.value?.validate(async (errors: any) => {
    if (!errors) {
      passwordLoading.value = true
      try {
        // 模拟修改密码请求
        await new Promise(resolve => setTimeout(resolve, 1000))
        message.success('密码修改成功')
        showChangePassword.value = false
        // 清空表单
        Object.assign(passwordForm, {
          currentPassword: '',
          newPassword: '',
          confirmPassword: ''
        })
      } catch (error) {
        message.error('密码修改失败')
      } finally {
        passwordLoading.value = false
      }
    }
  })
}

onMounted(() => {
  // 初始化表单数据
  if (userStore.user) {
    formValue.username = userStore.user.username
    formValue.email = userStore.user.email
  }
})
</script>

<style scoped>
.profile-container {
  padding: 24px;
}

.profile-avatar {
  text-align: center;
  margin-bottom: 24px;
}

.upload-btn {
  margin-top: 16px;
}

.user-info {
  text-align: center;
}

.user-info h3 {
  margin: 16px 0 8px;
  color: #333;
}

.user-info p {
  margin: 8px 0;
  color: #666;
}

.security-card {
  margin-top: 24px;
}
</style> 