import axios from 'axios'
import { useUserStore } from '../store/userStore'
import router from '../router'

// 创建axios实例
const api = axios.create({
  baseURL: '',
  timeout: 1000000,
})

// 请求拦截器
api.interceptors.request.use(
  (config) => {
    const userStore = useUserStore()
    if (userStore.token) {
      config.headers.Authorization = `Bearer ${userStore.token}`
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// 响应拦截器
api.interceptors.response.use(
  (response) => {
    return response.data
  },
  (error) => {
    if (error.response?.status === 401) {
      const userStore = useUserStore()
      userStore.logout()
      router.push('/login')
    }
    return Promise.reject(error)
  }
)

// 类型定义
export interface Channel {
  id: number
  user_id: number
  name: string
  type: string
  credentials: Record<string, any>
  created_at: string
  updated_at: string
}

export interface Topic {
  id: number
  user_id: number
  name: string
  webhook_key: string
  sending_strategy: string
  execution_mode: string
  description: string
  created_at: string
  updated_at: string
}

export interface Routing {
  topic_id: number
  channel_id: number
  priority: number
  variable_mappings: Record<string, any>
  message_template: string
  created_at: string
  updated_at: string
}

export interface CreateChannelRequest {
  name: string
  type: string
  credentials: Record<string, any>
}

export interface UpdateChannelRequest {
  name: string
  type: string
  credentials: Record<string, any>
}

export interface CreateTopicRequest {
  name: string
  sending_strategy: string
  execution_mode: string
  description: string
}

export interface UpdateTopicRequest {
  name: string
  sending_strategy: string
  execution_mode: string
  description: string
}

export interface CreateRoutingRequest {
  topic_id: number
  channel_id: number
  priority: number
  variable_mappings: Record<string, any>
  message_template: string
}

export interface UpdateRoutingRequest {
  priority: number
  variable_mappings: Record<string, any>
  message_template: string
}

// 通道API
export const channelApi = {
  // 获取通道列表
  getChannels: () => api.get<Channel[]>('/api/channels'),
  
  // 获取单个通道
  getChannel: (id: number) => api.get<Channel>(`/api/channels/${id}`),
  
  // 创建通道
  createChannel: (data: CreateChannelRequest) => api.post<Channel>('/api/channels', data),
  
  // 更新通道
  updateChannel: (id: number, data: UpdateChannelRequest) => api.put<Channel>(`/api/channels/${id}`, data),
  
  // 删除通道
  deleteChannel: (id: number) => api.delete(`/api/channels/${id}`),
  
  // 获取通道的路由
  getChannelRoutings: (id: number) => api.get<Routing[]>(`/api/channels/${id}/routings`),
}

// 主题API
export const topicApi = {
  // 获取主题列表
  getTopics: () => api.get<Topic[]>('/api/topics'),
  
  // 获取单个主题
  getTopic: (id: number) => api.get<Topic>(`/api/topics/${id}`),
  
  // 创建主题
  createTopic: (data: CreateTopicRequest) => api.post<Topic>('/api/topics', data),
  
  // 更新主题
  updateTopic: (id: number, data: UpdateTopicRequest) => api.put<Topic>(`/api/topics/${id}`, data),
  
  // 删除主题
  deleteTopic: (id: number) => api.delete(`/api/topics/${id}`),
  
  // 重新生成Webhook Key
  regenerateWebhookKey: (id: number) => api.post<Topic>(`/api/topics/${id}/regenerate-key`),
  
  // 获取主题的路由
  getTopicRoutings: (id: number) => api.get<Routing[]>(`/api/topics/${id}/routings`),
}

// 路由API
export const routingApi = {
  // 创建路由
  createRouting: (data: CreateRoutingRequest) => api.post<Routing>('/api/routings', data),
  
  // 更新路由
  updateRouting: (topicId: number, channelId: number, data: UpdateRoutingRequest) => 
    api.put<Routing>(`/api/routings/${topicId}/${channelId}`, data),
  
  // 删除路由
  deleteRouting: (topicId: number, channelId: number) => 
    api.delete(`/api/routings/${topicId}/${channelId}`),
}

export default api 