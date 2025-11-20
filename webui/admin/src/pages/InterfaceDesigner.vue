<template>
  <div class="page-container">
    <div class="page-header">
      <h1 class="page-title">接口设计器</h1>
      <div class="page-description">可视化配置API接口的DSL步骤</div>
    </div>

    <div class="designer-layout">
      <!-- 左侧面板：接口配置 -->
      <div class="designer-panel">
        <div class="panel-header">
          <h3 class="panel-title">📝 接口配置</h3>
        </div>
        <div class="panel-body">
          <div class="form-group">
            <label class="form-label">接口名称</label>
            <input v-model="interfaceConfig.name" class="tech-input" placeholder="用户信息查询" />
          </div>
          <div class="form-group">
            <label class="form-label">接口路径</label>
            <input v-model="interfaceConfig.path" class="tech-input" placeholder="/api/user/{id}" />
          </div>
          <div class="form-group">
            <label class="form-label">HTTP方法</label>
            <select v-model="interfaceConfig.method" class="tech-input">
              <option value="GET">GET</option>
              <option value="POST">POST</option>
              <option value="PUT">PUT</option>
              <option value="DELETE">DELETE</option>
              <option value="PATCH">PATCH</option>
            </select>
          </div>
          <div class="form-group">
            <label class="form-label">接口描述</label>
            <textarea v-model="interfaceConfig.description" class="tech-textarea" 
                      placeholder="查询用户详细信息" rows="3"></textarea>
          </div>
          <div class="form-group">
            <label class="form-label">认证方式</label>
            <select v-model="interfaceConfig.auth" class="tech-input">
              <option value="none">无需认证</option>
              <option value="jwt">JWT认证</option>
              <option value="apiKey">API Key</option>
              <option value="hmac">HMAC签名</option>
            </select>
          </div>
          <div class="form-group">
            <label class="form-label">权限要求</label>
            <input v-model="interfaceConfig.permission" class="tech-input" placeholder="user.read" />
          </div>
        </div>
      </div>

      <!-- 中间面板：DSL步骤设计器 -->
      <div class="designer-main">
        <div class="steps-toolbar">
          <div class="toolbar-left">
            <h3 class="panel-title">🔧 DSL步骤设计</h3>
          </div>
          <div class="toolbar-right">
            <button class="tech-btn secondary" @click="addStep">
              <span class="btn-icon">➕</span>
              添加步骤
            </button>
            <button class="tech-btn primary" @click="saveInterface">
              <span class="btn-icon">💾</span>
              保存接口
            </button>
          </div>
        </div>

        <div class="steps-container">
          <div v-for="(step, index) in dslSteps" :key="index" class="step-item">
            <div class="step-header">
              <div class="step-number">{{ index + 1 }}</div>
              <div class="step-type">{{ getStepTypeName(step.type) }}</div>
              <div class="step-actions">
                <button class="tech-btn ghost small" @click="editStep(index)">编辑</button>
                <button class="tech-btn danger small" @click="removeStep(index)">删除</button>
                <button class="tech-btn ghost small" @click="moveStep(index, -1)" 
                        :disabled="index === 0">上移</button>
                <button class="tech-btn ghost small" @click="moveStep(index, 1)" 
                        :disabled="index === dslSteps.length - 1">下移</button>
              </div>
            </div>
            <div class="step-content">
              <pre class="step-preview">{{ JSON.stringify(step.config, null, 2) }}</pre>
            </div>
          </div>
          
          <div v-if="dslSteps.length === 0" class="empty-steps">
            <div class="empty-icon">📋</div>
            <div class="empty-text">暂无步骤，点击"添加步骤"开始设计</div>
          </div>
        </div>
      </div>

      <!-- 右侧面板：步骤模板库 -->
      <div class="designer-panel">
        <div class="panel-header">
          <h3 class="panel-title">📚 步骤模板库</h3>
        </div>
        <div class="panel-body">
          <div class="step-templates">
            <div v-for="template in stepTemplates" :key="template.type" 
                 class="template-item" @click="addStepFromTemplate(template)">
              <div class="template-icon">{{ template.icon }}</div>
              <div class="template-info">
                <div class="template-name">{{ template.name }}</div>
                <div class="template-desc">{{ template.description }}</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 步骤编辑模态框 -->
    <div v-if="editingStep !== null" class="modal-overlay">
      <div class="modal-content">
        <div class="modal-header">
          <h3 class="modal-title">编辑步骤</h3>
          <button class="modal-close" @click="cancelEdit">×</button>
        </div>
        <div class="modal-body">
          <StepEditor 
            :step="editingStep" 
            @save="saveEditedStep" 
            @cancel="cancelEdit" 
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import StepEditor from '../components/StepEditor.vue'

const interfaceConfig = ref({
  name: '',
  path: '',
  method: 'GET',
  description: '',
  auth: 'none',
  permission: ''
})

const dslSteps = ref([])
const editingStep = ref(null)
const editingIndex = ref(-1)

const stepTemplates = [
  { type: 'validate', icon: '✅', name: '参数校验', description: '验证请求参数的有效性' },
  { type: 'db_query', icon: '🗄️', name: '数据库查询', description: '执行SQL查询操作' },
  { type: 'transform', icon: '🔄', name: '数据转换', description: '转换数据格式和结构' },
  { type: 'response', icon: '📤', name: '响应构建', description: '构建API响应数据' },
  { type: 'cache', icon: '⚡', name: '缓存操作', description: '读写缓存数据' },
  { type: 'hook', icon: '🔗', name: '钩子调用', description: '执行前后置钩子' },
  { type: 'condition', icon: '🔀', name: '条件分支', description: '根据条件执行不同逻辑' },
  { type: 'loop', icon: '🔄', name: '循环处理', description: '对数组数据进行循环处理' }
]

function getStepTypeName(type) {
  const names = {
    validate: '参数校验',
    db_query: '数据库查询',
    transform: '数据转换',
    response: '响应构建',
    cache: '缓存操作',
    hook: '钩子调用',
    condition: '条件分支',
    loop: '循环处理'
  }
  return names[type] || type
}

function addStep() {
  const newStep = {
    type: 'validate',
    config: {}
  }
  dslSteps.value.push(newStep)
  editStep(dslSteps.value.length - 1)
}

function addStepFromTemplate(template) {
  const newStep = {
    type: template.type,
    config: {}
  }
  dslSteps.value.push(newStep)
  editStep(dslSteps.value.length - 1)
}

function editStep(index) {
  editingIndex.value = index
  editingStep.value = JSON.parse(JSON.stringify(dslSteps.value[index]))
}

function saveEditedStep(updatedStep) {
  if (editingIndex.value >= 0) {
    dslSteps.value[editingIndex.value] = updatedStep
  }
  cancelEdit()
}

function cancelEdit() {
  editingStep.value = null
  editingIndex.value = -1
}

function removeStep(index) {
  dslSteps.value.splice(index, 1)
}

function moveStep(index, direction) {
  const newIndex = index + direction
  if (newIndex >= 0 && newIndex < dslSteps.value.length) {
    const temp = dslSteps.value[newIndex]
    dslSteps.value[newIndex] = dslSteps.value[index]
    dslSteps.value[index] = temp
  }
}

async function saveInterface() {
  if (!interfaceConfig.value.name || !interfaceConfig.value.path) {
    alert('请填写接口名称和路径')
    return
  }

  try {
    const interfaceData = {
      ...interfaceConfig.value,
      steps: dslSteps.value
    }
    
    // 这里调用保存接口的API
    console.log('保存接口:', interfaceData)
    alert('接口保存成功')
  } catch (error) {
    console.error('保存接口失败:', error)
    alert('保存失败')
  }
}
</script>

<style scoped>
.designer-layout {
  display: grid;
  grid-template-columns: 300px 1fr 300px;
  gap: 20px;
  height: calc(100vh - 200px);
}

.designer-panel {
  background: var(--surface-primary);
  border-radius: 12px;
  border: 1px solid var(--border-light);
  display: flex;
  flex-direction: column;
}

.panel-header {
  padding: 16px 20px;
  border-bottom: 1px solid var(--border-light);
}

.panel-title {
  margin: 0;
  font-size: 16px;
  font-weight: 600;
}

.panel-body {
  padding: 20px;
  flex: 1;
  overflow-y: auto;
}

.designer-main {
  background: var(--surface-primary);
  border-radius: 12px;
  border: 1px solid var(--border-light);
  display: flex;
  flex-direction: column;
}

.steps-toolbar {
  padding: 16px 20px;
  border-bottom: 1px solid var(--border-light);
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.steps-container {
  padding: 20px;
  flex: 1;
  overflow-y: auto;
}

.step-item {
  background: var(--surface-secondary);
  border-radius: 8px;
  border: 1px solid var(--border-light);
  margin-bottom: 12px;
  overflow: hidden;
}

.step-header {
  padding: 12px 16px;
  background: var(--accent-primary);
  color: white;
  display: flex;
  align-items: center;
  gap: 12px;
}

.step-number {
  background: rgba(255,255,255,0.2);
  border-radius: 50%;
  width: 24px;
  height: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 12px;
  font-weight: 600;
}

.step-type {
  flex: 1;
  font-weight: 600;
}

.step-actions {
  display: flex;
  gap: 8px;
}

.step-content {
  padding: 16px;
}

.step-preview {
  background: var(--surface-primary);
  border-radius: 4px;
  padding: 12px;
  font-size: 12px;
  line-height: 1.4;
  overflow-x: auto;
  max-height: 100px;
}

.empty-steps {
  text-align: center;
  padding: 60px 20px;
  color: var(--text-secondary);
}

.step-templates {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.template-item {
  background: var(--surface-secondary);
  border-radius: 8px;
  padding: 12px;
  cursor: pointer;
  transition: all 0.2s;
  border: 1px solid transparent;
}

.template-item:hover {
  border-color: var(--accent-primary);
  transform: translateY(-2px);
}

.template-item {
  display: flex;
  align-items: center;
  gap: 12px;
}

.template-icon {
  font-size: 20px;
}

.template-info {
  flex: 1;
}

.template-name {
  font-weight: 600;
  margin-bottom: 4px;
}

.template-desc {
  font-size: 12px;
  color: var(--text-secondary);
}

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0,0,0,0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal-content {
  background: var(--surface-primary);
  border-radius: 12px;
  width: 90%;
  max-width: 600px;
  max-height: 80vh;
  overflow: hidden;
}

.modal-header {
  padding: 20px;
  border-bottom: 1px solid var(--border-light);
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.modal-close {
  background: none;
  border: none;
  font-size: 24px;
  cursor: pointer;
  color: var(--text-secondary);
}

.modal-body {
  padding: 20px;
  max-height: calc(80vh - 80px);
  overflow-y: auto;
}

@media (max-width: 1200px) {
  .designer-layout {
    grid-template-columns: 1fr;
    grid-template-rows: auto auto auto;
  }
  
  .designer-panel {
    max-height: 300px;
  }
}
</style>