<template>
  <div class="tenants-page">
    <div class="page-header">
      <h1>租户管理</h1>
      <p>管理系统所有租户账户</p>
    </div>

    <!-- 创建租户卡片 -->
    <div class="card create-tenant-card">
      <div class="card-header">
        <h3>创建新租户</h3>
      </div>
      
      <form @submit.prevent="createTenant" class="tenant-form">
        <div class="form-row">
          <div class="form-group">
            <label for="name">租户名称</label>
            <input
              id="name"
              v-model="form.name"
              type="text"
              placeholder="请输入租户名称"
              required
              :disabled="loading"
            />
          </div>
          
          <div class="form-group">
            <label for="domain">域名</label>
            <input
              id="domain"
              v-model="form.domain"
              type="text"
              placeholder="请输入域名（可选）"
              :disabled="loading"
            />
          </div>
        </div>

        <div class="form-row">
          <div class="form-group">
            <label for="contact_email">联系邮箱</label>
            <input
              id="contact_email"
              v-model="form.contact_email"
              type="email"
              placeholder="请输入联系邮箱"
              required
              :disabled="loading"
            />
          </div>

          <div class="form-group">
            <label for="status">状态</label>
            <select
              id="status"
              v-model="form.status"
              :disabled="loading"
            >
              <option value="active">活跃</option>
              <option value="inactive">禁用</option>
              <option value="pending">待审核</option>
            </select>
          </div>
        </div>

        <div class="form-group">
          <label for="description">描述</label>
          <textarea
            id="description"
            v-model="form.description"
            placeholder="请输入租户描述（可选）"
            :disabled="loading"
            rows="3"
          ></textarea>
        </div>

        <div v-if="errorMessage" class="error-message">
          {{ errorMessage }}
        </div>

        <div class="form-actions">
          <button 
            type="submit" 
            class="btn btn-primary"
            :disabled="loading || !formValid"
          >
            <span v-if="loading">创建中...</span>
            <span v-else>创建租户</span>
          </button>
        </div>
      </form>
    </div>

    <!-- 租户列表卡片 -->
    <div class="card tenants-list-card">
      <div class="card-header">
        <h3>租户列表</h3>
        <div class="header-actions">
          <div class="search-box">
            <input
              v-model="searchQuery"
              type="text"
              placeholder="搜索租户名称或域名..."
              @input="onSearchInput"
              :disabled="loading"
            />
          </div>
          <button class="btn btn-secondary" @click="loadTenants" :disabled="loading">
            <span v-if="loading">刷新中...</span>
            <span v-else>刷新</span>
          </button>
        </div>
      </div>

      <div v-if="loading && tenants.length === 0" class="loading-state">
        <p>加载中...</p>
      </div>

      <div v-else-if="tenants.length === 0" class="empty-state">
        <p>暂无租户数据</p>
      </div>

      <div v-else class="tenants-table-container">
        <table class="tenants-table">
          <thead>
            <tr>
              <th>ID</th>
              <th>租户名称</th>
              <th>域名</th>
              <th>联系邮箱</th>
              <th>状态</th>
              <th>创建时间</th>
              <th>操作</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="tenant in tenants" :key="tenant.id" class="tenant-row">
              <td class="tenant-id">{{ tenant.id }}</td>
              <td class="tenant-name">
                <input
                  v-model="tenant.name"
                  :disabled="editingTenant !== tenant.id"
                  @blur="editingTenant = null"
                />
              </td>
              <td class="tenant-domain">
                <input
                  v-model="tenant.domain"
                  :disabled="editingTenant !== tenant.id"
                  @blur="editingTenant = null"
                />
              </td>
              <td class="tenant-contact">
                <input
                  v-model="tenant.contact_email"
                  type="email"
                  :disabled="editingTenant !== tenant.id"
                  @blur="editingTenant = null"
                />
              </td>
              <td class="tenant-status">
                <span :class="['status-badge', tenant.status]">
                  {{ getStatusText(tenant.status) }}
                </span>
                <button
                  v-if="editingTenant !== tenant.id"
                  class="btn btn-sm btn-outline"
                  @click="toggleTenantStatus(tenant)"
                  :title="tenant.status === 'active' ? '禁用租户' : '启用租户'"
                >
                  {{ tenant.status === 'active' ? '禁用' : '启用' }}
                </button>
              </td>
              <td class="tenant-created">
                {{ formatDate(tenant.created_at) }}
              </td>
              <td class="tenant-actions">
                <button
                  v-if="editingTenant === tenant.id"
                  class="btn btn-success btn-sm"
                  @click="updateTenant(tenant)"
                >
                  保存
                </button>
                <button
                  v-else
                  class="btn btn-primary btn-sm"
                  @click="startEdit(tenant)"
                >
                  编辑
                </button>
                
                <button
                  class="btn btn-danger btn-sm"
                  @click="deleteTenant(tenant)"
                >
                  删除
                </button>
              </td>
            </tr>
          </tbody>
        </table>

        <!-- 分页控件 -->
        <div v-if="pagination.total > pagination.page_size" class="pagination">
          <button
            class="btn btn-secondary"
            :disabled="pagination.page === 1"
            @click="goToPage(pagination.page - 1)"
          >
            上一页
          </button>
          
          <span class="page-info">
            第 {{ pagination.page }} 页 / 共 {{ Math.ceil(pagination.total / pagination.page_size) }} 页
          </span>
          
          <button
            class="btn btn-secondary"
            :disabled="pagination.page >= Math.ceil(pagination.total / pagination.page_size)"
            @click="goToPage(pagination.page + 1)"
          >
            下一页
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { tenantsApi } from '../utils/api.js'

const loading = ref(false)
const errorMessage = ref('')
const tenants = ref([])
const editingTenant = ref(null)

const form = reactive({
  name: '',
  domain: '',
  contact_email: '',
  status: 'active',
  description: ''
})

const pagination = reactive({
  page: 1,
  page_size: 10,
  total: 0
})

const searchQuery = ref('')
const searchTimeout = ref(null)

// 计算表单是否有效
const formValid = computed(() => {
  return form.name && form.contact_email
})

// 加载租户列表
const loadTenants = async (page = 1) => {
  loading.value = true
  errorMessage.value = ''
  
  try {
    const params = {
      page,
      page_size: pagination.page_size
    }
    
    // 添加搜索参数
    if (searchQuery.value.trim()) {
      params.search = searchQuery.value.trim()
    }
    
    const response = await tenantsApi.list(params)
    
    // 后端返回格式: {code: 0, msg: "ok", data: []}
    const responseData = response.data || response
    tenants.value = responseData.data || responseData.items || []
    pagination.total = responseData.total || responseData.data?.length || tenants.value.length
    pagination.page = page
    
  } catch (error) {
    errorMessage.value = error.message || '加载租户列表失败'
    console.error('加载租户列表失败:', error)
  } finally {
    loading.value = false
  }
}

// 创建租户
const createTenant = async () => {
  if (!formValid.value) return
  
  loading.value = true
  errorMessage.value = ''
  
  try {
    await tenantsApi.createTenant(form)
    
    // 清空表单
    form.name = ''
    form.domain = ''
    form.contact_email = ''
    form.status = 'active'
    form.description = ''
    
    // 重新加载租户列表
    await loadTenants(pagination.page)
    
  } catch (error) {
    errorMessage.value = error.message || '创建租户失败'
    console.error('创建租户失败:', error)
  } finally {
    loading.value = false
  }
}

// 更新租户
const updateTenant = async (tenant) => {
  try {
    await tenantsApi.updateTenant(tenant.id, {
      name: tenant.name,
      domain: tenant.domain,
      contact_email: tenant.contact_email
    })
    
    editingTenant.value = null
    
  } catch (error) {
    errorMessage.value = error.message || '更新租户失败'
    console.error('更新租户失败:', error)
    
    // 重新加载以恢复原始数据
    await loadTenants(pagination.page)
  }
}

// 切换租户状态
const toggleTenantStatus = async (tenant) => {
  if (!confirm(`确定要${tenant.status === 'active' ? '禁用' : '启用'}租户 "${tenant.name}" 吗？`)) {
    return
  }
  
  try {
    await tenantsApi.updateTenant(tenant.id, {
      status: tenant.status === 'active' ? 'inactive' : 'active'
    })
    
    // 重新加载租户列表
    await loadTenants(pagination.page)
    
  } catch (error) {
    errorMessage.value = error.message || '切换租户状态失败'
    console.error('切换租户状态失败:', error)
  }
}

// 删除租户
const deleteTenant = async (tenant) => {
  if (!confirm(`确定要删除租户 "${tenant.name}" 吗？此操作不可恢复。`)) {
    return
  }
  
  try {
    await tenantsApi.deleteTenant(tenant.id)
    await loadTenants(pagination.page)
    
  } catch (error) {
    errorMessage.value = error.message || '删除租户失败'
    console.error('删除租户失败:', error)
  }
}

// 开始编辑
const startEdit = (tenant) => {
  editingTenant.value = tenant.id
}

// 分页导航
const goToPage = (page) => {
  if (page < 1 || page > Math.ceil(pagination.total / pagination.page_size)) {
    return
  }
  loadTenants(page)
}

// 搜索输入处理
const onSearchInput = () => {
  // 清除之前的定时器
  if (searchTimeout.value) {
    clearTimeout(searchTimeout.value)
  }
  
  // 设置新的定时器，延迟搜索
  searchTimeout.value = setTimeout(() => {
    loadTenants(1) // 搜索时回到第一页
  }, 300)
}

// 获取状态文本
const getStatusText = (status) => {
  const statusMap = {
    'active': '活跃',
    'inactive': '禁用',
    'pending': '待审核'
  }
  return statusMap[status] || status
}

// 格式化日期
const formatDate = (dateString) => {
  if (!dateString) return '-'
  
  try {
    const date = new Date(dateString)
    return date.toLocaleDateString('zh-CN')
  } catch {
    return dateString
  }
}

// 组件挂载时加载数据
onMounted(() => {
  loadTenants()
})
</script>

<style scoped>
.tenants-page {
  padding: 2rem;
  max-width: 1200px;
  margin: 0 auto;
}

.page-header {
  margin-bottom: 2rem;
}

.page-header h1 {
  margin: 0 0 0.5rem 0;
  color: #2d3436;
  font-size: 2rem;
  font-weight: 600;
}

.page-header p {
  margin: 0;
  color: #636e72;
  font-size: 1.1rem;
}

.card {
  background: white;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  margin-bottom: 2rem;
}

.card-header {
  padding: 1.5rem 1.5rem 1rem;
  border-bottom: 1px solid #eee;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.card-header h3 {
  margin: 0;
  color: #2d3436;
  font-size: 1.25rem;
}

.header-actions {
  display: flex;
  gap: 0.5rem;
  align-items: center;
}

.search-box {
  position: relative;
}

.search-box input {
  padding: 0.5rem 1rem;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 0.9rem;
  min-width: 200px;
}

.search-box input:focus {
  outline: none;
  border-color: #3498db;
  box-shadow: 0 0 0 2px rgba(52, 152, 219, 0.1);
}

.tenant-form {
  padding: 1.5rem;
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1rem;
  margin-bottom: 1rem;
}

.form-group {
  display: flex;
  flex-direction: column;
}

.form-group label {
  margin-bottom: 0.5rem;
  font-weight: 500;
  color: #2d3436;
}

.form-group input,
.form-group select,
.form-group textarea {
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 1rem;
  transition: border-color 0.2s;
}

.form-group input:focus,
.form-group select:focus,
.form-group textarea:focus {
  outline: none;
  border-color: #3498db;
  box-shadow: 0 0 0 2px rgba(52, 152, 219, 0.1);
}

.form-group input:disabled,
.form-group select:disabled,
.form-group textarea:disabled {
  background-color: #f5f5f5;
  cursor: not-allowed;
}

.form-actions {
  margin-top: 1.5rem;
  text-align: right;
}

.error-message {
  background-color: #fee;
  color: #c33;
  padding: 0.75rem;
  border-radius: 6px;
  margin-bottom: 1rem;
  font-size: 0.9rem;
}

.loading-state,
.empty-state {
  padding: 3rem;
  text-align: center;
  color: #636e72;
}

.tenants-table-container {
  padding: 0 1.5rem 1.5rem;
}

.tenants-table {
  width: 100%;
  border-collapse: collapse;
  margin-bottom: 1.5rem;
}

.tenants-table th {
  background-color: #f8f9fa;
  padding: 1rem;
  text-align: left;
  font-weight: 600;
  color: #2d3436;
  border-bottom: 2px solid #eee;
}

.tenants-table td {
  padding: 1rem;
  border-bottom: 1px solid #eee;
}

.tenant-row:hover {
  background-color: #f8f9fa;
}

.tenant-id {
  font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
  font-size: 0.9rem;
  color: #636e72;
}

.tenant-name input,
.tenant-domain input,
.tenant-contact input {
  width: 100%;
  padding: 0.5rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 0.9rem;
}

.tenant-status {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.status-badge {
  padding: 0.25rem 0.75rem;
  border-radius: 20px;
  font-size: 0.8rem;
  font-weight: 500;
}

.status-badge.active {
  background-color: #d4edda;
  color: #155724;
}

.status-badge.inactive {
  background-color: #f8d7da;
  color: #721c24;
}

.status-badge.pending {
  background-color: #fff3cd;
  color: #856404;
}

.tenant-created {
  color: #636e72;
  font-size: 0.9rem;
}

.tenant-actions {
  display: flex;
  gap: 0.5rem;
}

.btn {
  padding: 0.5rem 1rem;
  border: none;
  border-radius: 6px;
  font-size: 0.9rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
}

.btn-primary {
  background-color: #3498db;
  color: white;
}

.btn-primary:hover:not(:disabled) {
  background-color: #2980b9;
}

.btn-secondary {
  background-color: #95a5a6;
  color: white;
}

.btn-secondary:hover:not(:disabled) {
  background-color: #7f8c8d;
}

.btn-success {
  background-color: #27ae60;
  color: white;
}

.btn-success:hover:not(:disabled) {
  background-color: #229954;
}

.btn-danger {
  background-color: #e74c3c;
  color: white;
}

.btn-danger:hover:not(:disabled) {
  background-color: #c0392b;
}

.btn-outline {
  background-color: transparent;
  border: 1px solid #95a5a6;
  color: #7f8c8d;
}

.btn-outline:hover:not(:disabled) {
  background-color: #95a5a6;
  color: white;
}

.btn-sm {
  padding: 0.25rem 0.5rem;
  font-size: 0.8rem;
}

.btn:disabled {
  background-color: #bdc3c7;
  cursor: not-allowed;
}

.pagination {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 1rem;
  margin-top: 1.5rem;
}

.page-info {
  color: #636e72;
  font-size: 0.9rem;
}

@media (max-width: 768px) {
  .tenants-page {
    padding: 1rem;
  }
  
  .form-row {
    grid-template-columns: 1fr;
  }
  
  .header-actions {
    flex-direction: column;
    align-items: stretch;
  }
  
  .search-box input {
    min-width: auto;
    width: 100%;
  }
  
  .tenants-table {
    font-size: 0.8rem;
  }
  
  .tenants-table th,
  .tenants-table td {
    padding: 0.5rem;
  }
  
  .tenant-actions {
    flex-direction: column;
  }
  
  .pagination {
    flex-direction: column;
    gap: 0.5rem;
  }
}
</style>