<template>
  <div class="dashboard-container">
    <n-grid :cols="24" :x-gap="16" :y-gap="16">
      <!-- 统计卡片 -->
      <n-grid-item :span="6">
        <n-card>
          <div class="stat-card">
            <n-icon size="48" color="#18a058">
              <svg viewBox="0 0 24 24">
                <path fill="currentColor" d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/>
              </svg>
            </n-icon>
            <div class="stat-info">
              <div class="stat-number">1,234</div>
              <div class="stat-label">总用户数</div>
            </div>
          </div>
        </n-card>
      </n-grid-item>
      
      <n-grid-item :span="6">
        <n-card>
          <div class="stat-card">
            <n-icon size="48" color="#2080f0">
              <svg viewBox="0 0 24 24">
                <path fill="currentColor" d="M13 2.05v3.03c3.39.49 6 3.39 6 6.92 0 .9-.18 1.75-.5 2.54l2.6 1.53c.56-1.24.9-2.62.9-4.07 0-5.18-3.95-9.45-9-9.95zM12 19c-3.87 0-7-3.13-7-7 0-3.53 2.61-6.43 6-6.92V2.05c-5.05.5-9 4.76-9 9.95 0 5.52 4.47 10 9.99 10 3.31 0 6.24-1.61 8.06-4.09l-2.6-1.53C16.17 17.98 14.21 19 12 19z"/>
              </svg>
            </n-icon>
            <div class="stat-info">
              <div class="stat-number">567</div>
              <div class="stat-label">活跃用户</div>
            </div>
          </div>
        </n-card>
      </n-grid-item>
      
      <n-grid-item :span="6">
        <n-card>
          <div class="stat-card">
            <n-icon size="48" color="#f0a020">
              <svg viewBox="0 0 24 24">
                <path fill="currentColor" d="M19 3H5c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2zm-5 14H7v-2h7v2zm3-4H7v-2h10v2zm0-4H7V7h10v2z"/>
              </svg>
            </n-icon>
            <div class="stat-info">
              <div class="stat-number">89</div>
              <div class="stat-label">今日订单</div>
            </div>
          </div>
        </n-card>
      </n-grid-item>
      
      <n-grid-item :span="6">
        <n-card>
          <div class="stat-card">
            <n-icon size="48" color="#d03050">
              <svg viewBox="0 0 24 24">
                <path fill="currentColor" d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm1 15h-2v-6h2v6zm0-8h-2V7h2v2z"/>
              </svg>
            </n-icon>
            <div class="stat-info">
              <div class="stat-number">¥12,345</div>
              <div class="stat-label">总收入</div>
            </div>
          </div>
        </n-card>
      </n-grid-item>
      
      <!-- 图表区域 -->
      <n-grid-item :span="16">
        <n-card title="数据趋势">
          <div class="chart-placeholder">
            <n-icon size="64" color="#d9d9d9">
              <svg viewBox="0 0 24 24">
                <path fill="currentColor" d="M3 13h2v-2H3v2zm0 4h2v-2H3v2zm0-8h2V7H3v2zm4 4h14v-2H7v2zm0 4h14v-2H7v2zM7 7v2h14V7H7z"/>
              </svg>
            </n-icon>
            <p>图表区域 - 这里可以集成 ECharts 或其他图表库</p>
          </div>
        </n-card>
      </n-grid-item>
      
      <n-grid-item :span="8">
        <n-card title="最近活动">
          <n-list>
            <n-list-item>
              <n-thing title="用户注册" description="新用户 john_doe 注册了账号">
                <template #avatar>
                  <n-avatar round size="small" src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg" />
                </template>
              </n-thing>
            </n-list-item>
            <n-list-item>
              <n-thing title="订单完成" description="订单 #12345 已完成支付">
                <template #avatar>
                  <n-avatar round size="small" src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg" />
                </template>
              </n-thing>
            </n-list-item>
            <n-list-item>
              <n-thing title="系统更新" description="系统已更新到最新版本">
                <template #avatar>
                  <n-avatar round size="small" src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg" />
                </template>
              </n-thing>
            </n-list-item>
          </n-list>
        </n-card>
      </n-grid-item>
    </n-grid>
  </div>
</template>

<script setup lang="ts">
// 仪表板组件
</script>

<style scoped>
.dashboard-container {
  padding: 24px;
}

.stat-card {
  display: flex;
  align-items: center;
  gap: 16px;
}

.stat-info {
  flex: 1;
}

.stat-number {
  font-size: 24px;
  font-weight: bold;
  color: #333;
}

.stat-label {
  font-size: 14px;
  color: #666;
  margin-top: 4px;
}

.chart-placeholder {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 300px;
  color: #999;
}

.chart-placeholder p {
  margin-top: 16px;
  font-size: 14px;
}
</style> 