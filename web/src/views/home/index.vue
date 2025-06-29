<template>
  <div class="home-container">
    <!-- 欢迎卡片 -->
    <n-card title="欢迎来到 Synapse" class="welcome-card">
      <div class="welcome-content">
        <n-icon size="64" color="#18a058">
          <svg viewBox="0 0 24 24">
            <path fill="currentColor" d="M12 2L2 7l10 5 10-5-10-5zM2 17l10 5 10-5M2 12l10 5 10-5"/>
          </svg>
        </n-icon>
        <h2>消息转发服务管理平台</h2>
        <p>灵活、高效的 Webhook 转通知中间件，支持多通道、多策略的消息分发</p>
      </div>
    </n-card>

    <!-- 统计卡片 -->
    <div class="stats-grid">
      <n-card class="stat-card">
        <div class="stat-content">
          <div class="stat-icon">
            <n-icon size="32" color="#18a058">
              <svg viewBox="0 0 24 24">
                <path fill="currentColor" d="M20 4H4c-1.1 0-1.99.9-1.99 2L2 18c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V6c0-1.1-.9-2-2-2zm-5 14H4v-4h11v4zm0-5H4V9h11v4zm5 5h-4V9h4v9z"/>
              </svg>
            </n-icon>
          </div>
          <div class="stat-info">
            <div class="stat-number">{{ stats.channels }}</div>
            <div class="stat-label">通知通道</div>
          </div>
        </div>
      </n-card>

      <n-card class="stat-card">
        <div class="stat-content">
          <div class="stat-icon">
            <n-icon size="32" color="#2080f0">
              <svg viewBox="0 0 24 24">
                <path fill="currentColor" d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z"/>
              </svg>
            </n-icon>
          </div>
          <div class="stat-info">
            <div class="stat-number">{{ stats.topics }}</div>
            <div class="stat-label">消息主题</div>
          </div>
        </div>
      </n-card>

      <n-card class="stat-card">
        <div class="stat-content">
          <div class="stat-icon">
            <n-icon size="32" color="#f0a020">
              <svg viewBox="0 0 24 24">
                <path fill="currentColor" d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/>
              </svg>
            </n-icon>
          </div>
          <div class="stat-info">
            <div class="stat-number">{{ stats.routings }}</div>
            <div class="stat-label">路由规则</div>
          </div>
        </div>
      </n-card>

      <n-card class="stat-card">
        <div class="stat-content">
          <div class="stat-icon">
            <n-icon size="32" color="#d03050">
              <svg viewBox="0 0 24 24">
                <path fill="currentColor" d="M20 2H4c-1.1 0-1.99.9-1.99 2L2 22l4-4h14c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2zm-2 12H6v-2h12v2zm0-3H6V9h12v2zm0-3H6V6h12v2z"/>
              </svg>
            </n-icon>
          </div>
          <div class="stat-info">
            <div class="stat-number">{{ stats.messages }}</div>
            <div class="stat-label">消息总数</div>
          </div>
        </div>
      </n-card>
    </div>

    <!-- 快速操作 -->
    <n-card title="快速操作" class="actions-card">
      <div class="actions-grid">
        <n-button type="primary" size="large" @click="$router.push('/channels')">
          <template #icon>
            <n-icon>
              <svg viewBox="0 0 24 24">
                <path fill="currentColor" d="M20 4H4c-1.1 0-1.99.9-1.99 2L2 18c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V6c0-1.1-.9-2-2-2zm-5 14H4v-4h11v4zm0-5H4V9h11v4zm5 5h-4V9h4v9z"/>
              </svg>
            </n-icon>
          </template>
          管理通知通道
        </n-button>

        <n-button type="info" size="large" @click="$router.push('/topics')">
          <template #icon>
            <n-icon>
              <svg viewBox="0 0 24 24">
                <path fill="currentColor" d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z"/>
              </svg>
            </n-icon>
          </template>
          创建消息主题
        </n-button>

        <n-button type="success" size="large" @click="$router.push('/dashboard')">
          <template #icon>
            <n-icon>
              <svg viewBox="0 0 24 24">
                <path fill="currentColor" d="M3 13h8V3H3v10zm0 8h8v-6H3v6zm10 0h8V11h-8v10zm0-18v6h8V3h-8z"/>
              </svg>
            </n-icon>
          </template>
          查看仪表板
        </n-button>

        <n-button type="warning" size="large" @click="$router.push('/profile')">
          <template #icon>
            <n-icon>
              <svg viewBox="0 0 24 24">
                <path fill="currentColor" d="M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm0 2c-2.67 0-8 1.34-8 4v2h16v-2c0-2.66-5.33-4-8-4z"/>
              </svg>
            </n-icon>
          </template>
          个人设置
        </n-button>
      </div>
    </n-card>

    <!-- 系统特性 -->
    <n-card title="系统特性" class="features-card">
      <div class="features-grid">
        <div class="feature-item">
          <n-icon size="24" color="#18a058">
            <svg viewBox="0 0 24 24">
              <path fill="currentColor" d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z"/>
            </svg>
          </n-icon>
          <div class="feature-content">
            <h4>多通道支持</h4>
            <p>支持 Telegram、Email、Slack、Webhook 等多种通知通道</p>
          </div>
        </div>

        <div class="feature-item">
          <n-icon size="24" color="#2080f0">
            <svg viewBox="0 0 24 24">
              <path fill="currentColor" d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/>
            </svg>
          </n-icon>
          <div class="feature-content">
            <h4>灵活路由</h4>
            <p>支持多种发送策略和优先级配置，满足不同业务需求</p>
          </div>
        </div>

        <div class="feature-item">
          <n-icon size="24" color="#f0a020">
            <svg viewBox="0 0 24 24">
              <path fill="currentColor" d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/>
            </svg>
          </n-icon>
          <div class="feature-content">
            <h4>异步处理</h4>
            <p>支持异步和同步执行模式，提高系统性能和可靠性</p>
          </div>
        </div>

        <div class="feature-item">
          <n-icon size="24" color="#d03050">
            <svg viewBox="0 0 24 24">
              <path fill="currentColor" d="M20 2H4c-1.1 0-1.99.9-1.99 2L2 22l4-4h14c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2zm-2 12H6v-2h12v2zm0-3H6V9h12v2zm0-3H6V6h12v2z"/>
            </svg>
          </n-icon>
          <div class="feature-content">
            <h4>消息历史</h4>
            <p>完整的消息发送历史和投递状态跟踪</p>
          </div>
        </div>
      </div>
    </n-card>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { channelApi, topicApi, routingApi } from '../../api'

// 统计数据
const stats = ref({
  channels: 0,
  topics: 0,
  routings: 0,
  messages: 0
})

// 加载统计数据
const loadStats = async () => {
  try {
    // 并行加载各项统计数据
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
      messages: 0 // TODO: 实现消息统计
    }
  } catch (error) {
    console.error('加载统计数据失败:', error)
  }
}

onMounted(() => {
  loadStats()
})
</script>

<style scoped>
.home-container {
  padding: 24px;
}

.welcome-card {
  margin-bottom: 24px;
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

.stat-info {
  flex: 1;
}

.stat-number {
  font-size: 32px;
  font-weight: bold;
  color: #18a058;
  line-height: 1;
}

.stat-label {
  margin-top: 4px;
  color: #666;
  font-size: 14px;
}

.actions-card {
  margin-bottom: 24px;
}

.actions-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 16px;
}

.features-card {
  margin-bottom: 24px;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 24px;
}

.feature-item {
  display: flex;
  align-items: flex-start;
  gap: 16px;
  padding: 16px;
  border-radius: 8px;
  background: #f9f9f9;
  transition: background-color 0.2s;
}

.feature-item:hover {
  background: #f0f0f0;
}

.feature-content h4 {
  margin: 0 0 8px 0;
  color: #333;
  font-size: 16px;
}

.feature-content p {
  margin: 0;
  color: #666;
  font-size: 14px;
  line-height: 1.5;
}

@media (max-width: 768px) {
  .stats-grid {
    grid-template-columns: repeat(2, 1fr);
  }
  
  .actions-grid {
    grid-template-columns: 1fr;
  }
  
  .features-grid {
    grid-template-columns: 1fr;
  }
}
</style>
