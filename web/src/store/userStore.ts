import { defineStore } from 'pinia'
import { ref } from 'vue'

export interface User {
  id: number
  username: string
  email: string
  avatar?: string
}

export const useUserStore = defineStore('user', () => {
  const user = ref<User | null>(null)
  const token = ref<string | null>(null)
  const isLoggedIn = ref(false)

  // 登录
  const login = (userData: User, userToken: string) => {
    user.value = userData
    token.value = userToken
    isLoggedIn.value = true
    localStorage.setItem('user', JSON.stringify(userData))
    localStorage.setItem('token', userToken)
  }

  // 登出
  const logout = () => {
    user.value = null
    token.value = null
    isLoggedIn.value = false
    localStorage.removeItem('user')
    localStorage.removeItem('token')
  }

  // 初始化用户状态
  const initUser = () => {
    const savedUser = localStorage.getItem('user')
    const savedToken = localStorage.getItem('token')
    if (savedUser && savedToken) {
      user.value = JSON.parse(savedUser)
      token.value = savedToken
      isLoggedIn.value = true
    }
  }

  return {
    user,
    token,
    isLoggedIn,
    login,
    logout,
    initUser
  }
})
