<template>
  <div class="dashboard-container">
    <div class="page-header">
      <h1>系统仪表板</h1>
      <n-button @click="refreshData" :loading="loading">
        <template #icon>
          <n-icon>↻</n-icon>
        </template>
        刷新数据
      </n-button>
    </div>

    <!-- 统计卡片 -->
    <div class="stats-grid">
      <n-card class="stat-card">
        <div class="stat-content">
          <div class="stat-icon success">
            <n-icon size="32">
              <svg viewBox="0 0 24 24">
                <path fill="currentColor" d="M20 4H4c-1.1 0-1.99.9-1.99 2L2 18c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V6c0-1.1-.9-2-2-2zm-5 14H4v-4h11v4zm0-5H4V9h11v4zm5 5h-4V9h4v9z"/>
              </svg>
            </n-icon>
          </div>
          <div class="stat-info">
            <div class="stat-number">{{ stats.channels }}</div>
            <div class="stat-label">通知通道</div>
            <div class="stat-trend">
              <n-icon size="16" color="#18a058">↑</n-icon>
              <span>活跃</span>
            </div>
          </div>
        </div>
      </n-card>

      <n-card class="stat-card">
        <div class="stat-content">
          <div class="stat-icon primary">
            <n-icon size="32">
              <svg viewBox="0 0 24 24">
                <path fill="currentColor" d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z"/>
              </svg>
            </n-icon>
          </div>
          <div class="stat-info">
            <div class="stat-number">{{ stats.topics }}</div>
            <div class="stat-label">消息主题</div>
            <div class="stat-trend">
              <n-icon size="16" color="#2080f0">→</n-icon>
              <span>稳定</span>
            </div>
          </div>
        </div>
      </n-card>

      <n-card class="stat-card">
        <div class="stat-content">
          <div class="stat-icon warning">
            <n-icon size="32">
              <svg viewBox="0 0 24 24">
                <path fill="currentColor" d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/>
              </svg>
            </n-icon>
          </div>
          <div class="stat-info">
            <div class="stat-number">{{ stats.routings }}</div>
            <div class="stat-label">路由规则</div>
            <div class="stat-trend">
              <n-icon size="16" color="#f0a020">↑</n-icon>
              <span>增长</span>
            </div>
          </div>
        </div>
      </n-card>

      <n-card class="stat-card">
        <div class="stat-content">
          <div class="stat-icon error">
            <n-icon size="32">
              <svg viewBox="0 0 24 24">
                <path fill="currentColor" d="M20 2H4c-1.1 0-1.99.9-1.99 2L2 22l4-4h14c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2zm-2 12H6v-2h12v2zm0-3H6V9h12v2zm0-3H6V6h12v2z"/>
              </svg>
            </n-icon>
          </div>
          <div class="stat-info">
            <div class="stat-number">{{ stats.messages }}</div>
            <div class="stat-label">消息总数</div>
            <div class="stat-trend">
              <n-icon size="16" color="#d03050">↑</n-icon>
              <span>活跃</span>
            </div>
          </div>
        </div>
      </n-card>
    </div>

    <!-- 图表区域 -->
    <div class="charts-grid">
      <!-- 通道类型分布 -->
      <n-card title="通道类型分布" class="chart-card">
        <div class="chart-placeholder">
          <n-icon size="48" color="#ccc">
            <svg viewBox="0 0 24 24">
              <path fill="currentColor" d="M3 13h8V3H3v10zm0 8h8v-6H3v6zm10 0h8V11h-8v10zm0-18v6h8V3h-8z"/>
            </svg>
          </n-icon>
          <p>图表功能开发中</p>
        </div>
      </n-card>

      <!-- 消息发送趋势 -->
      <n-card title="消息发送趋势" class="chart-card">
        <div class="chart-placeholder">
          <n-icon size="48" color="#ccc">
            <svg viewBox="0 0 24 24">
              <path fill="currentColor" d="M3 13h8V3H3v10zm0 8h8v-6H3v6zm10 0h8V11h-8v10zm0-18v6h8V3h-8z"/>
            </svg>
          </n-icon>
          <p>图表功能开发中</p>
        </div>
      </n-card>
    </div>

    <!-- 最近活动 -->
    <n-card title="最近活动" class="activity-card">
      <n-list>
        <n-list-item v-for="activity in recentActivities" :key="activity.id">
          <template #prefix>
            <n-avatar :color="activity.color" size="small">
              {{ activity.icon }}
            </n-avatar>
          </template>
          <n-thing :title="activity.title" :description="activity.description">
            <template #description>
              <div class="activity-meta">
                <span>{{ activity.description }}</span>
                <span class="activity-time">{{ activity.time }}</span>
              </div>
            </template>
          </n-thing>
        </n-list-item>
      </n-list>
    </n-card>

    <!-- 系统状态 -->
    <n-card title="系统状态" class="status-card">
      <n-space vertical>
        <div class="status-item">
          <span class="status-label">API 服务</span>
          <n-tag type="success" size="small">正常</n-tag>
        </div>
        <div class="status-item">
          <span class="status-label">数据库</span>
          <n-tag type="success" size="small">正常</n-tag>
        </div>
        <div class="status-item">
          <span class="status-label">消息队列</span>
          <n-tag type="warning" size="small">待实现</n-tag>
        </div>
        <div class="status-item">
          <span class="status-label">监控服务</span>
          <n-tag type="info" size="small">待实现</n-tag>
        </div>
      </n-space>
    </n-card>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useMessage } from 'naive-ui'
import { channelApi, topicApi } from '../../api'

const message = useMessage()

// 响应式数据
const loading = ref(false)
const stats = ref({
  channels: 0,
  topics: 0,
  routings: 0,
  messages: 0
})

// 最近活动
const recentActivities = ref([
  {
    id: 1,
    icon: '📡',
    title: '新通道创建',
    description: '创建了 Telegram 通知通道',
    time: '2分钟前',
    color: '#18a058'
  },
  {
    id: 2,
    icon: '🎯',
    title: '主题配置',
    description: '配置了 GitHub 消息主题',
    time: '5分钟前',
    color: '#2080f0'
  },
  {
    id: 3,
    icon: '🔄',
    title: '路由更新',
    description: '更新了消息路由规则',
    time: '10分钟前',
    color: '#f0a020'
  },
  {
    id: 4,
    icon: '📨',
    title: '消息发送',
    description: '成功发送 15 条消息',
    time: '15分钟前',
    color: '#d03050'
  }
])

// 加载统计数据
const loadStats = async () => {
  loading.value = true
  try {
    const [channelsRes, topicsRes] = await Promise.all([
      channelApi.getChannels(),
      topicApi.getTopics()
    ])
    
    const channels = channelsRes.data || channelsRes
    const topics = topicsRes.data || topicsRes
    
    // 计算路由总数
    let totalRoutings = 0
    for (const topic of topics) {
      try {
        const routingsRes = await topicApi.getTopicRoutings(topic.id)
        const routings = routingsRes.data || routingsRes
        totalRoutings += routings.length
      } catch (error) {
        // 忽略单个主题路由加载失败
      }
    }
    
    stats.value = {
      channels: channels.length,
      topics: topics.length,
      routings: totalRoutings,
      messages: Math.floor(Math.random() * 1000) // 模拟数据
    }
  } catch (error: any) {
    message.error('加载统计数据失败: ' + (error.message || '未知错误'))
  } finally {
    loading.value = false
  }
}

// 刷新数据
const refreshData = () => {
  loadStats()
  message.success('数据已刷新')
}

onMounted(() => {
  loadStats()
})
</script>

<style scoped>
.dashboard-container {
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

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 16px;
  margin-bottom: 24px;
}

.stat-card {
  transition: transform 0.2s;
}

.stat-card:hover {
  transform: translateY(-2px);
}

.stat-content {
  display: flex;
  align-items: center;
  gap: 16px;
}

.stat-icon {
  flex-shrink: 0;
}

.stat-icon.success {
  color: #18a058;
}

.stat-icon.primary {
  color: #2080f0;
}

.stat-icon.warning {
  color: #f0a020;
}

.stat-icon.error {
  color: #d03050;
}

.stat-info {
  flex: 1;
}

.stat-number {
  font-size: 32px;
  font-weight: bold;
  color: #333;
  line-height: 1;
}

.stat-label {
  margin-top: 4px;
  color: #666;
  font-size: 14px;
}

.stat-trend {
  display: flex;
  align-items: center;
  gap: 4px;
  margin-top: 4px;
  font-size: 12px;
  color: #999;
}

.charts-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(400px, 1fr));
  gap: 16px;
  margin-bottom: 24px;
}

.chart-card {
  min-height: 300px;
}

.chart-placeholder {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 200px;
  color: #ccc;
}

.chart-placeholder p {
  margin-top: 16px;
  font-size: 14px;
}

.activity-card {
  margin-bottom: 24px;
}

.activity-meta {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
}

.activity-time {
  color: #999;
  font-size: 12px;
}

.status-card {
  margin-bottom: 24px;
}

.status-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8px 0;
}

.status-label {
  font-size: 14px;
  color: #333;
}

@media (max-width: 768px) {
  .stats-grid {
    grid-template-columns: repeat(2, 1fr);
  }
  
  .charts-grid {
    grid-template-columns: 1fr;
  }
  
  .page-header {
    flex-direction: column;
    gap: 16px;
    align-items: flex-start;
  }
}
</style> 