<template>
  <div class="page-container">
    <div class="page-header">
      <h1 class="page-title">Hooks 检视</h1>
      <div class="page-description">检视接口的钩子执行链和匹配规则</div>
    </div>

    <div class="card tech-card">
      <div class="card-header">
        <h3 class="card-title">🔍 接口钩子检视</h3>
      </div>
      <div class="card-body">
        <div class="input-group">
          <div class="input-row">
            <select v-model="method" class="tech-input">
              <option value="GET">GET</option>
              <option value="POST">POST</option>
              <option value="PUT">PUT</option>
              <option value="DELETE">DELETE</option>
              <option value="PATCH">PATCH</option>
            </select>
            <input v-model="path" placeholder="/api/user/{id}" class="tech-input flex-1" />
            <button class="tech-btn primary" @click="inspect">
              <span class="btn-icon">🔍</span>
              检视
            </button>
          </div>
        </div>

        <div class="results-grid" v-if="chain.length > 0 || Object.keys(match).length > 0">
          <div class="tech-card">
            <div class="card-header">
              <h4 class="card-title">📋 策略执行链</h4>
            </div>
            <div class="card-body">
              <div class="chain-list">
                <div v-for="(item, index) in chain" :key="index" class="chain-item">
                  <div class="chain-index">{{ index + 1 }}</div>
                  <div class="chain-content">{{ item }}</div>
                </div>
                <div v-if="chain.length === 0" class="empty-state">
                  暂无策略链信息
                </div>
              </div>
            </div>
          </div>

          <div class="tech-card">
            <div class="card-header">
              <h4 class="card-title">🎯 匹配信息</h4>
            </div>
            <div class="card-body">
              <pre class="code-block">{{ json(match) }}</pre>
            </div>
          </div>
        </div>

        <div v-else class="empty-state">
          <div class="empty-icon">🔍</div>
          <div class="empty-text">输入接口路径和方法进行检视</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { hooksApi } from '../utils/api.js'

const method = ref('GET')
const path = ref('/api/admin/overview')
const chain = ref([])
const match = ref({})

const json = (v) => JSON.stringify(v, null, 2)

async function inspect() {
  if (!path.value) {
    alert('请输入接口路径')
    return
  }
  
  try {
    const r = await hooksApi.inspect({ 
      method: method.value, 
      path: path.value 
    })
    const d = r?.data ?? r ?? {}
    chain.value = d.chain || []
    match.value = d.match || {}
  } catch (error) {
    console.error('检视失败:', error)
    alert('检视失败，请检查接口路径是否正确')
  }
}
</script>

<style scoped>
.input-row {
  display: flex;
  gap: 12px;
  align-items: center;
  flex-wrap: wrap;
}

.results-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 20px;
  margin-top: 20px;
}

.chain-list {
  max-height: 300px;
  overflow-y: auto;
}

.chain-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 8px 12px;
  margin-bottom: 8px;
  background: var(--surface-secondary);
  border-radius: 8px;
  border-left: 3px solid var(--accent-primary);
}

.chain-index {
  background: var(--accent-primary);
  color: white;
  border-radius: 50%;
  width: 24px;
  height: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 12px;
  font-weight: 600;
}

.chain-content {
  flex: 1;
  font-family: 'Courier New', monospace;
  font-size: 13px;
}

.code-block {
  background: var(--surface-secondary);
  border-radius: 8px;
  padding: 16px;
  font-size: 13px;
  line-height: 1.4;
  overflow-x: auto;
  max-height: 300px;
  border: 1px solid var(--border-light);
}

@media (max-width: 768px) {
  .results-grid {
    grid-template-columns: 1fr;
  }
  
  .input-row {
    flex-direction: column;
    align-items: stretch;
  }
  
  .tech-input {
    width: 100%;
  }
}
</style>