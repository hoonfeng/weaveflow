<template>
  <div class="admin-app">
    <!-- 登录页面 -->
    <div v-if="!isAuthenticated || $route.path === '/login'" class="login-layout">
      <router-view />
    </div>
    
    <!-- 主布局 -->
    <div v-else class="main-layout">
      <!-- 侧边导航 -->
      <aside class="sidebar" :class="{ 'sidebar-collapsed': sidebarCollapsed }">
        <div class="sidebar-header">
          <div class="logo">
            <div class="logo-icon">
              <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z" stroke="currentColor" stroke-width="2"/>
                <path d="M9 12l2 2 4-4" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
              </svg>
            </div>
            <span class="logo-text" v-if="!sidebarCollapsed">IFaceConf</span>
          </div>
          <button class="sidebar-toggle" @click="sidebarCollapsed = !sidebarCollapsed">
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M3 12h18M3 6h18M3 18h18" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
            </svg>
          </button>
        </div>

        <nav class="sidebar-nav">
          <div class="nav-section">
            <div class="nav-section-label" v-if="!sidebarCollapsed">概览</div>
            <router-link to="/" class="nav-link" active-class="active">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M3 9l9-7 9 7v11a2 2 0 01-2 2H5a2 2 0 01-2-2V9z" stroke="currentColor" stroke-width="2"/>
                <path d="M9 22V12h6v10" stroke="currentColor" stroke-width="2"/>
              </svg>
              <span v-if="!sidebarCollapsed">仪表盘</span>
            </router-link>
          </div>

          <div class="nav-section">
            <div class="nav-section-label" v-if="!sidebarCollapsed">用户管理</div>
            <router-link to="/users" class="nav-link" active-class="active">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M17 21v-2a4 4 0 00-4-4H5a4 4 0 00-4 4v2" stroke="currentColor" stroke-width="2"/>
                <circle cx="9" cy="7" r="4" stroke="currentColor" stroke-width="2"/>
                <path d="M23 21v-2a4 4 0 00-3-3.87" stroke="currentColor" stroke-width="2"/>
                <path d="M16 3.13a4 4 0 010 7.75" stroke="currentColor" stroke-width="2"/>
              </svg>
              <span v-if="!sidebarCollapsed">用户管理</span>
            </router-link>
            <router-link to="/roles" class="nav-link" active-class="active">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M16 21v-2a4 4 0 00-4-4H5a4 4 0 00-4 4v2" stroke="currentColor" stroke-width="2"/>
                <circle cx="8.5" cy="7" r="4" stroke="currentColor" stroke-width="2"/>
                <path d="M20 8v6" stroke="currentColor" stroke-width="2"/>
                <path d="M23 11h-6" stroke="currentColor" stroke-width="2"/>
              </svg>
              <span v-if="!sidebarCollapsed">角色管理</span>
            </router-link>
            <router-link to="/permissions" class="nav-link" active-class="active">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                <rect x="3" y="11" width="18" height="11" rx="2" ry="2" stroke="currentColor" stroke-width="2"/>
                <path d="M7 11V7a5 5 0 0110 0v4" stroke="currentColor" stroke-width="2"/>
              </svg>
              <span v-if="!sidebarCollapsed">权限管理</span>
            </router-link>
          </div>

          <div class="nav-section">
            <div class="nav-section-label" v-if="!sidebarCollapsed">租户管理</div>
            <router-link to="/tenants" class="nav-link" active-class="active">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M19 21l-7-5-7 5V5a2 2 0 012-2h10a2 2 0 012 2v16z" stroke="currentColor" stroke-width="2"/>
              </svg>
              <span v-if="!sidebarCollapsed">租户列表</span>
            </router-link>
            <router-link to="/tenants/users" class="nav-link" active-class="active">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M17 21v-2a4 4 0 00-4-4H5a4 4 0 00-4 4v2" stroke="currentColor" stroke-width="2"/>
                <circle cx="9" cy="7" r="4" stroke="currentColor" stroke-width="2"/>
              </svg>
              <span v-if="!sidebarCollapsed">成员管理</span>
            </router-link>
            <router-link to="/tenants/apikeys" class="nav-link" active-class="active">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M15 7l-6 6m0-6l6 6" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
                <circle cx="12" cy="12" r="10" stroke="currentColor" stroke-width="2"/>
              </svg>
              <span v-if="!sidebarCollapsed">API Keys</span>
            </router-link>
          </div>

          <div class="nav-section">
            <div class="nav-section-label" v-if="!sidebarCollapsed">业务管理</div>
            <router-link to="/orders" class="nav-link" active-class="active">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M6 2L3 6v14a2 2 0 002 2h14a2 2 0 002-2V6l-3-4H6z" stroke="currentColor" stroke-width="2"/>
                <path d="M3 6h18" stroke="currentColor" stroke-width="2"/>
                <path d="M16 10a4 4 0 11-8 0" stroke="currentColor" stroke-width="2"/>
              </svg>
              <span v-if="!sidebarCollapsed">订单管理</span>
            </router-link>
            <router-link to="/payments" class="nav-link" active-class="active">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                <rect x="1" y="4" width="22" height="16" rx="2" ry="2" stroke="currentColor" stroke-width="2"/>
                <path d="M1 10h22" stroke="currentColor" stroke-width="2"/>
              </svg>
              <span v-if="!sidebarCollapsed">支付管理</span>
            </router-link>
          </div>

          <div class="nav-section">
            <div class="nav-section-label" v-if="!sidebarCollapsed">系统管理</div>
            <router-link to="/project" class="nav-link" active-class="active">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M14 2H6a2 2 0 00-2 2v16a2 2 0 002 2h12a2 2 0 002-2V8z" stroke="currentColor" stroke-width="2"/>
                <path d="M14 2v6h6" stroke="currentColor" stroke-width="2"/>
                <path d="M16 13H8" stroke="currentColor" stroke-width="2"/>
                <path d="M16 17H8" stroke="currentColor" stroke-width="2"/>
                <path d="M10 9H8" stroke="currentColor" stroke-width="2"/>
              </svg>
              <span v-if="!sidebarCollapsed">项目配置</span>
            </router-link>
            <router-link to="/plugins" class="nav-link" active-class="active">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z" stroke="currentColor" stroke-width="2"/>
              </svg>
              <span v-if="!sidebarCollapsed">插件管理</span>
            </router-link>
            <router-link to="/ops" class="nav-link" active-class="active">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                <circle cx="12" cy="12" r="10" stroke="currentColor" stroke-width="2"/>
                <path d="M12 6v6l4 2" stroke="currentColor" stroke-width="2"/>
              </svg>
              <span v-if="!sidebarCollapsed">运维工具</span>
            </router-link>
          </div>
        </nav>

        <div class="sidebar-footer">
          <div class="user-info" v-if="!sidebarCollapsed">
            <div class="user-avatar">{{ currentUser?.username?.charAt(0)?.toUpperCase() }}</div>
            <div class="user-details">
              <div class="user-name">{{ currentUser?.username }}</div>
              <div class="user-role">{{ currentUser?.role }}</div>
            </div>
          </div>
          <button class="logout-btn" @click="handleLogout">
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M9 21H5a2 2 0 01-2-2V5a2 2 0 012-2h4" stroke="currentColor" stroke-width="2"/>
              <path d="M16 17l5-5-5-5" stroke="currentColor" stroke-width="2"/>
              <path d="M21 12H9" stroke="currentColor" stroke-width="2"/>
            </svg>
            <span v-if="!sidebarCollapsed">退出</span>
          </button>
        </div>
      </aside>

      <!-- 主内容区域 -->
      <main class="main-content" :class="{ 'sidebar-collapsed': sidebarCollapsed }">
        <header class="main-header">
          <div class="header-left">
            <h1 class="page-title">{{ pageTitle }}</h1>
            <p class="page-description">{{ pageDescription }}</p>
          </div>
          <div class="header-right">
            <div class="header-actions">
              <button class="btn btn-outline" @click="toggleTheme">
                <svg width="16" height="16" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <path d="M21 12.79A9 9 0 1111.21 3 7 7 0 0021 12.79z" stroke="currentColor" stroke-width="2"/>
                </svg>
                主题
              </button>
            </div>
          </div>
        </header>

        <div class="content-area">
          <router-view />
        </div>
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuth } from './utils/auth.js'

const router = useRouter()
const route = useRoute()
const { isAuthenticated, currentUser, logout } = useAuth()

const sidebarCollapsed = ref(false)

// 页面标题映射
const pageTitles = {
  '/': '仪表盘',
  '/users': '用户管理',
  '/roles': '角色管理',
  '/permissions': '权限管理',
  '/tenants': '租户管理',
  '/tenants/users': '租户成员',
  '/tenants/apikeys': 'API密钥',
  '/orders': '订单管理',
  '/payments': '支付管理',
  '/project': '项目配置',
  '/plugins': '插件管理',
  '/ops': '运维工具'
}

const pageDescriptions = {
  '/': '系统概览和关键指标',
  '/users': '管理系统用户账户',
  '/roles': '管理用户角色和权限分配',
  '/permissions': '配置系统权限策略',
  '/tenants': '管理多租户配置',
  '/tenants/users': '管理租户下的用户成员',
  '/tenants/apikeys': '配置租户API访问密钥',
  '/orders': '查看和处理业务订单',
  '/payments': '管理支付记录和状态',
  '/project': '配置项目全局设置',
  '/plugins': '管理插件和扩展功能',
  '/ops': '系统运维和监控工具'
}

const pageTitle = computed(() => pageTitles[route.path] || '管理后台')
const pageDescription = computed(() => pageDescriptions[route.path] || '')

const handleLogout = () => {
  if (confirm('确定要退出登录吗？')) {
    logout()
    router.push('/login')
  }
}

const toggleTheme = () => {
  // 主题切换逻辑
  document.documentElement.classList.toggle('light-theme')
}
</script>

<style scoped>
.admin-app {
  min-height: 100vh;
  background: var(--bg-primary);
}

.login-layout {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: radial-gradient(ellipse at top, var(--bg-secondary) 0%, var(--bg-primary) 100%);
}

.main-layout {
  display: flex;
  min-height: 100vh;
}

/* 侧边栏样式 */
.sidebar {
  width: 280px;
  background: var(--bg-surface);
  border-right: 1px solid var(--border-light);
  display: flex;
  flex-direction: column;
  transition: width var(--transition-normal);
}

.sidebar-collapsed {
  width: 80px;
}

.sidebar-header {
  padding: var(--space-lg);
  border-bottom: 1px solid var(--border-light);
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.logo {
  display: flex;
  align-items: center;
  gap: var(--space-sm);
}

.logo-icon {
  width: 32px;
  height: 32px;
  background: var(--gradient-primary);
  border-radius: var(--radius-md);
  display: flex;
  align-items: center;
  justify-content: center;
}

.logo-icon svg {
  color: white;
}

.logo-text {
  font-size: 1.25rem;
  font-weight: 700;
  background: var(--gradient-primary);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.sidebar-toggle {
  background: none;
  border: none;
  color: var(--text-secondary);
  cursor: pointer;
  padding: var(--space-xs);
  border-radius: var(--radius-sm);
  transition: all var(--transition-fast);
}

.sidebar-toggle:hover {
  background: var(--bg-card);
  color: var(--text-primary);
}

.sidebar-nav {
  flex: 1;
  padding: var(--space-lg);
  overflow-y: auto;
}

.nav-section {
  margin-bottom: var(--space-xl);
}

.nav-section-label {
  font-size: 0.75rem;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  color: var(--text-muted);
  margin-bottom: var(--space-sm);
}

.nav-link {
  display: flex;
  align-items: center;
  gap: var(--space-md);
  padding: var(--space-md);
  color: var(--text-secondary);
  text-decoration: none;
  border-radius: var(--radius-md);
  transition: all var(--transition-fast);
  margin-bottom: var(--space-xs);
}

.nav-link:hover {
  background: var(--bg-card);
  color: var(--text-primary);
}

.nav-link.active {
  background: var(--primary-500);
  color: white;
}

.nav-link svg {
  flex-shrink: 0;
}

.sidebar-footer {
  padding: var(--space-lg);
  border-top: 1px solid var(--border-light);
}

.user-info {
  display: flex;
  align-items: center;
  gap: var(--space-sm);
  margin-bottom: var(--space-md);
}

.user-avatar {
  width: 40px;
  height: 40px;
  background: var(--gradient-primary);
  border-radius: var(--radius-full);
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
  color: white;
}

.user-details {
  flex: 1;
}

.user-name {
  font-weight: 600;
  margin-bottom: 2px;
}

.user-role {
  font-size: 0.75rem;
  color: var(--text-muted);
}

.logout-btn {
  width: 100%;
  display: flex;
  align-items: center;
  gap: var(--space-sm);
  padding: var(--space-md);
  background: rgba(239, 68, 68, 0.1);
  border: 1px solid rgba(239, 68, 68, 0.2);
  border-radius: var(--radius-md);
  color: var(--error);
  cursor: pointer;
  transition: all var(--transition-fast);
}

.logout-btn:hover {
  background: rgba(239, 68, 68, 0.2);
}

/* 主内容区域样式 */
.main-content {
  flex: 1;
  transition: margin-left var(--transition-normal);
}

.main-content.sidebar-collapsed {
  margin-left: 80px;
}

.main-header {
  padding: var(--space-xl);
  border-bottom: 1px solid var(--border-light);
  background: var(--bg-surface);
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.page-title {
  font-size: 1.875rem;
  font-weight: 700;
  margin-bottom: var(--space-xs);
}

.page-description {
  color: var(--text-secondary);
}

.header-actions {
  display: flex;
  gap: var(--space-sm);
}

.content-area {
  padding: var(--space-xl);
  max-height: calc(100vh - 120px);
  overflow-y: auto;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .sidebar {
    position: fixed;
    top: 0;
    left: 0;
    height: 100vh;
    z-index: 1000;
    transform: translateX(-100%);
  }
  
  .sidebar.mobile-open {
    transform: translateX(0);
  }
  
  .main-content {
    margin-left: 0 !important;
  }
}
</style>