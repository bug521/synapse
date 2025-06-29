<template>
  <div class="topic-detail-page">
    <div class="page-header">
      <n-button @click="router.back()">
        <template #icon>
          <n-icon>←</n-icon>
        </template>
        返回
      </n-button>
      <h1>{{ topic?.name }} - 主题详情</h1>
    </div>

    <!-- 主题信息 -->
    <n-card title="主题信息" class="info-card">
      <n-descriptions :column="2" bordered>
        <n-descriptions-item label="主题ID">{{ topic?.id }}</n-descriptions-item>
        <n-descriptions-item label="主题名称">{{ topic?.name }}</n-descriptions-item>
        <n-descriptions-item label="Webhook Key">
          <div class="webhook-key-display">
            <span class="key-text">{{ topic?.webhook_key }}</span>
            <n-button size="tiny" type="primary" ghost @click="copyWebhookKey">
              复制
            </n-button>
          </div>
        </n-descriptions-item>
        <n-descriptions-item label="发送策略">
          {{ getSendingStrategyText(topic?.sending_strategy) }}
        </n-descriptions-item>
        <n-descriptions-item label="执行模式">
          {{ getExecutionModeText(topic?.execution_mode) }}
        </n-descriptions-item>
        <n-descriptions-item label="创建时间">
          {{ topic?.created_at ? new Date(topic.created_at).toLocaleString() : '' }}
        </n-descriptions-item>
        <n-descriptions-item label="描述" :span="2">
          {{ topic?.description || '暂无描述' }}
        </n-descriptions-item>
      </n-descriptions>
    </n-card>

    <!-- 路由配置 -->
    <n-card title="路由配置" class="routing-card">
      <template #header-extra>
        <n-button type="primary" @click="showCreateRoutingModal = true">
          <template #icon>
            <n-icon>+</n-icon>
          </template>
          添加路由
        </n-button>
      </template>

      <n-data-table
        :columns="routingColumns"
        :data="routings"
        :loading="routingLoading"
        :bordered="false"
      />
    </n-card>

    <!-- 创建路由模态框 -->
    <n-modal v-model:show="showCreateRoutingModal" preset="card" style="width: 700px">
      <template #header>
        <h3>添加路由规则</h3>
      </template>
      
      <n-form
        ref="routingFormRef"
        :model="routingFormData"
        :rules="routingRules"
        label-placement="left"
        label-width="auto"
        require-mark-placement="right-hanging"
      >
        <n-form-item label="选择通道" path="channel_id">
          <n-select
            v-model:value="routingFormData.channel_id"
            :options="channelOptions"
            placeholder="请选择通知通道"
            filterable
          />
        </n-form-item>
        
        <n-form-item label="优先级" path="priority">
          <n-input-number
            v-model:value="routingFormData.priority"
            :min="0"
            :max="100"
            placeholder="数字越小优先级越高"
          />
        </n-form-item>
        
        <n-form-item label="变量映射" path="variable_mappings">
          <n-input
            v-model:value="variableMappingsText"
            type="textarea"
            placeholder="请输入JSON格式的变量映射，例如：&#10;{&#10;  &quot;title&quot;: &quot;repository.name&quot;,&#10;  &quot;action&quot;: &quot;action&quot;&#10;}"
            :rows="6"
          />
        </n-form-item>
        
        <n-form-item label="消息模板" path="message_template">
          <n-input
            v-model:value="routingFormData.message_template"
            type="textarea"
            placeholder="请输入消息模板，支持变量替换，例如：&#10;仓库 {{.title}} 有新活动: {{.action}}"
            :rows="4"
          />
        </n-form-item>
      </n-form>
      
      <template #footer>
        <div class="modal-footer">
          <n-button @click="showCreateRoutingModal = false">取消</n-button>
          <n-button type="primary" :loading="routingSubmitting" @click="handleCreateRouting">
            创建
          </n-button>
        </div>
      </template>
    </n-modal>

    <!-- 编辑路由模态框 -->
    <n-modal v-model:show="showEditRoutingModal" preset="card" style="width: 700px">
      <template #header>
        <h3>编辑路由规则</h3>
      </template>
      
      <n-form
        ref="editRoutingFormRef"
        :model="editRoutingFormData"
        :rules="routingRules"
        label-placement="left"
        label-width="auto"
        require-mark-placement="right-hanging"
      >
        <n-form-item label="选择通道" path="channel_id">
          <n-select
            v-model:value="editRoutingFormData.channel_id"
            :options="channelOptions"
            placeholder="请选择通知通道"
            filterable
            disabled
          />
        </n-form-item>
        
        <n-form-item label="优先级" path="priority">
          <n-input-number
            v-model:value="editRoutingFormData.priority"
            :min="0"
            :max="100"
            placeholder="数字越小优先级越高"
          />
        </n-form-item>
        
        <n-form-item label="变量映射" path="variable_mappings">
          <n-input
            v-model:value="editVariableMappingsText"
            type="textarea"
            placeholder="请输入JSON格式的变量映射"
            :rows="6"
          />
        </n-form-item>
        
        <n-form-item label="消息模板" path="message_template">
          <n-input
            v-model:value="editRoutingFormData.message_template"
            type="textarea"
            placeholder="请输入消息模板，支持变量替换"
            :rows="4"
          />
        </n-form-item>
      </n-form>
      
      <template #footer>
        <div class="modal-footer">
          <n-button @click="showEditRoutingModal = false">取消</n-button>
          <n-button type="primary" :loading="routingSubmitting" @click="handleUpdateRouting">
            更新
          </n-button>
        </div>
      </template>
    </n-modal>

    <!-- 删除确认模态框 -->
    <n-modal v-model:show="showDeleteRoutingModal" preset="card" style="width: 400px">
      <template #header>
        <h3>确认删除</h3>
      </template>
      
      <p>确定要删除这个路由规则吗？此操作不可恢复。</p>
      
      <template #footer>
        <div class="modal-footer">
          <n-button @click="showDeleteRoutingModal = false">取消</n-button>
          <n-button type="error" :loading="routingDeleting" @click="handleDeleteRouting">
            删除
          </n-button>
        </div>
      </template>
    </n-modal>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted, h } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useMessage } from 'naive-ui'
import { topicApi, channelApi, routingApi, type Topic, type Channel, type Routing, type CreateRoutingRequest } from '../../api'

const route = useRoute()
const router = useRouter()
const message = useMessage()

// 响应式数据
const topic = ref<Topic | null>(null)
const channels = ref<Channel[]>([])
const routings = ref<Routing[]>([])
const routingLoading = ref(false)
const routingSubmitting = ref(false)
const routingDeleting = ref(false)

// 模态框状态
const showCreateRoutingModal = ref(false)
const showEditRoutingModal = ref(false)
const showDeleteRoutingModal = ref(false)

// 表单引用
const routingFormRef = ref()
const editRoutingFormRef = ref()

// 编辑状态
const editingRouting = ref<Routing | null>(null)
const deletingRouting = ref<Routing | null>(null)

// 表单数据
const routingFormData = reactive<CreateRoutingRequest>({
  topic_id: 0,
  channel_id: 0,
  priority: 0,
  variable_mappings: {},
  message_template: ''
})

const editRoutingFormData = reactive<CreateRoutingRequest>({
  topic_id: 0,
  channel_id: 0,
  priority: 0,
  variable_mappings: {},
  message_template: ''
})

// 计算属性
const channelOptions = computed(() => {
  return channels.value.map(channel => ({
    label: `${channel.name} (${channel.type})`,
    value: channel.id
  }))
})

const variableMappingsText = computed({
  get: () => JSON.stringify(routingFormData.variable_mappings, null, 2),
  set: (value: string) => {
    try {
      routingFormData.variable_mappings = JSON.parse(value)
    } catch (e) {
      // 忽略解析错误
    }
  }
})

const editVariableMappingsText = computed({
  get: () => JSON.stringify(editRoutingFormData.variable_mappings, null, 2),
  set: (value: string) => {
    try {
      editRoutingFormData.variable_mappings = JSON.parse(value)
    } catch (e) {
      // 忽略解析错误
    }
  }
})

// 表单验证规则
const routingRules = {
  channel_id: {
    required: true,
    message: '请选择通知通道',
    trigger: 'change'
  },
  priority: {
    required: true,
    message: '请输入优先级',
    trigger: 'blur'
  }
}

// 表格列配置
const routingColumns = [
  {
    title: '通道',
    key: 'channel_id',
    width: 200,
    render: (row: Routing) => {
      const channel = channels.value.find(c => c.id === row.channel_id)
      return channel ? `${channel.name} (${channel.type})` : '未知通道'
    }
  },
  {
    title: '优先级',
    key: 'priority',
    width: 100
  },
  {
    title: '变量映射',
    key: 'variable_mappings',
    width: 200,
    render: (row: Routing) => {
      const mappings = Object.keys(row.variable_mappings || {})
      return mappings.length > 0 ? mappings.join(', ') : '无'
    }
  },
  {
    title: '消息模板',
    key: 'message_template',
    width: 300,
    render: (row: Routing) => {
      return row.message_template || '使用默认模板'
    }
  },
  {
    title: '操作',
    key: 'actions',
    width: 150,
    render: (row: Routing) => {
      return h('div', { class: 'action-buttons' }, [
        h(
          'n-button',
          {
            size: 'small',
            type: 'info',
            onClick: () => handleEditRouting(row)
          },
          { default: () => '编辑' }
        ),
        h(
          'n-button',
          {
            size: 'small',
            type: 'error',
            onClick: () => handleDeleteRoutingClick(row)
          },
          { default: () => '删除' }
        )
      ])
    }
  }
]

// 方法
const loadTopic = async () => {
  const topicId = parseInt(route.params.id as string)
  try {
    const response = await topicApi.getTopic(topicId)
    topic.value = response.data || response
  } catch (error: any) {
    message.error('加载主题信息失败: ' + (error.message || '未知错误'))
  }
}

const loadChannels = async () => {
  try {
    const response = await channelApi.getChannels()
    channels.value = response.data || response
  } catch (error: any) {
    message.error('加载通道列表失败: ' + (error.message || '未知错误'))
  }
}

const loadRoutings = async () => {
  const topicId = parseInt(route.params.id as string)
  routingLoading.value = true
  try {
    const response = await topicApi.getTopicRoutings(topicId)
    routings.value = response.data || response
  } catch (error: any) {
    message.error('加载路由列表失败: ' + (error.message || '未知错误'))
  } finally {
    routingLoading.value = false
  }
}

const handleCreateRouting = async () => {
  try {
    await routingFormRef.value?.validate()
    routingSubmitting.value = true
    
    const topicId = parseInt(route.params.id as string)
    routingFormData.topic_id = topicId
    
    await routingApi.createRouting(routingFormData)
    message.success('路由创建成功')
    showCreateRoutingModal.value = false
    loadRoutings()
  } catch (error: any) {
    message.error('创建失败: ' + (error.message || '未知错误'))
  } finally {
    routingSubmitting.value = false
  }
}

const handleEditRouting = (routing: Routing) => {
  editingRouting.value = routing
  editRoutingFormData.topic_id = routing.topic_id
  editRoutingFormData.channel_id = routing.channel_id
  editRoutingFormData.priority = routing.priority
  editRoutingFormData.variable_mappings = { ...routing.variable_mappings }
  editRoutingFormData.message_template = routing.message_template
  showEditRoutingModal.value = true
}

const handleUpdateRouting = async () => {
  if (!editingRouting.value) return
  
  try {
    await editRoutingFormRef.value?.validate()
    routingSubmitting.value = true
    
    await routingApi.updateRouting(
      editingRouting.value.topic_id,
      editingRouting.value.channel_id,
      {
        priority: editRoutingFormData.priority,
        variable_mappings: editRoutingFormData.variable_mappings,
        message_template: editRoutingFormData.message_template
      }
    )
    message.success('路由更新成功')
    showEditRoutingModal.value = false
    loadRoutings()
  } catch (error: any) {
    message.error('更新失败: ' + (error.message || '未知错误'))
  } finally {
    routingSubmitting.value = false
  }
}

const handleDeleteRoutingClick = (routing: Routing) => {
  deletingRouting.value = routing
  showDeleteRoutingModal.value = true
}

const handleDeleteRouting = async () => {
  if (!deletingRouting.value) return
  
  try {
    routingDeleting.value = true
    await routingApi.deleteRouting(
      deletingRouting.value.topic_id,
      deletingRouting.value.channel_id
    )
    message.success('路由删除成功')
    showDeleteRoutingModal.value = false
    loadRoutings()
  } catch (error: any) {
    message.error('删除失败: ' + (error.message || '未知错误'))
  } finally {
    routingDeleting.value = false
  }
}

const copyWebhookKey = () => {
  if (!topic.value?.webhook_key) return
  
  navigator.clipboard.writeText(topic.value.webhook_key).then(() => {
    message.success('Webhook Key已复制到剪贴板')
  }).catch(() => {
    message.error('复制失败')
  })
}

const getSendingStrategyText = (strategy?: string) => {
  const strategyMap: Record<string, string> = {
    'all': '发送给所有',
    'failover': '故障转移'
  }
  return strategyMap[strategy || ''] || strategy || '未知'
}

const getExecutionModeText = (mode?: string) => {
  const modeMap: Record<string, string> = {
    'async': '异步',
    'sync': '同步'
  }
  return modeMap[mode || ''] || mode || '未知'
}

// 生命周期
onMounted(() => {
  loadTopic()
  loadChannels()
  loadRoutings()
})
</script>

<style scoped>
.topic-detail-page {
  padding: 24px;
}

.page-header {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-bottom: 24px;
}

.page-header h1 {
  margin: 0;
  font-size: 24px;
  font-weight: 600;
}

.info-card {
  margin-bottom: 24px;
}

.routing-card {
  margin-bottom: 24px;
}

.webhook-key-display {
  display: flex;
  align-items: center;
  gap: 8px;
}

.key-text {
  font-family: monospace;
  font-size: 12px;
  color: #666;
  word-break: break-all;
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
</style> 