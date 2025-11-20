import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'
import App from './App.vue'
import Home from './pages/Home.vue'
import Auth from './pages/Auth.vue'
import Docs from './pages/Docs.vue'
import OpenAPI from './pages/OpenAPI.vue'
import Metrics from './pages/Metrics.vue'
import ConsoleDashboard from './pages/console/Dashboard.vue'
import ConsoleUsage from './pages/console/Usage.vue'
import ConsoleAPIKeys from './pages/console/APIKeys.vue'
import ConsoleWebhooks from './pages/console/Webhooks.vue'
import ConsoleBilling from './pages/console/Billing.vue'
import '../../packages/ui/theme.css'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', component: Home },
    { path: '/login', component: Auth },
    { path: '/docs', component: Docs },
    { path: '/openapi', component: OpenAPI },
    { path: '/metrics', component: Metrics },
    { path: '/console', component: ConsoleDashboard, meta: { requireAuth: true } },
    { path: '/console/usage', component: ConsoleUsage, meta: { requireAuth: true } },
    { path: '/console/apikeys', component: ConsoleAPIKeys, meta: { requireAuth: true } },
    { path: '/console/webhooks', component: ConsoleWebhooks, meta: { requireAuth: true } },
    { path: '/console/billing', component: ConsoleBilling, meta: { requireAuth: true } },
  ]
})

function hasToken(){ return !!localStorage.getItem('jwt') }
router.beforeEach((to, _from, next)=>{ if(to.meta?.requireAuth && !hasToken()) return next('/login'); next() })

createApp(App).use(router).mount('#app')