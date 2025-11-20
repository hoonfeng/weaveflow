<template>
  <div class="dashboard">
    <!-- 概览统计 -->
    <div class="stats-grid">
      <div class="stat-card">
        <div class="stat-icon">
          <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M17 21v-2a4 4 0 00-4-4H5a4 4 0 00-4 4v2" stroke="currentColor" stroke-width="2"/>
            <circle cx="9" cy="7" r="4" stroke="currentColor" stroke-width="2"/>
            <path d="M23 21v-2a4 4 0 00-3-3.87" stroke="currentColor" stroke-width="2"/>
            <path d="M16 3.13a4 4 0 010 7.75" stroke="currentColor" stroke-width="2"/>
          </svg>
        </div>
        <div class="stat-content">
          <div class="stat-value">{{ stats.users || 0 }}</div>
          <div class="stat-label">用户总数</div>
        </div>
        <div class="stat-trend positive">+12%</div>
      </div>

      <div class="stat-card">
        <div class="stat-icon">
          <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M19 21l-7-5-7 5V5a2 2 0 012-2h10a2 2 0 012 2v16z" stroke="currentColor" stroke-width="2"/>
          </svg>
        </div>
        <div class="stat-content">
          <div class="stat-value">{{ stats.tenants || 0 }}</div>
          <div class="stat-label">租户数量</div>
        </div>
        <div class="stat-trend positive">+8%</div>
      </div>

      <div class="stat-card">
        <div class="stat-icon">
          <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M6 2L3 6v14a2 2 0 002 2h14a2 2 0 002-2V6l-3-4H6z" stroke="currentColor" stroke-width="2"/>
            <path d="M3 6h18" stroke="currentColor" stroke-width="2"/>
            <path d="M16 10a4 4 0 11-8 0" stroke="currentColor" stroke-width="2"/>
          </svg>
        </div>
        <div class="stat-content">
          <div class="stat-value">{{ stats.orders || 0 }}</div>
          <div class="stat-label">今日订单</div>
        </div>
        <div class="stat-trend negative">-5%</div>
      </div>

      <div class="stat-card">
        <div class="stat-icon">
          <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <rect x="1" y="4" width="22" height="16" rx="2" ry="2" stroke="currentColor" stroke-width="2"/>
            <path d="M1 10h22" stroke="currentColor" stroke-width="2"/>
          </svg>
        </div>
        <div class="stat-content">
          <div class="stat-value">¥{{ stats.revenue || 0 }}</div>
          <div class="stat-label">今日收入</div>
        </div>
        <div class="stat-trend positive">+15%</div>
      </div>
    </div>

    <!-- 图表区域 -->
    <div class="charts-grid">
      <div class="chart-card">
        <div class="card-header">
          <h3>用户增长趋势</h3>
          <div class="chart-actions">
            <button class="btn btn-sm btn-outline">7天</button>
            <button class="btn btn-sm btn-outline active">30天</button>
            <button class="btn btn-sm btn-outline">90天</button>
          </div>
        </div>
        <div class="chart-placeholder">
          <div class="placeholder-content">
            <svg width="48" height="48" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M21 16V8a2 2 0 00-1-1.73l-7-4a2 2 0 00-2 0l-7 4A2 2 0 003 8v8a2 2 0 001 1.73l7 4a2 2 0 002 0l7-4A2 2 0 0021 16z" stroke="currentColor" stroke-width="2"/>
              <path d="M7.5 4.21l4.5 2.6 4.5-2.6" stroke="currentColor" stroke-width="2"/>
              <path d="M7.5 19.79V14.6L3 12" stroke="currentColor" stroke-width="2"/>
              <path d="M21 12l-4.5 2.6v5.19" stroke="currentColor" stroke-width="2"/>
              <path d="M3.27 6.96L12 12.01l8.73-5.05" stroke="currentColor" stroke-width="2"/>
              <path d="M12 22.08V12" stroke="currentColor" stroke-width="2"/>
            </svg>
            <p>图表组件待集成</p>
            <pre style="text-align:left">{{ JSON.stringify(trends, null, 2) }}</pre>
            <div class="chart-data" style="margin-top:12px">
              <div style="display:grid;grid-template-columns:1fr 1fr;gap:16px">
                <div>
                  <h4 style="margin-bottom:8px">用户趋势</h4>
                  <table class="table" v-if="(trends.users||[]).length">
                    <thead><tr><th>月份</th><th>用户数</th></tr></thead>
                    <tbody>
                      <tr v-for="(row,idx) in trends.users" :key="'u'+idx"><td>{{ row.ym || row.month || row.date }}</td><td>{{ row.users || row.cnt || 0 }}</td></tr>
                    </tbody>
                  </table>
                  <div v-else>暂无数据</div>
                </div>
                <div>
                  <h4 style="margin-bottom:8px">订单趋势</h4>
                  <table class="table" v-if="(trends.orders||[]).length">
                    <thead><tr><th>月份</th><th>订单数</th></tr></thead>
                    <tbody>
                      <tr v-for="(row,idx) in trends.orders" :key="'o'+idx"><td>{{ row.ym || row.month || row.date }}</td><td>{{ row.orders || row.cnt || 0 }}</td></tr>
                    </tbody>
                  </table>
                  <div v-else>暂无数据</div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="chart-card">
        <div class="card-header">
          <h3>收入分布</h3>
          <div class="chart-actions">
            <button class="btn btn-sm btn-outline">月</button>
            <button class="btn btn-sm btn-outline active">季度</button>
            <button class="btn btn-sm btn-outline">年</button>
          </div>
        </div>
        <div class="chart-placeholder">
          <div class="placeholder-content">
            <svg width="48" height="48" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z" stroke="currentColor" stroke-width="2"/>
            </svg>
            <p>饼图组件待集成</p>
            <pre style="text-align:left">{{ JSON.stringify(paymentsStats, null, 2) }}</pre>
            <div class="chart-data" style="margin-top:12px">
              <div style="display:grid;grid-template-columns:1fr 1fr;gap:16px">
                <div>
                  <h4 style="margin-bottom:8px">按状态</h4>
                  <table class="table" v-if="(paymentsStats.by_status||[]).length">
                    <thead><tr><th>状态</th><th>数量</th></tr></thead>
                    <tbody>
                      <tr v-for="(row,idx) in paymentsStats.by_status" :key="'ps'+idx"><td>{{ row.status }}</td><td>{{ row.cnt }}</td></tr>
                    </tbody>
                  </table>
                  <div v-else>暂无数据</div>
                </div>
                <div>
                  <h4 style="margin-bottom:8px">按渠道</h4>
                  <table class="table" v-if="(paymentsStats.by_provider||[]).length">
                    <thead><tr><th>渠道</th><th>数量</th></tr></thead>
                    <tbody>
                      <tr v-for="(row,idx) in paymentsStats.by_provider" :key="'pp'+idx"><td>{{ row.provider }}</td><td>{{ row.cnt }}</td></tr>
                    </tbody>
                  </table>
                  <div v-else>暂无数据</div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 系统状态 -->
    <div class="system-status">
      <div class="card">
        <div class="card-header">
          <h3>系统健康状态</h3>
          <button class="btn btn-sm btn-primary" @click="loadHealth" :disabled="loading">
            <span v-if="loading" class="loading-spinner"></span>
            <span v-else>刷新</span>
          </button>
        </div>
        <div class="status-grid">
          <div class="status-item" v-for="item in healthStatus" :key="item.name">
            <div class="status-indicator" :class="item.status"></div>
            <div class="status-info">
              <div class="status-name">{{ item.name }}</div>
              <div class="status-value">{{ item.value }}</div>
            </div>
          </div>
        </div>
      </div>

      <div class="card">
        <div class="card-header">
          <h3>最近活动</h3>
          <button class="btn btn-sm btn-outline">查看全部</button>
        </div>
        <div class="activity-list">
          <div class="activity-item" v-for="activity in recentActivities" :key="activity.id">
            <div class="activity-icon" :class="activity.type">
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z" stroke="currentColor" stroke-width="2"/>
              </svg>
            </div>
            <div class="activity-content">
              <div class="activity-title">{{ activity.title }}</div>
              <div class="activity-time">{{ activity.time }}</div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 插件状态 -->
    <div class="plugins-status" v-if="plugins && plugins.length">
      <div class="card">
        <div class="card-header">
          <h3>插件状态</h3>
          <button class="btn btn-sm btn-outline" @click="showPlugins = !showPlugins">
            {{ showPlugins ? '收起' : '展开' }}
          </button>
        </div>
        <div v-show="showPlugins" class="plugins-table">
          <table class="table">
            <thead>
              <tr>
                <th>插件名称</th>
                <th>状态</th>
                <th>运行状态</th>
                <th>函数数量</th>
                <th>操作</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="plugin in plugins" :key="plugin.name">
                <td>{{ plugin.name }}</td>
                <td>
                  <span class="badge" :class="{ 'badge-success': plugin.enabled, 'badge-outline': !plugin.enabled }">
                    {{ plugin.enabled ? '启用' : '禁用' }}
                  </span>
                </td>
                <td>
                  <span class="badge" :class="{ 'badge-primary': plugin.running, 'badge-outline': !plugin.running }">
                    {{ plugin.running ? '运行中' : '已停止' }}
                  </span>
                </td>
                <td>{{ plugin.functions?.length || plugin.fnCount || 0 }}</td>
                <td>
                  <div class="action-buttons">
                    <button class="btn btn-sm btn-outline" v-if="plugin.enabled && !plugin.running">启动</button>
                    <button class="btn btn-sm btn-outline" v-if="plugin.running">停止</button>
                    <button class="btn btn-sm btn-outline">配置</button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { overviewApi, systemApi, pluginsApi } from '../utils/api.js'

const stats = ref({})
const healthStatus = ref([])
const recentActivities = ref([])
const plugins = ref([])
const trends = ref({ users: [], orders: [] })
const paymentsStats = ref({ by_status: [], by_provider: [] })
const showPlugins = ref(false)
const loading = ref(false)

// 模拟数据 - 实际应该从API获取
const mockStats = {
  users: 1524,
  tenants: 89,
  orders: 234,
  revenue: 15680
}

const mockHealthStatus = [
  { name: 'API服务', status: 'healthy', value: '正常' },
  { name: '数据库', status: 'healthy', value: '正常' },
  { name: '缓存服务', status: 'warning', value: '延迟较高' },
  { name: '文件存储', status: 'healthy', value: '正常' },
  { name: '消息队列', status: 'error', value: '离线' }
]

const mockActivities = [
  { id: 1, type: 'user', title: '新用户注册：user123', time: '2分钟前' },
  { id: 2, type: 'order', title: '订单完成：ORD-20241119001', time: '5分钟前' },
  { id: 3, type: 'system', title: '系统备份完成', time: '1小时前' },
  { id: 4, type: 'payment', title: '支付成功：¥1,280.00', time: '2小时前' }
]

const loadDashboardData = async () => {
  loading.value = true
  try {
    // 调用仪表盘概览API - 这个接口专门调用插件能力
    const dashboardData = await overviewApi.getDashboard()
    if (dashboardData && dashboardData.data) {
      stats.value = {
        users: dashboardData.data.users || 0,
        tenants: dashboardData.data.tenants || 0,
        orders: dashboardData.data.orders || 0,
        revenue: dashboardData.data.revenue || 0
      }
    } else {
      stats.value = mockStats
    }
    
    // 加载插件状态 - 使用实际接口
    try {
      const pluginsData = await pluginsApi.getExternal()
      plugins.value = pluginsData.data || []
    } catch (pluginError) {
      console.warn('加载插件状态失败:', pluginError)
      plugins.value = []
    }
    // 加载概览趋势与支付统计
    try { const t = await overviewApi.getTrends(); trends.value = t?.data || t || { users: [], orders: [] } } catch (e) { console.warn('加载概览趋势失败:', e) }
    try { const ps = await overviewApi.getPaymentsStats(); paymentsStats.value = ps?.data || ps || { by_status: [], by_provider: [] } } catch (e) { console.warn('加载支付统计失败:', e) }
    
    // 加载健康检查数据 - 汇总CPU/内存/存储/插件与主库状态
    try {
      const healthData = await systemApi.health()
      const d = healthData?.data || healthData || {}
      const ok = (d.sql_main === 'ok' || d.sql_main === 1)
      healthStatus.value = [
        { name: '数据库', status: ok ? 'healthy' : 'error', value: ok ? '正常' : '异常' },
        { name: '插件数量', status: 'healthy', value: String(d.plugin_count ?? 0) },
        { name: '禁用插件', status: (d.plugin_disabled_count ?? 0) > 0 ? 'warning' : 'healthy', value: String(d.plugin_disabled_count ?? 0) },
        { name: 'CPU使用率', status: 'healthy', value: d.cpu_percent ?? d.cpu ?? '0%' },
        { name: '内存使用', status: 'healthy', value: d.memory_used_bytes && d.memory_total_bytes ? `${Math.round(d.memory_used_bytes/1024/1024)}MB / ${Math.round(d.memory_total_bytes/1024/1024)}MB` : '未知' },
        { name: '存储可用', status: 'healthy', value: d.storage_free_bytes && d.storage_total_bytes ? `${Math.round(d.storage_free_bytes/1024/1024/1024)}GB / ${Math.round(d.storage_total_bytes/1024/1024/1024)}GB` : '未知' }
      ]
    } catch (healthError) {
      console.warn('加载健康状态失败:', healthError)
      healthStatus.value = mockHealthStatus
    }
    
    // 使用模拟数据作为后备
    recentActivities.value = mockActivities
    
  } catch (error) {
    console.error('加载仪表盘数据失败:', error)
    // 使用模拟数据作为后备
    stats.value = mockStats
    healthStatus.value = mockHealthStatus
    recentActivities.value = mockActivities
    plugins.value = []
  } finally {
    loading.value = false
  }
}

const loadHealth = async () => {
  try {
    const healthData = await systemApi.health()
    // 处理健康检查数据
    console.log('健康状态:', healthData)
  } catch (error) {
    console.error('健康检查失败:', error)
  }
}

onMounted(() => {
  loadDashboardData()
})
</script>

<style scoped>
.dashboard {
  display: flex;
  flex-direction: column;
  gap: var(--space-xl);
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: var(--space-md);
}

.stat-card {
  background: var(--gradient-dark);
  border: 1px solid var(--border-light);
  border-radius: var(--radius-xl);
  padding: var(--space-lg);
  position: relative;
  overflow: hidden;
  transition: all var(--transition-normal);
}

.stat-card:hover {
  border-color: var(--primary-500);
  transform: translateY(-2px);
  box-shadow: var(--shadow-lg);
}

.stat-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 3px;
  background: var(--gradient-primary);
}

.stat-icon {
  width: 48px;
  height: 48px;
  background: rgba(76, 187, 76, 0.1);
  border-radius: var(--radius-lg);
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: var(--space-md);
}

.stat-icon svg {
  color: var(--primary-500);
}

.stat-content {
  margin-bottom: var(--space-sm);
}

.stat-value {
  font-size: 2.5rem;
  font-weight: 700;
  background: var(--gradient-primary);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  line-height: 1;
}

.stat-label {
  font-size: 0.875rem;
  color: var(--text-secondary);
  margin-top: var(--space-xs);
}

.stat-trend {
  position: absolute;
  top: var(--space-lg);
  right: var(--space-lg);
  font-size: 0.875rem;
  font-weight: 600;
  padding: var(--space-xs) var(--space-sm);
  border-radius: var(--radius-full);
}

.stat-trend.positive {
  background: rgba(16, 185, 129, 0.1);
  color: var(--success);
}

.stat-trend.negative {
  background: rgba(239, 68, 68, 0.1);
  color: var(--error);
}

.charts-grid {
  display: grid;
  grid-template-columns: 2fr 1fr;
  gap: var(--space-xl);
}

.chart-card {
  background: var(--bg-card);
  border: 1px solid var(--border-light);
  border-radius: var(--radius-lg);
  overflow: hidden;
}

.chart-card .card-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: var(--space-lg);
  border-bottom: 1px solid var(--border-light);
}

.chart-actions {
  display: flex;
  gap: var(--space-xs);
}

.chart-actions .btn.active {
  background: var(--primary-500);
  color: white;
}

.chart-placeholder {
  height: 300px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--bg-surface);
}

.placeholder-content {
  text-align: center;
  color: var(--text-muted);
}

.placeholder-content svg {
  margin-bottom: var(--space-md);
  opacity: 0.5;
}

.system-status {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: var(--space-xl);
}

.status-grid {
  display: grid;
  gap: var(--space-md);
  padding: var(--space-lg);
}

.status-item {
  display: flex;
  align-items: center;
  gap: var(--space-md);
  padding: var(--space-md);
  background: var(--bg-surface);
  border-radius: var(--radius-md);
}

.status-indicator {
  width: 12px;
  height: 12px;
  border-radius: 50%;
}

.status-indicator.healthy {
  background: var(--success);
}

.status-indicator.warning {
  background: var(--warning);
}

.status-indicator.error {
  background: var(--error);
}

.status-name {
  font-weight: 600;
  margin-bottom: 2px;
}

.status-value {
  font-size: 0.875rem;
  color: var(--text-secondary);
}

.activity-list {
  padding: var(--space-lg);
}

.activity-item {
  display: flex;
  align-items: center;
  gap: var(--space-md);
  padding: var(--space-md);
  border-radius: var(--radius-md);
  transition: background var(--transition-fast);
}

.activity-item:hover {
  background: var(--bg-surface);
}

.activity-icon {
  width: 32px;
  height: 32px;
  border-radius: var(--radius-full);
  display: flex;
  align-items: center;
  justify-content: center;
}

.activity-icon.user {
  background: rgba(76, 187, 76, 0.1);
}

.activity-icon.order {
  background: rgba(251, 146, 60, 0.1);
}

.activity-icon.payment {
  background: rgba(16, 185, 129, 0.1);
}

.activity-icon.system {
  background: rgba(59, 130, 246, 0.1);
}

.activity-icon svg {
  width: 16px;
  height: 16px;
}

.activity-content {
  flex: 1;
}

.activity-title {
  font-weight: 500;
  margin-bottom: 2px;
}

.activity-time {
  font-size: 0.75rem;
  color: var(--text-muted);
}

.plugins-table {
  padding: var(--space-lg);
}

.action-buttons {
  display: flex;
  gap: var(--space-xs);
}

@media (max-width: 1024px) {
  .charts-grid {
    grid-template-columns: 1fr;
  }
  
  .system-status {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 768px) {
  .stats-grid {
    grid-template-columns: 1fr;
  }
}
</style>