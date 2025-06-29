<template>
  <div class="channels-page">
    <div class="page-header">
      <h1>通知通道管理</h1>
      <n-button type="primary" @click="showCreateModal = true">
        <template #icon>
          <n-icon>+</n-icon>
        </template>
        添加通道
      </n-button>
    </div>

    <!-- 通道列表 -->
    <n-card>
      <n-data-table
        :columns="columns"
        :data="channels"
        :loading="loading"
        :pagination="pagination"
        :bordered="false"
      />
    </n-card>

    <!-- 创建/编辑通道模态框 -->
    <n-modal v-model:show="showCreateModal" preset="card" style="width: 600px">
      <template #header>
        <h3>{{ editingChannel ? '编辑通道' : '添加通道' }}</h3>
      </template>
      
      <n-form
        ref="formRef"
        :model="formData"
        :rules="rules"
        label-placement="left"
        label-width="auto"
        require-mark-placement="right-hanging"
      >
        <n-form-item label="通道名称" path="name">
          <n-input v-model:value="formData.name" placeholder="请输入通道名称" />
        </n-form-item>
        
        <n-form-item label="通道类型" path="type">
          <n-select
            v-model:value="formData.type"
            :options="channelTypeOptions"
            placeholder="请选择通道类型"
            @update:value="handleTypeChange"
          />
        </n-form-item>
        
        <!-- Telegram配置 -->
        <template v-if="formData.type === 'telegram'">
          <n-form-item label="Bot Token" path="credentials.bot_token">
            <n-input
              v-model:value="formData.credentials.bot_token"
              placeholder="请输入Telegram Bot Token"
              type="password"
            />
          </n-form-item>
          <n-form-item label="Chat ID" path="credentials.chat_id">
            <n-input
              v-model:value="formData.credentials.chat_id"
              placeholder="请输入Chat ID"
            />
          </n-form-item>
          <n-form-item label="解析模式" path="credentials.parse_mode">
            <n-select
              v-model:value="formData.credentials.parse_mode"
              :options="parseModeOptions"
              placeholder="请选择解析模式"
            />
          </n-form-item>
        </template>
        
        <!-- Email配置 -->
        <template v-if="formData.type === 'email'">
          <n-form-item label="SMTP主机" path="credentials.smtp_host">
            <n-input
              v-model:value="formData.credentials.smtp_host"
              placeholder="例如: smtp.gmail.com"
            />
          </n-form-item>
          <n-form-item label="SMTP端口" path="credentials.smtp_port">
            <n-input-number
              v-model:value="formData.credentials.smtp_port"
              placeholder="例如: 587"
              :min="1"
              :max="65535"
            />
          </n-form-item>
          <n-form-item label="用户名" path="credentials.smtp_username">
            <n-input
              v-model:value="formData.credentials.smtp_username"
              placeholder="请输入邮箱用户名"
            />
          </n-form-item>
          <n-form-item label="密码" path="credentials.smtp_password">
            <n-input
              v-model:value="formData.credentials.smtp_password"
              placeholder="请输入邮箱密码"
              type="password"
            />
          </n-form-item>
          <n-form-item label="发件人" path="credentials.sender">
            <n-input
              v-model:value="formData.credentials.sender"
              placeholder="请输入发件人邮箱"
            />
          </n-form-item>
        </template>
      </n-form>
      
      <template #footer>
        <div class="modal-footer">
          <n-button @click="showCreateModal = false">取消</n-button>
          <n-button type="primary" :loading="submitting" @click="handleSubmit">
            {{ editingChannel ? '更新' : '创建' }}
          </n-button>
        </div>
      </template>
    </n-modal>

    <!-- 删除确认模态框 -->
    <n-modal v-model:show="showDeleteModal" preset="card" style="width: 400px">
      <template #header>
        <h3>确认删除</h3>
      </template>
      
      <p>确定要删除通道 "{{ deletingChannel?.name }}" 吗？此操作不可恢复。</p>
      
      <template #footer>
        <div class="modal-footer">
          <n-button @click="showDeleteModal = false">取消</n-button>
          <n-button type="error" :loading="deleting" @click="handleDelete">
            删除
          </n-button>
        </div>
      </template>
    </n-modal>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, h } from 'vue'
import { useMessage } from 'naive-ui'
import { channelApi, type Channel, type CreateChannelRequest } from '../../api'
import { NButton, NIcon } from 'naive-ui'
const message = useMessage()

// 响应式数据
const loading = ref(false)
const submitting = ref(false)
const deleting = ref(false)
const channels = ref<Channel[]>([])
const showCreateModal = ref(false)
const showDeleteModal = ref(false)
const editingChannel = ref<Channel | null>(null)
const deletingChannel = ref<Channel | null>(null)
const formRef = ref()

// 表单数据
const formData = reactive<CreateChannelRequest>({
  name: '',
  type: '',
  credentials: {}
})

// 通道类型选项
const channelTypeOptions = [
  { label: 'Telegram', value: 'telegram' },
  { label: 'Email', value: 'email' },
  { label: 'Slack', value: 'slack' },
  { label: 'Webhook', value: 'webhook' }
]

// 解析模式选项
const parseModeOptions = [
  { label: 'HTML', value: 'HTML' },
  { label: 'Markdown', value: 'Markdown' }
]

// 表单验证规则
const rules = {
  name: {
    required: true,
    message: '请输入通道名称',
    trigger: 'blur'
  },
  type: {
    required: true,
    message: '请选择通道类型',
    trigger: 'change'
  }
}

// 表格列配置
const columns = [
  {
    title: 'ID',
    key: 'id',
    width: 80
  },
  {
    title: '名称',
    key: 'name',
    width: 200
  },
  {
    title: '类型',
    key: 'type',
    width: 120,
    render: (row: Channel) => {
      const typeMap: Record<string, string> = {
        telegram: 'Telegram',
        email: 'Email',
        slack: 'Slack',
        webhook: 'Webhook'
      }
      return typeMap[row.type] || row.type
    }
  },
  {
    title: '创建时间',
    key: 'createdAt',
    width: 180,
    render: (row: Channel) => new Date(row.createdAt).toLocaleString()
  },
  {
    title: '操作',
    key: 'actions',
    width: 200,
    render: (row: Channel) => {
      return h('div', { class: 'action-buttons' }, [
        h(
          NButton,
          {
            size: 'small',
            onClick: () => handleView(row)
          },
          { default: () => '查看' }
        ),
        h(
          NButton,
          {
            size: 'small',
            type: 'primary',
            onClick: () => handleEdit(row)
          },
          { default: () => '编辑' }
        ),
        h(
          NButton,
          {
            size: 'small',
            type: 'error',
            onClick: () => handleDeleteClick(row)
          },
          { default: () => '删除' }
        )
      ])
    }
  }
]

// 分页配置
const pagination = {
  pageSize: 10
}

// 方法
const loadChannels = async () => {
  loading.value = true
  try {
    const response = await channelApi.getChannels()
    channels.value = response.data || response
  } catch (error: any) {
    message.error('加载通道列表失败: ' + (error.message || '未知错误'))
  } finally {
    loading.value = false
  }
}

const handleTypeChange = (type: string) => {
  // 重置凭证
  formData.credentials = {}
  
  // 根据类型设置默认凭证结构
  if (type === 'telegram') {
    formData.credentials = {
      bot_token: '',
      chat_id: '',
      parse_mode: 'HTML'
    }
  } else if (type === 'email') {
    formData.credentials = {
      smtp_host: '',
      smtp_port: 587,
      smtp_username: '',
      smtp_password: '',
      sender: ''
    }
  }
}

const handleSubmit = async () => {
  try {
    await formRef.value?.validate()
    submitting.value = true
    
    if (editingChannel.value) {
      await channelApi.updateChannel(editingChannel.value.id, formData)
      message.success('通道更新成功')
    } else {
      await channelApi.createChannel(formData)
      message.success('通道创建成功')
    }
    
    showCreateModal.value = false
    loadChannels()
  } catch (error: any) {
    message.error('操作失败: ' + (error.message || '未知错误'))
  } finally {
    submitting.value = false
  }
}

const handleEdit = (channel: Channel) => {
  editingChannel.value = channel
  formData.name = channel.name
  formData.type = channel.type
  formData.credentials = { ...channel.credentials }
  showCreateModal.value = true
}

const handleView = (channel: Channel) => {
  // TODO: 实现查看详情功能
  message.info('查看功能开发中')
}

const handleDeleteClick = (channel: Channel) => {
  deletingChannel.value = channel
  showDeleteModal.value = true
}

const handleDelete = async () => {
  if (!deletingChannel.value) return
  
  try {
    deleting.value = true
    await channelApi.deleteChannel(deletingChannel.value.id)
    message.success('通道删除成功')
    showDeleteModal.value = false
    loadChannels()
  } catch (error: any) {
    message.error('删除失败: ' + (error.message || '未知错误'))
  } finally {
    deleting.value = false
  }
}

const resetForm = () => {
  editingChannel.value = null
  formData.name = ''
  formData.type = ''
  formData.credentials = {}
  formRef.value?.restoreValidation()
}

// 监听模态框关闭
const handleModalClose = () => {
  resetForm()
}

// 生命周期
onMounted(() => {
  loadChannels()
})
</script>

<style scoped>
.channels-page {
  padding: 24px;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.page-header h1 {
  margin: 0;
  font-size: 24px;
  font-weight: 600;
}

.modal-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}

:deep(.action-buttons) {
  display: flex;
  gap: 8px;
}
</style> 