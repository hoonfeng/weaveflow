<template>
  <div class="page-container">
    <div class="page-header">
      <h1 class="page-title">项目配置</h1>
      <div class="page-description">管理系统配置参数和运行环境</div>
    </div>

    <div class="grid tech-grid">
      <!-- 基本配置 -->
      <div class="tech-card">
        <div class="card-header">
          <h3 class="card-title">⚙️ 基本配置</h3>
          <div class="card-actions">
            <button class="tech-btn secondary" @click="loadConfig">
              <span class="btn-icon">🔄</span>
              刷新
            </button>
            <button class="tech-btn primary" @click="saveConfig">
              <span class="btn-icon">💾</span>
              保存
            </button>
          </div>
        </div>
        <div class="card-body">
          <div class="config-form">
            <div class="form-group">
              <label class="form-label">项目名称</label>
              <input v-model="config.name" class="tech-input" placeholder="接口配置化系统" />
            </div>
            <div class="form-group">
              <label class="form-label">监听端口</label>
              <input v-model="config.port" type="number" class="tech-input" placeholder="8080" />
            </div>
            <div class="form-group">
              <label class="form-label">环境模式</label>
              <select v-model="config.env" class="tech-input">
                <option value="development">开发环境</option>
                <option value="testing">测试环境</option>
                <option value="production">生产环境</option>
              </select>
            </div>
            <div class="form-group">
              <label class="form-label">日志级别</label>
              <select v-model="config.logLevel" class="tech-input">
                <option value="debug">DEBUG</option>
                <option value="info">INFO</option>
                <option value="warn">WARN</option>
                <option value="error">ERROR</option>
              </select>
            </div>
            <div class="form-group">
              <label class="form-label">API前缀</label>
              <input v-model="config.apiPrefix" class="tech-input" placeholder="/api" />
            </div>
            <div class="form-group">
              <label class="form-label">JWT密钥</label>
              <div class="input-with-action">
                <input v-model="config.jwtSecret" type="password" class="tech-input" />
                <button class="tech-btn ghost small" @click="toggleJwtVisibility">
                  {{ showJwt ? '隐藏' : '显示' }}
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- 数据库配置 -->
      <div class="tech-card">
        <div class="card-header">
          <h3 class="card-title">🗄️ 数据库配置</h3>
        </div>
        <div class="card-body">
          <div class="config-form">
            <div class="form-group">
              <label class="form-label">数据库类型</label>
              <select v-model="config.dbType" class="tech-input">
                <option value="sqlite">SQLite</option>
                <option value="mysql">MySQL</option>
                <option value="postgres">PostgreSQL</option>
              </select>
            </div>
            <div class="form-group">
              <label class="form-label">连接字符串</label>
              <input v-model="config.dbConn" class="tech-input" 
                     placeholder="host=localhost port=5432 user=postgres password=123456 dbname=ifaceconf" />
            </div>
            <div class="form-group">
              <label class="form-label">最大连接数</label>
              <input v-model="config.maxConnections" type="number" class="tech-input" placeholder="10" />
            </div>
            <div class="form-group">
              <label class="form-label">连接超时(秒)</label>
              <input v-model="config.connTimeout" type="number" class="tech-input" placeholder="30" />
            </div>
          </div>
        </div>
      </div>

      <!-- 缓存配置 -->
      <div class="tech-card">
        <div class="card-header">
          <h3 class="card-title">⚡ 缓存配置</h3>
        </div>
        <div class="card-body">
          <div class="config-form">
            <div class="form-group">
              <label class="form-label">缓存类型</label>
              <select v-model="config.cacheType" class="tech-input">
                <option value="memory">内存缓存</option>
                <option value="redis">Redis</option>
                <option value="none">无缓存</option>
              </select>
            </div>
            <div class="form-group">
              <label class="form-label">Redis地址</label>
              <input v-model="config.redisAddr" class="tech-input" placeholder="localhost:6379" />
            </div>
            <div class="form-group">
              <label class="form-label">Redis密码</label>
              <input v-model="config.redisPassword" type="password" class="tech-input" />
            </div>
            <div class="form-group">
              <label class="form-label">默认TTL(秒)</label>
              <input v-model="config.defaultTTL" type="number" class="tech-input" placeholder="3600" />
            </div>
          </div>
        </div>
      </div>

      <!-- 安全配置 -->
      <div class="tech-card">
        <div class="card-header">
          <h3 class="card-title">🔒 安全配置</h3>
        </div>
        <div class="card-body">
          <div class="config-form">
            <div class="form-group">
              <label class="form-label">CORS域名</label>
              <input v-model="config.corsOrigins" class="tech-input" placeholder="http://localhost:3000" />
            </div>
            <div class="form-group">
              <label class="form-label">API限流(请求/秒)</label>
              <input v-model="config.rateLimit" type="number" class="tech-input" placeholder="100" />
            </div>
            <div class="form-group">
              <label class="form-label">会话超时(分钟)</label>
              <input v-model="config.sessionTimeout" type="number" class="tech-input" placeholder="30" />
            </div>
            <div class="form-group">
              <label class="form-label">密码强度要求</label>
              <select v-model="config.passwordStrength" class="tech-input">
                <option value="low">低强度</option>
                <option value="medium">中强度</option>
                <option value="high">高强度</option>
              </select>
            </div>
          </div>
        </div>
      </div>

      <!-- 系统状态 -->
      <div class="tech-card">
        <div class="card-header">
          <h3 class="card-title">📊 系统状态</h3>
        </div>
        <div class="card-body">
          <div class="status-grid">
            <div class="status-item">
              <div class="status-label">启动时间</div>
              <div class="status-value">{{ systemStatus.startTime || '未知' }}</div>
            </div>
            <div class="status-item">
              <div class="status-label">运行时长</div>
              <div class="status-value">{{ systemStatus.uptime || '未知' }}</div>
            </div>
            <div class="status-item">
              <div class="status-label">内存使用</div>
              <div class="status-value">{{ systemStatus.memoryUsage || '未知' }}</div>
            </div>
            <div class="status-item">
              <div class="status-label">接口数量</div>
              <div class="status-value">{{ systemStatus.interfaceCount || '未知' }}</div>
            </div>
            <div class="status-item">
              <div class="status-label">活跃会话</div>
              <div class="status-value">{{ systemStatus.activeSessions || '未知' }}</div>
            </div>
            <div class="status-item">
              <div class="status-label">请求总数</div>
              <div class="status-value">{{ systemStatus.totalRequests || '未知' }}</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import api from '../utils/api.js'

const config = ref({
  name: '',
  port: 8080,
  env: 'development',
  logLevel: 'info',
  apiPrefix: '/api',
  jwtSecret: '',
  dbType: 'sqlite',
  dbConn: '',
  maxConnections: 10,
  connTimeout: 30,
  cacheType: 'memory',
  redisAddr: '',
  redisPassword: '',
  defaultTTL: 3600,
  corsOrigins: '',
  rateLimit: 100,
  sessionTimeout: 30,
  passwordStrength: 'medium'
})

const systemStatus = ref({})
const showJwt = ref(false)

onMounted(() => {
  loadConfig()
  loadSystemStatus()
})

async function loadConfig() {
  try {
    const res = await api.get('/api/admin/config')
    if (res?.data) {
      config.value = { ...config.value, ...res.data }
    }
  } catch (error) {
    console.error('加载配置失败:', error)
  }
}

async function loadSystemStatus() {
  try {
    const res = await api.get('/api/admin/status')
    systemStatus.value = res?.data || {}
  } catch (error) {
    console.error('加载系统状态失败:', error)
  }
}

async function saveConfig() {
  try {
    await api.post('/api/admin/config', config.value)
    alert('配置保存成功')
  } catch (error) {
    console.error('保存配置失败:', error)
    alert('保存失败')
  }
}

function toggleJwtVisibility() {
  showJwt.value = !showJwt.value
  const input = document.querySelector('input[type="password"]')
  if (input) {
    input.type = showJwt.value ? 'text' : 'password'
  }
}
</script>

<style scoped>
.config-form {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.form-label {
  font-weight: 600;
  color: var(--text-secondary);
  font-size: 14px;
}

.input-with-action {
  display: flex;
  gap: 8px;
  align-items: center;
}

.input-with-action .tech-input {
  flex: 1;
}

.status-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
  gap: 16px;
}

.status-item {
  background: var(--surface-secondary);
  padding: 12px;
  border-radius: 8px;
  text-align: center;
}

.status-label {
  font-size: 12px;
  color: var(--text-secondary);
  margin-bottom: 4px;
}

.status-value {
  font-weight: 600;
  font-size: 14px;
  color: var(--accent-primary);
}

@media (max-width: 768px) {
  .status-grid {
    grid-template-columns: 1fr 1fr;
  }
  
  .input-with-action {
    flex-direction: column;
  }
}
</style>