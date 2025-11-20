import { createRouter, createWebHistory } from 'vue-router'
import { isLoggedIn } from '../utils/auth.js'

// 页面组件
import Login from '../pages/Login.vue'
import Dashboard from '../pages/Dashboard.vue'
import Users from '../pages/Users.vue'
import Tenants from '../pages/Tenants.vue'
import ProjectConfig from '../pages/ProjectConfig.vue'
import Plugins from '../pages/Plugins.vue'
import Hooks from '../pages/Hooks.vue'
import OpsTools from '../pages/OpsTools.vue'
import InterfaceDesigner from '../pages/InterfaceDesigner.vue'
import Forbidden from '../pages/Forbidden.vue'

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: Login,
    meta: { requiresAuth: false }
  },
  {
    path: '/',
    name: 'Dashboard',
    component: Dashboard,
    meta: { requiresAuth: true, title: '仪表板' }
  },
  {
    path: '/users',
    name: 'Users',
    component: Users,
    meta: { requiresAuth: true, title: '用户管理', permission: 'user.read' }
  },
  {
    path: '/tenants',
    name: 'Tenants',
    component: Tenants,
    meta: { requiresAuth: true, title: '租户管理', permission: 'tenant.read' }
  },
  {
    path: '/config',
    name: 'ProjectConfig',
    component: ProjectConfig,
    meta: { requiresAuth: true, title: '项目配置', permission: 'config.read' }
  },
  {
    path: '/plugins',
    name: 'Plugins',
    component: Plugins,
    meta: { requiresAuth: true, title: '插件管理', permission: 'plugin.read' }
  },
  {
    path: '/hooks',
    name: 'Hooks',
    component: Hooks,
    meta: { requiresAuth: true, title: '钩子检视', permission: 'hook.read' }
  },
  {
    path: '/ops-tools',
    name: 'OpsTools',
    component: OpsTools,
    meta: { requiresAuth: true, title: '运维工具', permission: 'system.read' }
  },
  {
    path: '/interface-designer',
    name: 'InterfaceDesigner',
    component: InterfaceDesigner,
    meta: { requiresAuth: true, title: '接口设计器', permission: 'interface.write' }
  },
  {
    path: '/forbidden',
    name: 'Forbidden',
    component: Forbidden,
    meta: { requiresAuth: false }
  },
  {
    path: '/:pathMatch(.*)*',
    redirect: '/'
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// 路由守卫
router.beforeEach((to, from, next) => {
  // 设置页面标题
  if (to.meta.title) {
    document.title = `${to.meta.title} - 接口配置化系统`
  }

  // 检查是否需要认证
  if (to.meta.requiresAuth && !isLoggedIn()) {
    next('/login')
    return
  }

  // 检查权限
  if (to.meta.permission) {
    const { hasPermission } = require('../utils/auth.js')
    if (!hasPermission(to.meta.permission)) {
      next('/forbidden')
      return
    }
  }

  // 如果已经登录但访问登录页，重定向到首页
  if (to.path === '/login' && isLoggedIn()) {
    next('/')
    return
  }

  next()
})

export default router