<template>
  <div class="step-editor">
    <div class="editor-header">
      <h4 class="editor-title">编辑步骤配置</h4>
    </div>
    
    <div class="editor-body">
      <div class="form-group">
        <label class="form-label">步骤类型</label>
        <select v-model="step.type" class="tech-input" @change="resetConfig">
          <option value="validate">参数校验</option>
          <option value="db_query">数据库查询</option>
          <option value="transform">数据转换</option>
          <option value="response">响应构建</option>
          <option value="cache">缓存操作</option>
          <option value="hook">钩子调用</option>
          <option value="condition">条件分支</option>
          <option value="loop">循环处理</option>
        </select>
      </div>

      <!-- 参数校验配置 -->
      <div v-if="step.type === 'validate'" class="config-section">
        <h5 class="section-title">参数校验配置</h5>
        <div class="form-group">
          <label class="form-label">校验规则</label>
          <textarea v-model="step.config.rules" class="tech-textarea" rows="4" 
                    placeholder="例如: {
  \"id\": \"required,numeric\",
  \"name\": \"required,min:2,max:50\"
}"></textarea>
        </div>
        <div class="form-group">
          <label class="form-label">错误消息</label>
          <textarea v-model="step.config.errorMessage" class="tech-textarea" rows="2" 
                    placeholder="自定义错误消息"></textarea>
        </div>
      </div>

      <!-- 数据库查询配置 -->
      <div v-if="step.type === 'db_query'" class="config-section">
        <h5 class="section-title">数据库查询配置</h5>
        <div class="form-group">
          <label class="form-label">SQL模板</label>
          <textarea v-model="step.config.sql" class="tech-textarea" rows="4" 
                    placeholder="SELECT * FROM users WHERE id = {{.id}}"></textarea>
        </div>
        <div class="form-group">
          <label class="form-label">数据源</label>
          <input v-model="step.config.dataSource" class="tech-input" placeholder="default" />
        </div>
        <div class="form-group">
          <label class="form-label">超时时间(秒)</label>
          <input v-model="step.config.timeout" type="number" class="tech-input" placeholder="30" />
        </div>
      </div>

      <!-- 数据转换配置 -->
      <div v-if="step.type === 'transform'" class="config-section">
        <h5 class="section-title">数据转换配置</h5>
        <div class="form-group">
          <label class="form-label">转换模板</label>
          <textarea v-model="step.config.template" class="tech-textarea" rows="4" 
                    placeholder="{
  \"userInfo\": {
    \"name\": \"{{.name}}\",
    \"email\": \"{{.email}}\"
  }
}"></textarea>
        </div>
        <div class="form-group">
          <label class="form-label">输入字段映射</label>
          <textarea v-model="step.config.fieldMapping" class="tech-textarea" rows="3" 
                    placeholder="name: fullName, email: userEmail"></textarea>
        </div>
      </div>

      <!-- 响应构建配置 -->
      <div v-if="step.type === 'response'" class="config-section">
        <h5 class="section-title">响应构建配置</h5>
        <div class="form-group">
          <label class="form-label">HTTP状态码</label>
          <input v-model="step.config.statusCode" type="number" class="tech-input" placeholder="200" />
        </div>
        <div class="form-group">
          <label class="form-label">响应头</label>
          <textarea v-model="step.config.headers" class="tech-textarea" rows="3" 
                    placeholder="Content-Type: application/json
Cache-Control: no-cache"></textarea>
        </div>
        <div class="form-group">
          <label class="form-label">响应体模板</label>
          <textarea v-model="step.config.body" class="tech-textarea" rows="4" 
                    placeholder="{
  \"success\": true,
  \"data\": {{.result}}
}"></textarea>
        </div>
      </div>

      <!-- 条件分支配置 -->
      <div v-if="step.type === 'condition'" class="config-section">
        <h5 class="section-title">条件分支配置</h5>
        <div class="form-group">
          <label class="form-label">条件表达式</label>
          <textarea v-model="step.config.condition" class="tech-textarea" rows="3" 
                    placeholder="{{.user.role}} == 'admin'"></textarea>
        </div>
        <div class="form-group">
          <label class="form-label">真分支步骤</label>
          <textarea v-model="step.config.trueSteps" class="tech-textarea" rows="3" 
                    placeholder="步骤配置数组"></textarea>
        </div>
        <div class="form-group">
          <label class="form-label">假分支步骤</label>
          <textarea v-model="step.config.falseSteps" class="tech-textarea" rows="3" 
                    placeholder="步骤配置数组"></textarea>
        </div>
      </div>

      <!-- 其他步骤类型的通用配置 -->
      <div v-else class="config-section">
        <h5 class="section-title">步骤配置</h5>
        <div class="form-group">
          <label class="form-label">配置参数</label>
          <textarea v-model="step.config" class="tech-textarea" rows="6" 
                    placeholder="请输入步骤的具体配置参数"></textarea>
        </div>
      </div>
    </div>

    <div class="editor-footer">
      <button class="tech-btn secondary" @click="$emit('cancel')">取消</button>
      <button class="tech-btn primary" @click="saveStep">保存</button>
    </div>
  </div>
</template>

<script setup>
defineProps({
  step: {
    type: Object,
    required: true
  }
})

const emit = defineEmits(['save', 'cancel'])

function resetConfig() {
  // 根据步骤类型重置配置
  const defaultConfigs = {
    validate: { rules: '', errorMessage: '' },
    db_query: { sql: '', dataSource: 'default', timeout: 30 },
    transform: { template: '', fieldMapping: '' },
    response: { statusCode: 200, headers: '', body: '' },
    condition: { condition: '', trueSteps: '', falseSteps: '' },
    cache: { key: '', ttl: 3600, action: 'get' },
    hook: { hookName: '', parameters: '' },
    loop: { items: '', template: '' }
  }
  
  if (defaultConfigs[this.step.type]) {
    this.step.config = { ...defaultConfigs[this.step.type] }
  } else {
    this.step.config = {}
  }
}

function saveStep() {
  try {
    // 验证配置格式
    if (this.step.config) {
      // 尝试解析JSON格式的配置
      if (typeof this.step.config === 'string') {
        this.step.config = JSON.parse(this.step.config)
      }
    }
    
    emit('save', this.step)
  } catch (error) {
    alert('配置格式错误，请检查JSON格式')
    console.error('配置解析错误:', error)
  }
}
</script>

<style scoped>
.step-editor {
  background: var(--surface-primary);
  border-radius: 8px;
  overflow: hidden;
}

.editor-header {
  padding: 16px 20px;
  border-bottom: 1px solid var(--border-light);
  background: var(--accent-primary);
  color: white;
}

.editor-title {
  margin: 0;
  font-size: 16px;
  font-weight: 600;
}

.editor-body {
  padding: 20px;
  max-height: 400px;
  overflow-y: auto;
}

.config-section {
  margin-bottom: 20px;
}

.section-title {
  margin: 0 0 12px 0;
  font-size: 14px;
  font-weight: 600;
  color: var(--text-secondary);
  border-bottom: 1px solid var(--border-light);
  padding-bottom: 8px;
}

.editor-footer {
  padding: 16px 20px;
  border-top: 1px solid var(--border-light);
  display: flex;
  gap: 12px;
  justify-content: flex-end;
}

.tech-textarea {
  width: 100%;
  padding: 12px;
  border: 1px solid var(--border-light);
  border-radius: 8px;
  background: var(--surface-secondary);
  color: var(--text-primary);
  font-family: 'Courier New', monospace;
  font-size: 13px;
  line-height: 1.4;
  resize: vertical;
  transition: border-color 0.2s;
}

.tech-textarea:focus {
  outline: none;
  border-color: var(--accent-primary);
  box-shadow: 0 0 0 2px var(--accent-light);
}
</style>