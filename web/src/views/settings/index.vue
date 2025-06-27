<template>
  <div class="settings-container">
    <n-card title="系统设置">
      <n-tabs type="line" animated>
        <!-- 基本设置 -->
        <n-tab-pane name="basic" tab="基本设置">
          <n-form
            ref="basicFormRef"
            :model="basicForm"
            label-placement="left"
            label-width="auto"
            require-mark-placement="right-hanging"
          >
            <n-form-item label="网站名称">
              <n-input v-model:value="basicForm.siteName" placeholder="请输入网站名称" />
            </n-form-item>
            
            <n-form-item label="网站描述">
              <n-input
                v-model:value="basicForm.siteDescription"
                type="textarea"
                placeholder="请输入网站描述"
                :autosize="{ minRows: 3, maxRows: 5 }"
              />
            </n-form-item>
            
            <n-form-item label="主题模式">
              <n-select
                v-model:value="basicForm.theme"
                :options="themeOptions"
                placeholder="请选择主题模式"
              />
            </n-form-item>
            
            <n-form-item>
              <n-button type="primary" @click="handleSaveBasic" :loading="basicLoading">
                保存设置
              </n-button>
            </n-form-item>
          </n-form>
        </n-tab-pane>
        
        <!-- 通知设置 -->
        <n-tab-pane name="notification" tab="通知设置">
          <n-space vertical>
            <n-card title="邮件通知" size="small">
              <n-space vertical>
                <n-switch v-model:value="notificationSettings.emailEnabled" />
                <span>启用邮件通知</span>
              </n-space>
            </n-card>
            
            <n-card title="系统通知" size="small">
              <n-space vertical>
                <n-switch v-model:value="notificationSettings.systemEnabled" />
                <span>启用系统通知</span>
              </n-space>
            </n-card>
            
            <n-card title="推送通知" size="small">
              <n-space vertical>
                <n-switch v-model:value="notificationSettings.pushEnabled" />
                <span>启用推送通知</span>
              </n-space>
            </n-card>
            
            <n-button type="primary" @click="handleSaveNotification" :loading="notificationLoading">
              保存通知设置
            </n-button>
          </n-space>
        </n-tab-pane>
        
        <!-- 安全设置 -->
        <n-tab-pane name="security" tab="安全设置">
          <n-space vertical>
            <n-card title="登录安全" size="small">
              <n-space vertical>
                <n-switch v-model:value="securitySettings.twoFactorEnabled" />
                <span>启用两步验证</span>
              </n-space>
            </n-card>
            
            <n-card title="会话管理" size="small">
              <n-space vertical>
                <n-input-number
                  v-model:value="securitySettings.sessionTimeout"
                  :min="1"
                  :max="24"
                  placeholder="会话超时时间（小时）"
                />
                <span>会话超时时间</span>
              </n-space>
            </n-card>
            
            <n-card title="密码策略" size="small">
              <n-space vertical>
                <n-input-number
                  v-model:value="securitySettings.passwordMinLength"
                  :min="6"
                  :max="20"
                  placeholder="最小密码长度"
                />
                <span>最小密码长度</span>
              </n-space>
            </n-card>
            
            <n-button type="primary" @click="handleSaveSecurity" :loading="securityLoading">
              保存安全设置
            </n-button>
          </n-space>
        </n-tab-pane>
      </n-tabs>
    </n-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useMessage } from 'naive-ui'

const message = useMessage()

// 基本设置表单
const basicFormRef = ref()
const basicLoading = ref(false)
const basicForm = reactive({
  siteName: 'Synapse',
  siteDescription: '一个现代化的管理系统',
  theme: 'light'
})

const themeOptions = [
  { label: '浅色主题', value: 'light' },
  { label: '深色主题', value: 'dark' },
  { label: '自动切换', value: 'auto' }
]

// 通知设置
const notificationLoading = ref(false)
const notificationSettings = reactive({
  emailEnabled: true,
  systemEnabled: true,
  pushEnabled: false
})

// 安全设置
const securityLoading = ref(false)
const securitySettings = reactive({
  twoFactorEnabled: false,
  sessionTimeout: 2,
  passwordMinLength: 8
})

// 保存基本设置
const handleSaveBasic = async () => {
  basicLoading.value = true
  try {
    await new Promise(resolve => setTimeout(resolve, 1000))
    message.success('基本设置保存成功')
  } catch (error) {
    message.error('保存失败')
  } finally {
    basicLoading.value = false
  }
}

// 保存通知设置
const handleSaveNotification = async () => {
  notificationLoading.value = true
  try {
    await new Promise(resolve => setTimeout(resolve, 1000))
    message.success('通知设置保存成功')
  } catch (error) {
    message.error('保存失败')
  } finally {
    notificationLoading.value = false
  }
}

// 保存安全设置
const handleSaveSecurity = async () => {
  securityLoading.value = true
  try {
    await new Promise(resolve => setTimeout(resolve, 1000))
    message.success('安全设置保存成功')
  } catch (error) {
    message.error('保存失败')
  } finally {
    securityLoading.value = false
  }
}
</script>

<style scoped>
.settings-container {
  padding: 24px;
}
</style> 