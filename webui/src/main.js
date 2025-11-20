import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'
import App from './App.vue'
import Home from './pages/Home.vue'
import Docs from './pages/Docs.vue'
import OpenAPI from './pages/OpenAPI.vue'
import Metrics from './pages/Metrics.vue'
import Login from './pages/Login.vue'
import Reverse from './pages/Reverse.vue'
import Upload from './pages/Upload.vue'
import './styles.css'
import AdminDashboard from './pages/admin/AdminDashboard.vue'
import AdminUsers from './pages/admin/AdminUsers.vue'
import AdminRoles from './pages/admin/AdminRoles.vue'
import AdminTenants from './pages/admin/AdminTenants.vue'
import AdminPlans from './pages/admin/AdminPlans.vue'
import AdminSubscriptions from './pages/admin/AdminSubscriptions.vue'
import AdminApiKeys from './pages/admin/AdminApiKeys.vue'
import AdminOrders from './pages/admin/AdminOrders.vue'
import AdminPayments from './pages/admin/AdminPayments.vue'
import AdminTools from './pages/admin/AdminTools.vue'
import ConsoleDashboard from './pages/console/ConsoleDashboard.vue'
import ConsoleUsage from './pages/console/ConsoleUsage.vue'
import ConsoleWebhooks from './pages/console/ConsoleWebhooks.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', redirect: '/admin' },
    { path: '/docs', component: Docs },
    { path: '/openapi', component: OpenAPI },
    { path: '/metrics', component: Metrics },
    { path: '/login', component: Login },
    { path: '/reverse', component: Reverse },
    { path: '/upload', component: Upload },
    { path: '/admin', component: AdminDashboard, meta: { requireAdmin: true } },
    { path: '/admin/users', component: AdminUsers, meta: { requireAdmin: true } },
    { path: '/admin/roles', component: AdminRoles, meta: { requireAdmin: true } },
    { path: '/admin/tenants', component: AdminTenants, meta: { requireAdmin: true } },
    { path: '/admin/plans', component: AdminPlans, meta: { requireAdmin: true } },
    { path: '/admin/subscriptions', component: AdminSubscriptions, meta: { requireAdmin: true } },
    { path: '/admin/apikeys', component: AdminApiKeys, meta: { requireAdmin: true } },
    { path: '/admin/orders', component: AdminOrders, meta: { requireAdmin: true } },
    { path: '/admin/payments', component: AdminPayments, meta: { requireAdmin: true } },
    { path: '/admin/tools', component: AdminTools, meta: { requireAdmin: true } },
    { path: '/console', component: ConsoleDashboard, meta: { requireAuth: true } },
    { path: '/console/usage', component: ConsoleUsage, meta: { requireAuth: true } },
    { path: '/console/webhooks', component: ConsoleWebhooks, meta: { requireAuth: true } }
  ]
})

function parseJwt(t){
  if(!t) return {}
  const p = t.split('.')[1]
  if(!p) return {}
  try { return JSON.parse(atob(p.replace(/-/g,'+').replace(/_/g,'/'))) } catch { return {} }
}
function hasToken(){ return !!localStorage.getItem('jwt') }
function isAdmin(){ const c = parseJwt(localStorage.getItem('jwt')); const r = c.roles || c.role || []; if(Array.isArray(r)) return r.includes('admin'); if(typeof r==='string') return r==='admin'; return false }

router.beforeEach((to, from, next)=>{
  if(to.meta && to.meta.requireAdmin){ if(!hasToken() || !isAdmin()) return next('/login') }
  if(to.meta && to.meta.requireAuth){ if(!hasToken()) return next('/login') }
  next()
})

createApp(App).use(router).mount('#app')