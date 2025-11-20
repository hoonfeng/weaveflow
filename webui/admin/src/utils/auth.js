// 认证工具函数
import { ref, reactive, readonly } from 'vue'

// 存储token到localStorage
export function setToken(token) {
  if (typeof window !== 'undefined') {
    localStorage.setItem('ifaceconf_token', token)
  }
}

// 从localStorage获取token
export function getToken() {
  if (typeof window !== 'undefined') {
    return localStorage.getItem('ifaceconf_token')
  }
  return null
}

// 移除token
export function removeToken() {
  if (typeof window !== 'undefined') {
    localStorage.removeItem('ifaceconf_token')
  }
}

// 检查是否已登录
export function isLoggedIn() {
  const token = getToken()
  if (!token) return false
  
  // 检查token是否过期
  try {
    const payload = JSON.parse(atob(token.split('.')[1]))
    const now = Math.floor(Date.now() / 1000)
    if (typeof payload.exp === 'number') {
      return payload.exp > now
    }
    // 无过期字段时视为有效（兼容后端未设置exp的令牌）
    return true
  } catch (error) {
    return false
  }
}

// 获取用户信息
export function getUserInfo() {
  const token = getToken()
  if (!token) return null
  
  try {
    const payload = JSON.parse(atob(token.split('.')[1]))
    return {
      id: payload.uid || payload.sub,
      username: payload.username,
      role: payload.role,
      roles: payload.roles || [],
      perms: payload.perms || [],
      exp: payload.exp
    }
  } catch (error) {
    return null
  }
}

// 检查用户权限
export function hasPermission(permission) {
  const userInfo = getUserInfo()
  if (!userInfo) return false
  
  // 超级管理员拥有所有权限
  if (userInfo.role === 'superadmin') return true
  
  // 这里可以根据实际权限系统进行扩展
  // 例如检查用户角色或权限列表
  const userPermissions = {
    admin: ['user.read', 'user.write', 'config.read'],
    user: ['user.read']
  }
  
  const permissions = userPermissions[userInfo.role] || []
  return permissions.includes(permission)
}

// 登录处理
export async function login(credentials) {
  try {
    const response = await fetch('/api/auth/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(credentials)
    })
    
    if (!response.ok) {
      throw new Error('登录失败')
    }
    
    const data = await response.json()
    const token = (data && data.token) || (data && data.data && data.data.token)
    if (token) {
      setToken(token)
      return { success: true, token }
    }
    throw new Error('无效的响应格式')
  } catch (error) {
    console.error('登录错误:', error)
    throw error
  }
}

// 登出处理
export function logout() {
  removeToken()
  // 跳转到登录页
  if (typeof window !== 'undefined') {
    window.location.href = '/login'
  }
}

// 自动刷新token
export async function refreshToken() {
  try {
    const response = await fetch('/api/auth/refresh', {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${getToken()}`
      }
    })
    
    if (response.ok) {
      const data = await response.json()
      const token = (data && data.token) || (data && data.data && data.data.token)
      if (token) { setToken(token); return true }
    }
  } catch (error) {
    console.error('刷新token失败:', error)
  }
  
  // 刷新失败，需要重新登录
  logout()
  return false
}

// 全局认证状态（避免每次useAuth创建独立ref导致不同组件不共享状态）
const userState = ref(getUserInfo())
const isAuthenticatedState = ref(isLoggedIn())

// Vue 3 Composition API 认证钩子（返回共享状态）
export function useAuth() {
  const loginUser = async (credentials) => {
    try {
      const result = await login(credentials)
      userState.value = getUserInfo()
      isAuthenticatedState.value = true
      return result
    } catch (error) {
      throw error
    }
  }

  const logoutUser = () => {
    logout()
    userState.value = null
    isAuthenticatedState.value = false
  }

  return {
    user: readonly(userState),
    currentUser: readonly(userState),
    isAuthenticated: readonly(isAuthenticatedState),
    login: loginUser,
    logout: logoutUser,
    hasPermission
  }
}

// 认证管理器（兼容旧版本）
export const authManager = {
  setToken,
  getToken,
  removeToken,
  isLoggedIn,
  checkAuth: () => isLoggedIn(),
  getUserInfo,
  isAdmin: () => {
    const info = getUserInfo()
    const roles = (info && info.roles) || []
    return !!info && (roles.includes('admin') || roles.includes('superadmin') || info.role === 'admin' || info.role === 'superadmin')
  },
  hasPermission,
  login,
  logout,
  refreshToken
}