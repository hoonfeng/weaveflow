import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'
import App from './App.vue'
import Login from './pages/Login.vue'
import Dashboard from './pages/Dashboard.vue'
import Users from './pages/Users.vue'
import Roles from './pages/Roles.vue'
import Permissions from './pages/Permissions.vue'
import Tenants from './pages/Tenants.vue'
import TenantsUsers from './pages/TenantsUsers.vue'
import TenantsSubscriptions from './pages/TenantsSubscriptions.vue'
import TenantLimits from './pages/TenantLimits.vue'
import TenantUsage from './pages/TenantUsage.vue'
import TenantApiKeys from './pages/TenantApiKeys.vue'
import Plans from './pages/Plans.vue'
import Subscriptions from './pages/Subscriptions.vue'
import Orders from './pages/Orders.vue'
import Payments from './pages/Payments.vue'
import ProjectConfig from './pages/ProjectConfig.vue'
import Plugins from './pages/Plugins.vue'
import PluginsInternal from './pages/PluginsInternal.vue'
import OpsTools from './pages/OpsTools.vue'
import Hooks from './pages/Hooks.vue'
import PayReconcile from './pages/PayReconcile.vue'
import PaymentChannels from './pages/PaymentChannels.vue'
import PaymentProviders from './pages/PaymentProviders.vue'
import Forbidden from './pages/Forbidden.vue'

// 导入样式
import './styles/theme.css'
import './styles/components.css'

import { authManager } from './utils/auth.js'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/login', component: Login, meta: { public: true } },
    { path: '/forbidden', component: Forbidden, meta: { public: true } },
    { path: '/', component: Dashboard, meta: { requireAuth: true, requireAdmin: true } },
    { path: '/users', component: Users, meta: { requireAuth: true, requireAdmin: true } },
    { path: '/roles', component: Roles, meta: { requireAuth: true, requireAdmin: true } },
    { path: '/permissions', component: Permissions, meta: { requireAuth: true, requireAdmin: true } },
    { path: '/tenants', component: Tenants, meta: { requireAuth: true, requireAdmin: true } },
    { path: '/tenants/users', component: TenantsUsers, meta: { requireAuth: true, requireAdmin: true } },
    { path: '/tenants/subscriptions', component: TenantsSubscriptions, meta: { requireAuth: true, requireAdmin: true } },
    { path: '/tenants/limits', component: TenantLimits, meta: { requireAuth: true, requireAdmin: true } },
    { path: '/tenants/usage', component: TenantUsage, meta: { requireAuth: true, requireAdmin: true } },
    { path: '/tenants/apikeys', component: TenantApiKeys, meta: { requireAuth: true, requireAdmin: true } },
    { path: '/plans', component: Plans, meta: { requireAuth: true, requireAdmin: true } },
    { path: '/subscriptions', component: Subscriptions, meta: { requireAuth: true, requireAdmin: true } },
    { path: '/orders', component: Orders, meta: { requireAuth: true, requireAdmin: true } },
    { path: '/payments', component: Payments, meta: { requireAuth: true, requireAdmin: true } },
    { path: '/pay/reconcile', component: PayReconcile, meta: { requireAuth: true, requireAdmin: true } },
    { path: '/pay/channels', component: PaymentChannels, meta: { requireAuth: true, requireAdmin: true } },
    { path: '/pay/providers', component: PaymentProviders, meta: { requireAuth: true, requireAdmin: true } },
    { path: '/project', component: ProjectConfig, meta: { requireAuth: true, requireAdmin: true } },
    { path: '/plugins', component: Plugins, meta: { requireAuth: true, requireAdmin: true } },
    { path: '/plugins/internal', component: PluginsInternal, meta: { requireAuth: true, requireAdmin: true } },
    { path: '/ops', component: OpsTools, meta: { requireAuth: true, requireAdmin: true } },
    { path: '/hooks', component: Hooks, meta: { requireAuth: true, requireAdmin: true } }
  ]
})

// 路由守卫 - 权限控制
router.beforeEach((to, from, next) => {
  // 公开路由直接放行
  if (to.meta.public) {
    return next()
  }

  // 检查认证状态
  const isAuthenticated = authManager.checkAuth()
  
  // 需要认证但未登录，跳转到登录页
  if (to.meta.requireAuth && !isAuthenticated) {
    return next('/login')
  }

  // 需要管理员权限但当前用户不是管理员，跳转到无权限页面
  if (to.meta.requireAdmin && !authManager.isAdmin()) {
    return next('/forbidden')
  }

  next()
})

const app = createApp(App)
app.use(router)
app.mount('#app')