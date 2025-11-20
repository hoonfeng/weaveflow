<template>
  <div class="users-page">
    <!-- 页面标题和操作 -->
    <div class="page-header">
      <div class="header-content">
        <h1 class="page-title">用户管理</h1>
        <p class="page-description">管理系统所有用户账户，支持搜索、筛选和批量操作</p>
      </div>
      <div class="header-actions">
        <button class="btn btn-primary" @click="openCreateModal">
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M12 5v14M5 12h14" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
          </svg>
          创建用户
        </button>
      </div>
    </div>

    <!-- 筛选和搜索 -->
    <div class="filters-card">
      <div class="filters-row">
        <div class="search-box">
          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <circle cx="11" cy="11" r="8" stroke="currentColor" stroke-width="2"/>
            <path d="M21 21l-4.35-4.35" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
          </svg>
          <input
            v-model="searchQuery"
            type="text"
            placeholder="搜索用户名、邮箱或角色..."
            class="search-input"
            @input="handleSearch"
          />
        </div>
        
        <div class="filter-controls">
          <select v-model="statusFilter" class="filter-select" @change="loadUsers">
            <option value="">全部状态</option>
            <option value="active">活跃</option>
            <option value="inactive">禁用</option>
            <option value="pending">待审核</option>
          </select>
          
          <select v-model="roleFilter" class="filter-select" @change="loadUsers">
            <option value="">全部角色</option>
            <option value="admin">管理员</option>
            <option value="user">普通用户</option>
            <option value="guest">访客</option>
          </select>
        </div>
      </div>
    </div>

    <!-- 用户列表 -->
    <div class="users-card">
      <div class="card-header">
        <h3>用户列表</h3>
        <div class="card-actions">
          <button class="btn btn-sm btn-outline" @click="exportUsers">
            <svg width="14" height="14" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M14 2H6a2 2 0 00-2 2v16a2 2 0 002 2h12a2 2 0 002-2V8z" stroke="currentColor" stroke-width="2"/>
              <path d="M14 2v6h6" stroke="currentColor" stroke-width="2"/>
              <path d="M16 13H8" stroke="currentColor" stroke-width="2"/>
              <path d="M16 17H8" stroke="currentColor" stroke-width="2"/>
              <path d="M10 9H8" stroke="currentColor" stroke-width="2"/>
            </svg>
            导出
          </button>
          <button class="btn btn-sm btn-outline" @click="loadUsers" :disabled="loading">
            <svg width="14" height="14" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M23 4v6h-6" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
              <path d="M20.49 9a9 9 0 11-2.12-9.36" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
            </svg>
            刷新
          </button>
        </div>
      </div>

      <!-- 加载状态 -->
      <div v-if="loading && users.length === 0" class="loading-state">
        <div class="loading-content">
          <div class="loading-spinner"></div>
          <p>加载中...</p>
        </div>
      </div>

      <!-- 空状态 -->
      <div v-else-if="users.length === 0" class="empty-state">
        <div class="empty-content">
          <svg width="64" height="64" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M17 21v-2a4 4 0 00-4-4H5a4 4 0 00-4 4v2" stroke="currentColor" stroke-width="2"/>
            <circle cx="9" cy="7" r="4" stroke="currentColor" stroke-width="2"/>
            <path d="M23 21v-2a4 4 0 00-3-3.87" stroke="currentColor" stroke-width="2"/>
            <path d="M16 3.13a4 4 0 010 7.75" stroke="currentColor" stroke-width="2"/>
          </svg>
          <h4>暂无用户数据</h4>
          <p>当前没有符合条件的用户，请尝试调整筛选条件</p>
          <button class="btn btn-primary" @click="openCreateModal">创建第一个用户</button>
        </div>
      </div>

      <!-- 用户表格 -->
      <div v-else class="table-container">
        <table class="table">
          <thead>
            <tr>
              <th class="checkbox-column">
                <input type="checkbox" v-model="selectAll" @change="toggleSelectAll" />
              </th>
              <th>ID</th>
              <th>用户名</th>
              <th>邮箱</th>
              <th>角色</th>
              <th>状态</th>
              <th>创建时间</th>
              <th>操作</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="user in users" :key="user.id" :class="{ selected: selectedUsers.includes(user.id) }">
              <td class="checkbox-column">
                <input type="checkbox" :value="user.id" v-model="selectedUsers" />
              </td>
              <td class="user-id">{{ user.id }}</td>
              <td class="user-username">
                <div class="user-avatar">{{ user.username?.charAt(0)?.toUpperCase() }}</div>
                <span>{{ user.username }}</span>
              </td>
              <td class="user-email">{{ user.email }}</td>
              <td class="user-roles">
                <span v-for="role in parseRoles(user.roles || user.role)" :key="role" 
                      class="role-badge" :class="`role-${role.toLowerCase()}`">
                  {{ role }}
                </span>
              </td>
              <td class="user-status">
                <span class="status-badge" :class="user.status || 'active'">
                  {{ getStatusText(user.status) }}
                </span>
              </td>
              <td class="user-created">{{ formatDate(user.created_at) }}</td>
              <td class="user-actions">
                <div class="action-buttons">
                  <button class="btn btn-sm btn-outline" @click="editUser(user)" title="编辑">
                    <svg width="12" height="12" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                      <path d="M11 4H4a2 2 0 00-2 2v14a2 2 0 002 2h14a2 2 0 002-2v-7" stroke="currentColor" stroke-width="2"/>
                      <path d="M18.5 2.5a2.121 2.121 0 013 3L12 15l-4 1 1-4 9.5-9.5z" stroke="currentColor" stroke-width="2"/>
                    </svg>
                  </button>
                  <button class="btn btn-sm btn-outline" @click="viewUser(user)" title="查看">
                    <svg width="12" height="12" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                      <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z" stroke="currentColor" stroke-width="2"/>
                      <circle cx="12" cy="12" r="3" stroke="currentColor" stroke-width="2"/>
                    </svg>
                  </button>
                  <button class="btn btn-sm btn-outline danger" @click="deleteUser(user)" title="删除">
                    <svg width="12" height="12" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                      <path d="M3 6h18" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
                      <path d="M19 6v14a2 2 0 01-2 2H7a2 2 0 01-2-2V6m3 0V4a2 2 0 012-2h4a2 2 0 012 2v2" stroke="currentColor" stroke-width="2"/>
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>

        <!-- 分页 -->
        <div class="pagination" v-if="pagination.total > 0">
          <div class="pagination-info">
            显示 {{ pagination.start }} - {{ pagination.end }} 条，共 {{ pagination.total }} 条
          </div>
          <div class="pagination-controls">
            <button class="btn btn-sm btn-outline" :disabled="pagination.page === 1" @click="goToPage(pagination.page - 1)">
              上一页
            </button>
            <span class="pagination-current">第 {{ pagination.page }} 页</span>
            <button class="btn btn-sm btn-outline" :disabled="pagination.page >= pagination.pages" @click="goToPage(pagination.page + 1)">
              下一页
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- 批量操作栏 -->
    <div v-if="selectedUsers.length > 0" class="batch-actions">
      <div class="batch-info">
        已选择 {{ selectedUsers.length }} 个用户
      </div>
      <div class="batch-buttons">
        <button class="btn btn-sm btn-outline" @click="batchEnable">启用</button>
        <button class="btn btn-sm btn-outline" @click="batchDisable">禁用</button>
        <button class="btn btn-sm btn-outline danger" @click="batchDelete">删除</button>
        <button class="btn btn-sm btn-outline" @click="clearSelection">取消选择</button>
      </div>
    </div>

    <!-- 创建/编辑用户模态框 -->
    <div v-if="showModal" class="modal-overlay" @click="closeModal">
      <div class="modal" @click.stop>
        <div class="modal-header">
          <h3>{{ modalUser.id ? '编辑用户' : '创建用户' }}</h3>
          <button class="modal-close" @click="closeModal">
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M18 6L6 18M6 6l12 12" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
            </svg>
          </button>
        </div>
        
        <div class="modal-body">
          <form @submit.prevent="saveUser">
            <div class="form-group">
              <label class="form-label">用户名</label>
              <input v-model="modalUser.username" type="text" class="form-input" required />
            </div>
            
            <div class="form-group">
              <label class="form-label">邮箱</label>
              <input v-model="modalUser.email" type="email" class="form-input" required />
            </div>
            
            <div class="form-group">
              <label class="form-label">密码</label>
              <input v-model="modalUser.password" type="password" class="form-input" :required="!modalUser.id" />
            </div>
            
            <div class="form-row">
              <div class="form-group">
                <label class="form-label">角色</label>
                <select v-model="modalUser.role" class="form-select">
                  <option value="user">普通用户</option>
                  <option value="admin">管理员</option>
                </select>
              </div>
              
              <div class="form-group">
                <label class="form-label">状态</label>
                <select v-model="modalUser.status" class="form-select">
                  <option value="active">活跃</option>
                  <option value="inactive">禁用</option>
                </select>
              </div>
            </div>
            
            <div class="modal-actions">
              <button type="button" class="btn btn-outline" @click="closeModal">取消</button>
              <button type="submit" class="btn btn-primary" :disabled="saving">
                {{ saving ? '保存中...' : '保存' }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { usersApi } from '../utils/api.js'

const users = ref([])
const selectedUsers = ref([])
const searchQuery = ref('')
const statusFilter = ref('')
const roleFilter = ref('')
const loading = ref(false)
const saving = ref(false)
const showModal = ref(false)
const modalUser = ref({})
const selectAll = ref(false)

const pagination = ref({
  page: 1,
  pageSize: 20,
  total: 0,
  pages: 0,
  start: 0,
  end: 0
})

// 计算属性
const filteredUsers = computed(() => {
  return users.value.filter(user => {
    const matchesSearch = !searchQuery.value || 
      user.username.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      user.email.toLowerCase().includes(searchQuery.value.toLowerCase())
    
    const matchesStatus = !statusFilter.value || user.status === statusFilter.value
    const matchesRole = !roleFilter.value || (user.role || '').includes(roleFilter.value)
    
    return matchesSearch && matchesStatus && matchesRole
  })
})

// 方法
const loadUsers = async () => {
  loading.value = true
  try {
    const params = {
      page: pagination.value.page,
      page_size: pagination.value.pageSize,
      search: searchQuery.value
    }
    
    const response = await usersApi.list(params)
    // 后端返回格式: {code: 0, msg: "ok", data: {items: [], total: 0, page: 1, page_size: 20}}
    const responseData = response.data || response
    users.value = responseData.data?.items || responseData.data || responseData.items || []
    
    // 更新分页信息
    if (responseData.data) {
      pagination.value.total = responseData.data.total || responseData.total || users.value.length
      pagination.value.pages = Math.ceil(pagination.value.total / pagination.value.pageSize)
      pagination.value.start = (pagination.value.page - 1) * pagination.value.pageSize + 1
      pagination.value.end = Math.min(pagination.value.page * pagination.value.pageSize, pagination.value.total)
    }
  } catch (error) {
    console.error('加载用户列表失败:', error)
    users.value = []
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  // 防抖搜索
  clearTimeout(searchTimeout)
  searchTimeout = setTimeout(() => {
    pagination.value.page = 1
    loadUsers()
  }, 300)
}

let searchTimeout

const parseRoles = (roles) => {
  if (typeof roles === 'string') {
    return roles.split(',').filter(r => r.trim())
  }
  return roles || []
}

const getStatusText = (status) => {
  const statusMap = {
    active: '活跃',
    inactive: '禁用',
    pending: '待审核'
  }
  return statusMap[status] || '未知'
}

const formatDate = (dateString) => {
  if (!dateString) return '-'
  return new Date(dateString).toLocaleDateString('zh-CN')
}

const openCreateModal = () => {
  modalUser.value = { role: 'user', status: 'active' }
  showModal.value = true
}

const editUser = (user) => {
  modalUser.value = { ...user }
  showModal.value = true
}

const viewUser = (user) => {
  // 查看用户详情
  console.log('查看用户:', user)
}

const saveUser = async () => {
  saving.value = true
  try {
    if (modalUser.value.id) {
      await usersApi.updateUser(modalUser.value.id, modalUser.value)
    } else {
      await usersApi.createUser(modalUser.value)
    }
    closeModal()
    loadUsers()
  } catch (error) {
    console.error('保存用户失败:', error)
  } finally {
    saving.value = false
  }
}

const deleteUser = async (user) => {
  if (confirm(`确定要删除用户 "${user.username}" 吗？此操作不可撤销。`)) {
    try {
      await usersApi.deleteUser(user.id)
      loadUsers()
    } catch (error) {
      console.error('删除用户失败:', error)
    }
  }
}

const closeModal = () => {
  showModal.value = false
  modalUser.value = {}
}

const toggleSelectAll = () => {
  if (selectAll.value) {
    selectedUsers.value = users.value.map(user => user.id)
  } else {
    selectedUsers.value = []
  }
}

const batchEnable = async () => {
  // 批量启用用户
  console.log('批量启用:', selectedUsers.value)
}

const batchDisable = async () => {
  // 批量禁用用户
  console.log('批量禁用:', selectedUsers.value)
}

const batchDelete = async () => {
  if (confirm(`确定要删除选中的 ${selectedUsers.value.length} 个用户吗？此操作不可撤销。`)) {
    // 批量删除用户
    console.log('批量删除:', selectedUsers.value)
  }
}

const clearSelection = () => {
  selectedUsers.value = []
  selectAll.value = false
}

const goToPage = (page) => {
  if (page >= 1 && page <= pagination.value.pages) {
    pagination.value.page = page
    loadUsers()
  }
}

const exportUsers = () => {
  // 导出用户数据
  console.log('导出用户数据')
}

onMounted(() => {
  loadUsers()
})
</script>

<style scoped>
.users-page {
  display: flex;
  flex-direction: column;
  gap: var(--space-xl);
}

.page-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: var(--space-lg);
}

.header-content h1 {
  font-size: 2rem;
  font-weight: 700;
  margin-bottom: var(--space-xs);
}

.header-content p {
  color: var(--text-secondary);
}

.filters-card {
  background: var(--bg-card);
  border: 1px solid var(--border-light);
  border-radius: var(--radius-lg);
  padding: var(--space-lg);
}

.filters-row {
  display: flex;
  gap: var(--space-md);
  align-items: center;
}

.search-box {
  flex: 1;
  position: relative;
  display: flex;
  align-items: center;
}

.search-box svg {
  position: absolute;
  left: var(--space-md);
  color: var(--text-muted);
}

.search-input {
  width: 100%;
  padding: var(--space-md) var(--space-md) var(--space-md) calc(var(--space-md) * 2 + 20px);
  background: var(--bg-surface);
  border: 1px solid var(--border-light);
  border-radius: var(--radius-md);
  color: var(--text-primary);
  font-size: 0.875rem;
}

.search-input:focus {
  outline: none;
  border-color: var(--primary-500);
}

.filter-controls {
  display: flex;
  gap: var(--space-sm);
}

.filter-select {
  padding: var(--space-md);
  background: var(--bg-surface);
  border: 1px solid var(--border-light);
  border-radius: var(--radius-md);
  color: var(--text-primary);
  font-size: 0.875rem;
}

.users-card {
  background: var(--bg-card);
  border: 1px solid var(--border-light);
  border-radius: var(--radius-lg);
  overflow: hidden;
}

.card-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: var(--space-lg);
  border-bottom: 1px solid var(--border-light);
}

.card-actions {
  display: flex;
  gap: var(--space-sm);
}

.loading-state, .empty-state {
  padding: var(--space-3xl);
  text-align: center;
}

.loading-content, .empty-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: var(--space-md);
}

.loading-content p, .empty-content p {
  color: var(--text-muted);
}

.table-container {
  overflow-x: auto;
}

.table {
  width: 100%;
  border-collapse: collapse;
}

.table th {
  background: var(--bg-surface);
  padding: var(--space-md);
  text-align: left;
  font-size: 0.75rem;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  color: var(--text-secondary);
  border-bottom: 1px solid var(--border-light);
}

.table td {
  padding: var(--space-md);
  border-bottom: 1px solid var(--border-light);
  font-size: 0.875rem;
}

.table tbody tr:hover {
  background: rgba(76, 187, 76, 0.05);
}

.table tbody tr.selected {
  background: rgba(76, 187, 76, 0.1);
}

.checkbox-column {
  width: 40px;
  text-align: center;
}

.user-avatar {
  width: 32px;
  height: 32px;
  background: var(--gradient-primary);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-weight: 600;
  margin-right: var(--space-sm);
}

.user-username {
  display: flex;
  align-items: center;
  gap: var(--space-sm);
}

.role-badge {
  display: inline-block;
  padding: 2px 8px;
  font-size: 0.75rem;
  border-radius: var(--radius-full);
  margin-right: 4px;
}

.role-badge.role-admin {
  background: rgba(239, 68, 68, 0.1);
  color: var(--error);
}

.role-badge.role-user {
  background: rgba(76, 187, 76, 0.1);
  color: var(--primary-500);
}

.status-badge {
  padding: 4px 8px;
  font-size: 0.75rem;
  border-radius: var(--radius-full);
  font-weight: 500;
}

.status-badge.active {
  background: rgba(16, 185, 129, 0.1);
  color: var(--success);
}

.status-badge.inactive {
  background: rgba(107, 114, 128, 0.1);
  color: var(--gray-500);
}

.status-badge.pending {
  background: rgba(245, 158, 11, 0.1);
  color: var(--warning);
}

.action-buttons {
  display: flex;
  gap: var(--space-xs);
}

.action-buttons .btn.danger {
  color: var(--error);
  border-color: var(--error);
}

.action-buttons .btn.danger:hover {
  background: var(--error);
  color: white;
}

.pagination {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: var(--space-lg);
  border-top: 1px solid var(--border-light);
}

.pagination-info {
  color: var(--text-muted);
  font-size: 0.875rem;
}

.pagination-controls {
  display: flex;
  align-items: center;
  gap: var(--space-md);
}

.pagination-current {
  font-size: 0.875rem;
  color: var(--text-primary);
}

.batch-actions {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  background: var(--bg-surface);
  border-top: 1px solid var(--border-light);
  padding: var(--space-md) var(--space-xl);
  display: flex;
  align-items: center;
  justify-content: space-between;
  z-index: 100;
}

.batch-buttons {
  display: flex;
  gap: var(--space-sm);
}

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal {
  background: var(--bg-surface);
  border-radius: var(--radius-xl);
  width: 90%;
  max-width: 500px;
  max-height: 90vh;
  overflow-y: auto;
}

.modal-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: var(--space-xl);
  border-bottom: 1px solid var(--border-light);
}

.modal-close {
  background: none;
  border: none;
  color: var(--text-secondary);
  cursor: pointer;
  padding: var(--space-xs);
  border-radius: var(--radius-sm);
}

.modal-close:hover {
  background: var(--bg-card);
}

.modal-body {
  padding: var(--space-xl);
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: var(--space-md);
}

.modal-actions {
  display: flex;
  gap: var(--space-md);
  justify-content: flex-end;
  margin-top: var(--space-xl);
}

@media (max-width: 768px) {
  .page-header {
    flex-direction: column;
    align-items: flex-start;
    gap: var(--space-md);
  }
  
  .filters-row {
    flex-direction: column;
    align-items: stretch;
  }
  
  .form-row {
    grid-template-columns: 1fr;
  }
  
  .modal {
    width: 95%;
    margin: var(--space-md);
  }
}
</style>