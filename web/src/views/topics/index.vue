<template>
  <div class="topics-page">
    <div class="page-header">
      <h1>消息主题管理</h1>
      <n-button type="primary" @click="showCreateModal = true">
        <template #icon>
          <n-icon>+</n-icon>
        </template>
        创建主题
      </n-button>
    </div>

    <!-- 主题列表 -->
    <n-card>
      <n-data-table
        :columns="columns"
        :data="topics"
        :loading="loading"
        :pagination="pagination"
        :bordered="false"
      />
    </n-card>

    <!-- 创建/编辑主题模态框 -->
    <n-modal v-model:show="showCreateModal" :title="editingTopic ? '编辑主题' : '创建主题'" preset="card" style="width: 600px">
      
      <n-form
        ref="formRef"
        :model="formData"
        :rules="rules"
        label-placement="left"
        label-width="auto"
        require-mark-placement="right-hanging"
      >
        <n-form-item label="主题名称" path="name">
          <n-input v-model:value="formData.name" placeholder="请输入主题名称" />
        </n-form-item>
        
        <n-form-item label="发送策略" path="sendingStrategy">
          <n-select
            v-model:value="formData.sendingStrategy"
            :options="sendingStrategyOptions"
            placeholder="请选择发送策略"
          />
        </n-form-item>
        
        <n-form-item label="执行模式" path="executionMode">
          <n-select
            v-model:value="formData.executionMode"
            :options="executionModeOptions"
            placeholder="请选择执行模式"
          />
        </n-form-item>
        
        <n-form-item label="描述" path="description">
          <n-input
            v-model:value="formData.description"
            type="textarea"
            placeholder="请输入主题描述"
            :rows="3"
          />
        </n-form-item>
      </n-form>
      
      <template #footer>
        <div class="modal-footer">
          <n-button @click="showCreateModal = false">取消</n-button>
          <n-button type="primary" :loading="submitting" @click="handleSubmit">
            {{ editingTopic ? '更新' : '创建' }}
          </n-button>
        </div>
      </template>
    </n-modal>

    <!-- 删除确认模态框 -->
    <n-modal v-model:show="showDeleteModal" preset="card" style="width: 400px">
      <template #header>
        <h3>确认删除</h3>
      </template>
      
      <p>确定要删除主题 "{{ deletingTopic?.name }}" 吗？此操作不可恢复。</p>
      
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
import { useRouter } from 'vue-router'
import {NButton, useMessage} from 'naive-ui'
import { topicApi, type Topic, type CreateTopicRequest } from '../../api'

const router = useRouter()
const message = useMessage()

// 响应式数据
const loading = ref(false)
const submitting = ref(false)
const deleting = ref(false)
const topics = ref<Topic[]>([])
const showCreateModal = ref(false)
const showDeleteModal = ref(false)
const editingTopic = ref<Topic | null>(null)
const deletingTopic = ref<Topic | null>(null)
const formRef = ref()

// 表单数据
const formData = reactive<CreateTopicRequest>({
  name: '',
  sendingStrategy: 'all',
  executionMode: 'async',
  description: ''
})

// 选项配置
const sendingStrategyOptions = [
  { label: '发送给所有', value: 'all' },
  { label: '故障转移', value: 'failover' }
]

const executionModeOptions = [
  { label: '异步', value: 'async' },
  { label: '同步', value: 'sync' }
]

// 表单验证规则
const rules = {
  name: {
    required: true,
    message: '请输入主题名称',
    trigger: 'blur'
  },
  sending_strategy: {
    required: true,
    message: '请选择发送策略',
    trigger: 'change'
  },
  execution_mode: {
    required: true,
    message: '请选择执行模式',
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
    title: 'Webhook Key',
    key: 'webhookKey',
    width: 200,
    render: (row: Topic) => {
      return h('div', { class: 'webhook-key' }, [
        h('span', { class: 'key-text' }, row.webhookKey.substring(0, 8) + '...'),
        h('n-button', {
          size: 'tiny',
          type: 'primary',
          ghost: true,
          onClick: () => copyWebhookKey(row.webhookKey)
        }, { default: () => '复制' })
      ])
    }
  },
  {
    title: '发送策略',
    key: 'sendingStrategy',
    width: 120,
    render: (row: Topic) => {
      const strategyMap: Record<string, string> = {
        'all': '发送给所有',
        'failover': '故障转移'
      }
      return strategyMap[row.sendingStrategy] || row.sendingStrategy
    }
  },
  {
    title: '执行模式',
    key: 'executionMode',
    width: 120,
    render: (row: Topic) => {
      const modeMap: Record<string, string> = {
        'async': '异步',
        'sync': '同步'
      }
      return modeMap[row.executionMode] || row.executionMode
    }
  },
  {
    title: '创建时间',
    key: 'createdAt',
    width: 180,
    render: (row: Topic) => new Date(row.createdAt).toLocaleString()
  },
  {
    title: '操作',
    key: 'actions',
    width: 250,
    render: (row: Topic) => {
      return h('div', { class: 'action-buttons' }, [
        h(
          NButton,
          {
            size: 'small',
            type: 'primary',
            onClick: () => handleViewDetail(row)
          },
          { default: () => '详情' }
        ),
        h(
            NButton,
          {
            size: 'small',
            type: 'info',
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
const loadTopics = async () => {
  loading.value = true
  try {
    const response = await topicApi.getTopics()
    topics.value = response.data || response
  } catch (error: any) {
    message.error('加载主题列表失败: ' + (error.message || '未知错误'))
  } finally {
    loading.value = false
  }
}

const handleSubmit = async () => {
  try {
    await formRef.value?.validate()
    submitting.value = true
    
    if (editingTopic.value) {
      await topicApi.updateTopic(editingTopic.value.id, formData)
      message.success('主题更新成功')
    } else {
      await topicApi.createTopic(formData)
      message.success('主题创建成功')
    }
    
    showCreateModal.value = false
    loadTopics()
  } catch (error: any) {
    message.error('操作失败: ' + (error.message || '未知错误'))
  } finally {
    submitting.value = false
  }
}

const handleEdit = (topic: Topic) => {
  editingTopic.value = topic
  formData.name = topic.name
  formData.sendingStrategy = topic.sendingStrategy
  formData.executionMode = topic.executionMode
  formData.description = topic.description
  showCreateModal.value = true
}

const handleViewDetail = (topic: Topic) => {
  router.push(`/topics/${topic.id}`)
}

const handleRegenerateKey = async (topic: Topic) => {
  try {
    await topicApi.regenerateWebhookKey(topic.id)
    message.success('Webhook Key重新生成成功')
    loadTopics()
  } catch (error: any) {
    message.error('重新生成失败: ' + (error.message || '未知错误'))
  }
}

const copyWebhookKey = (key: string) => {
  navigator.clipboard.writeText(key).then(() => {
    message.success('Webhook Key已复制到剪贴板')
  }).catch(() => {
    message.error('复制失败')
  })
}

const handleDeleteClick = (topic: Topic) => {
  deletingTopic.value = topic
  showDeleteModal.value = true
}

const handleDelete = async () => {
  if (!deletingTopic.value) return
  
  try {
    deleting.value = true
    await topicApi.deleteTopic(deletingTopic.value.id)
    message.success('主题删除成功')
    showDeleteModal.value = false
    loadTopics()
  } catch (error: any) {
    message.error('删除失败: ' + (error.message || '未知错误'))
  } finally {
    deleting.value = false
  }
}

const resetForm = () => {
  editingTopic.value = null
  formData.name = ''
  formData.sendingStrategy = 'all'
  formData.executionMode = 'async'
  formData.description = ''
  formRef.value?.restoreValidation()
}

// 监听模态框关闭
const handleModalClose = () => {
  resetForm()
}

// 生命周期
onMounted(() => {
  loadTopics()
})
</script>

<style scoped>
.topics-page {
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

.action-buttons {
  display: flex;
  gap: 8px;
}

.webhook-key {
  display: flex;
  align-items: center;
  gap: 8px;
}

.key-text {
  font-family: monospace;
  font-size: 12px;
  color: #666;
}

:deep(.action-buttons) {
  display: flex;
  gap: 8px;
}
</style> 