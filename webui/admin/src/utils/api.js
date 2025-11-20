import axios from 'axios'
import { getToken } from './auth.js'

// 创建axios实例
const api = axios.create({
  baseURL: '/api',
  timeout: 30000,
  headers: {
    'Content-Type': 'application/json'
  }
})

// 请求拦截器
api.interceptors.request.use(
  (config) => {
    const token = getToken()
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
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
      // token过期，跳转到登录页
      window.location.href = '/login'
    }
    return Promise.reject(error)
  }
)

// 仪表盘概览API
export const overviewApi = {
  // 获取仪表盘概览数据
  getDashboard: () => api.get('/admin/dashboard/overview'),
  
  // 获取概览统计
  getOverview: () => api.get('/admin/overview'),
  // 获取概览趋势（近12月用户与订单）
  getTrends: () => api.get('/admin/overview/trends'),
  // 获取支付统计（状态与渠道分布）
  getPaymentsStats: () => api.get('/admin/overview/payments_stats')
}

// 系统管理API
export const systemApi = {
  // 健康检查
  health: () => api.get('/admin/health'),
  
  // 接口规范检查
  lint: () => api.get('/admin/lint'),
  
  // 重载配置
  reload: () => api.post('/admin/reload'),
  
  // 模型迁移
  migrate: () => api.post('/admin/migrate'),
  
  // 队列状态
  queueStatus: () => api.get('/admin/queue/status'),
  
  // 清理Nonces
  purgeNonces: (ttlSeconds) => api.post('/admin/nonces/purge', null, { params: { ttlSeconds } }),
  
  // 远程调试
  remoteDebug: () => api.get('/admin/remote/debug'),
  
  // 远程调用测试
  remoteTest: () => api.get('/admin/remote/test')
}

// 插件管理API
export const pluginsApi = {
  // 获取外部插件列表
  getExternal: () => api.get('/admin/plugins/external'),
  
  // 启用/禁用插件
  enable: (data) => api.post('/admin/plugins/enable', data),
  
  // 添加插件
  add: (data) => api.post('/admin/plugins/add', data),
  
  // 移除插件
  remove: (data) => api.post('/admin/plugins/remove', data),
  
  // 停止插件
  stop: (data) => api.post('/admin/plugins/stop', data),
  
  // 启动插件
  start: (data) => api.post('/admin/plugins/start', data),
  
  // 重启插件
  restart: (data) => api.post('/admin/plugins/restart', data)
}

// 钩子管理API
export const hooksApi = {
  // 检视钩子
  inspect: (data) => api.post('/admin/hooks/inspect', data)
}

// 配置管理API
export const configApi = {
  // 获取配置
  get: () => api.get('/admin/config'),
  
  // 保存配置
  save: (data) => api.post('/admin/config', data)
}

// 用户管理API
export const usersApi = {
  // 获取用户列表 - 实际接口为GET /api/admin/users
  list: (params) => api.get('/admin/users', { params }),
  
  // 创建用户 - 实际接口为POST /api/admin/users
  create: (data) => api.post('/admin/users', data),
  
  // 更新用户 - 实际接口为PUT /api/admin/users/{id}
  update: (id, data) => api.put(`/admin/users/${id}`, data),
  
  // 删除用户 - 实际接口为DELETE /api/admin/users/{id}
  delete: (id) => api.delete(`/admin/users/${id}`),
  
  // 获取用户详情 - 实际接口为GET /api/admin/users/{id}
  get: (id) => api.get(`/admin/users/${id}`)
}

// 租户管理API
export const tenantsApi = {
  // 获取租户列表 - 实际接口为GET /api/admin/tenants
  list: (params) => api.get('/admin/tenants', { params }),
  
  // 创建租户 - 实际接口为POST /api/admin/tenants
  create: (data) => api.post('/admin/tenants', data),
  
  // 更新租户 - 实际接口为PUT /api/admin/tenants/{id}
  update: (id, data) => api.put(`/admin/tenants/${id}`, data),
  
  // 删除租户 - 实际接口为DELETE /api/admin/tenants/{id}
  delete: (id) => api.delete(`/admin/tenants/${id}`),
  
  // 获取租户详情 - 实际接口为GET /api/admin/tenants/{id}
  get: (id) => api.get(`/admin/tenants/${id}`)
}

// 接口管理API
export const interfacesApi = {
  // 获取接口列表
  list: (params) => api.get('/admin/interfaces', { params }),
  
  // 获取接口详情
  get: (id) => api.get(`/admin/interfaces/${id}`),
  
  // 创建接口
  create: (data) => api.post('/admin/interfaces', data),
  
  // 更新接口
  update: (id, data) => api.put(`/admin/interfaces/${id}`, data),
  
  // 删除接口
  delete: (id) => api.delete(`/admin/interfaces/${id}`),
  
  // 测试接口
  test: (id, data) => api.post(`/admin/interfaces/${id}/test`, data)
}

// 文件上传API
export const uploadApi = {
  // 上传文件
  upload: (formData) => api.post('/files/upload', formData, {
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })
}

// 认证API
export const authApi = {
  // 登录
  login: (data) => api.post('/auth/login', data),
  
  // 登出
  logout: () => api.post('/auth/logout'),
  
  // 刷新token
  refresh: () => api.post('/auth/refresh'),
  
  // 获取用户信息
  getProfile: () => api.get('/auth/profile')
}

// 支付相关API - 添加缺失的导出
export const payApi = {
  // 支付对账
  reconcile: (data) => api.post('/admin/pay/reconcile', data)
}

export const payProvidersApi = {
  // 获取支付提供商列表
  list: () => api.get('/admin/pay/providers'),
  
  // 添加支付提供商
  create: (data) => api.post('/admin/pay/providers', data),
  
  // 更新支付提供商
  update: (id, data) => api.put(`/admin/pay/providers/${id}`, data),
  
  // 删除支付提供商
  delete: (id) => api.delete(`/admin/pay/providers/${id}`)
}

export const paymentsApi = {
  // 获取支付列表
  list: (params) => api.get('/admin/payments', { params }),
  
  // 获取支付详情
  get: (id) => api.get(`/admin/payments/${id}`),
  
  // 创建支付
  create: (data) => api.post('/admin/payments', data),
  
  // 更新支付
  update: (id, data) => api.put(`/admin/payments/${id}`, data)
}

export const rolesApi = {
  // 获取角色列表
  list: () => api.get('/admin/roles'),
  
  // 创建角色
  create: (data) => api.post('/admin/roles', data),
  
  // 更新角色
  update: (id, data) => api.put(`/admin/roles/${id}`, data),
  
  // 删除角色
  delete: (id) => api.delete(`/admin/roles/${id}`)
}

export const subscriptionsApi = {
  // 获取订阅列表
  list: (params) => api.get('/admin/subscriptions', { params }),
  
  // 获取订阅详情
  get: (id) => api.get(`/admin/subscriptions/${id}`),
  
  // 创建订阅
  create: (data) => api.post('/admin/subscriptions', data),
  
  // 更新订阅
  update: (id, data) => api.put(`/admin/subscriptions/${id}`, data),
  
  // 取消订阅
  cancel: (id) => api.post(`/admin/subscriptions/${id}/cancel`)
}

export default api