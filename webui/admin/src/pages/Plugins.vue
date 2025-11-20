<template>
  <div class="page-container">
    <div class="page-header">
      <h1 class="page-title">插件管理</h1>
      <div class="page-description">管理外部插件和内置插件的状态与配置</div>
    </div>

    <div class="grid tech-grid">
      <!-- 外置插件状态 -->
      <div class="tech-card">
        <div class="card-header">
          <h3 class="card-title">🔌 外置插件</h3>
          <div class="card-actions">
            <button class="tech-btn secondary" @click="status">
              <span class="btn-icon">🔄</span>
              刷新
            </button>
          </div>
        </div>
        <div class="card-body">
          <div class="table-container">
            <table class="tech-table" v-if="list.length > 0">
              <thead>
                <tr>
                  <th>插件名称</th>
                  <th>端点数量</th>
                  <th>状态</th>
                  <th>操作</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="it in list" :key="it.name">
                  <td>
                    <div class="plugin-name">
                      <span class="plugin-icon">🔧</span>
                      {{ it.name }}
                    </div>
                  </td>
                  <td>
                    <span class="badge info">{{ (it.endpoints || []).length }}</span>
                  </td>
                  <td>
                    <span class="status-badge" :class="{ active: it.enabled }">
                      {{ it.enabled ? '已启用' : '已禁用' }}
                    </span>
                  </td>
                  <td>
                    <div class="action-buttons">
                      <button 
                        v-if="it.enabled" 
                        class="tech-btn danger small"
                        @click="togglePlugin(it.name, false)"
                      >
                        禁用
                      </button>
                      <button 
                        v-else 
                        class="tech-btn success small"
                        @click="togglePlugin(it.name, true)"
                      >
                        启用
                      </button>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
            <div v-else class="empty-state">
              <div class="empty-icon">🔌</div>
              <div class="empty-text">暂无外置插件</div>
            </div>
          </div>
        </div>
      </div>

      <!-- 动态插件管理 -->
      <div class="tech-card">
        <div class="card-header">
          <h3 class="card-title">⚡ 动态插件管理</h3>
        </div>
        <div class="card-body">
          <div class="form-grid">
            <div class="form-group">
              <label class="form-label">插件名称</label>
              <input v-model="addForm.name" placeholder="metrics" class="tech-input" />
            </div>
            <div class="form-group">
              <label class="form-label">实例数量</label>
              <input v-model="addForm.instances" type="number" min="1" class="tech-input" />
            </div>
            <div class="form-group">
              <label class="form-label">超时时间</label>
              <input v-model="addForm.timeout" placeholder="300ms" class="tech-input" />
            </div>
            <div class="form-group">
              <label class="form-label">队列大小</label>
              <input v-model="addForm.queueSize" type="number" min="1" class="tech-input" />
            </div>
            <div class="form-group">
              <label class="form-label">Windows路径</label>
              <input v-model="addForm.windows" placeholder="C:\\path\\to\\plugin.exe" class="tech-input" />
            </div>
            <div class="form-group">
              <label class="form-label">Unix路径</label>
              <input v-model="addForm.unix" placeholder="/usr/local/bin/plugin" class="tech-input" />
            </div>
            <div class="form-group full-width">
              <label class="form-label">函数列表（逗号分隔）</label>
              <input v-model="addForm.functions" placeholder="function1, function2, function3" class="tech-input" />
            </div>
          </div>
          
          <div class="action-bar">
            <button class="tech-btn primary" @click="add">
              <span class="btn-icon">➕</span>
              新增插件
            </button>
            <button class="tech-btn danger" @click="remove">
              <span class="btn-icon">🗑️</span>
              移除插件
            </button>
            <div class="plugin-controls">
              <button class="tech-btn warning" @click="stop">停止</button>
              <button class="tech-btn success" @click="start">启动</button>
              <button class="tech-btn info" @click="restart">重启</button>
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

const list = ref([])
const addForm = ref({
  name: 'metrics',
  windows: '',
  unix: '',
  instances: 2,
  timeout: '300ms',
  queueSize: 1024,
  functions: ''
})

onMounted(() => {
  status()
})

async function status() {
  try {
    const res = await api.get('/api/admin/plugins/external')
    list.value = res?.data || res || []
  } catch (error) {
    console.error('获取插件状态失败:', error)
    list.value = []
  }
}

async function togglePlugin(name, enabled) {
  try {
    await api.post('/api/admin/plugins/enable', {
      names: [name],
      enabled: enabled
    })
    await status()
  } catch (error) {
    console.error('切换插件状态失败:', error)
    alert('操作失败')
  }
}

async function add() {
  if (!addForm.value.name) {
    alert('请输入插件名称')
    return
  }
  
  try {
    const fns = addForm.value.functions.split(',')
      .map(s => s.trim())
      .filter(Boolean)
    
    await api.post('/api/admin/plugins/add', {
      name: addForm.value.name,
      windows: addForm.value.windows,
      unix: addForm.value.unix,
      instances: Number(addForm.value.instances || 2),
      timeout: addForm.value.timeout,
      queueSize: Number(addForm.value.queueSize || 1024),
      functions: fns
    })
    
    await status()
    alert('插件添加成功')
  } catch (error) {
    console.error('添加插件失败:', error)
    alert('添加失败')
  }
}

async function remove() {
  if (!addForm.value.name) {
    alert('请输入要移除的插件名称')
    return
  }
  
  try {
    await api.post('/api/admin/plugins/remove', {
      names: [addForm.value.name]
    })
    await status()
    alert('插件移除成功')
  } catch (error) {
    console.error('移除插件失败:', error)
    alert('移除失败')
  }
}

async function stop() {
  if (!addForm.value.name) {
    alert('请输入插件名称')
    return
  }
  
  try {
    await api.post('/api/admin/plugins/stop', {
      names: [addForm.value.name]
    })
    await status()
  } catch (error) {
    console.error('停止插件失败:', error)
  }
}

async function start() {
  if (!addForm.value.name) {
    alert('请输入插件名称')
    return
  }
  
  try {
    await api.post('/api/admin/plugins/start', {
      names: [addForm.value.name]
    })
    await status()
  } catch (error) {
    console.error('启动插件失败:', error)
  }
}

async function restart() {
  if (!addForm.value.name) {
    alert('请输入插件名称')
    return
  }
  
  try {
    await api.post('/api/admin/plugins/restart', {
      names: [addForm.value.name]
    })
    await status()
  } catch (error) {
    console.error('重启插件失败:', error)
  }
}
</script>

<style scoped>
.form-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
  margin-bottom: 20px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.form-group.full-width {
  grid-column: 1 / -1;
}

.form-label {
  font-weight: 600;
  color: var(--text-secondary);
  font-size: 14px;
}

.action-bar {
  display: flex;
  gap: 12px;
  align-items: center;
  flex-wrap: wrap;
}

.plugin-controls {
  display: flex;
  gap: 8px;
  margin-left: auto;
}

.plugin-name {
  display: flex;
  align-items: center;
  gap: 8px;
  font-weight: 600;
}

.plugin-icon {
  font-size: 16px;
}

.status-badge {
  padding: 4px 8px;
  border-radius: 12px;
  font-size: 12px;
  font-weight: 600;
}

.status-badge.active {
  background: var(--success-light);
  color: var(--success-dark);
}

.status-badge:not(.active) {
  background: var(--danger-light);
  color: var(--danger-dark);
}

.action-buttons {
  display: flex;
  gap: 8px;
}

@media (max-width: 768px) {
  .form-grid {
    grid-template-columns: 1fr;
  }
  
  .action-bar {
    flex-direction: column;
    align-items: stretch;
  }
  
  .plugin-controls {
    margin-left: 0;
    justify-content: center;
  }
}
</style>