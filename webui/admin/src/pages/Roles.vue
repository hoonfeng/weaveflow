<template>
  <div class="roles-page">
    <div class="page-header">
      <h1>角色管理</h1>
      <div class="header-actions">
        <button class="btn btn-primary" @click="openCreate=true">创建角色</button>
        <button class="btn btn-secondary" @click="loadRoles">刷新</button>
        <button class="btn btn-info" @click="syncPerms">同步权限</button>
      </div>
    </div>

    <div class="card" v-if="permOpen">
      <div class="card-header"><h3>权限 - {{ currentRole?.name }}</h3></div>
      <div class="roles-table">
        <div style="display:flex;gap:16px">
          <div style="flex:1">
            <h4>全部权限</h4>
            <div style="display:flex;gap:8px;margin:8px 0">
              <button class="btn btn-sm btn-outline" @click="selectAllPerms">全选</button>
              <button class="btn btn-sm btn-outline" @click="invertPerms">反选</button>
              <button class="btn btn-sm btn-outline" @click="clearPerms">清空</button>
            </div>
            <div style="display:flex;flex-direction:column;gap:8px;max-height:50vh;overflow:auto">
              <label v-for="p in allPerms" :key="p.id" style="display:flex;gap:8px;align-items:center">
                <input type="checkbox" :value="p.id" v-model="selectedPermIds" />
                <span>{{ p.code }}（{{ p.name }}）</span>
              </label>
            </div>
          </div>
          <div style="flex:1">
            <h4>已分配</h4>
            <div>
              <span v-for="rp in rolePerms" :key="rp.perm_id" class="chip">
                <span class="chip-title">{{ rp.name || rp.code }}</span>
                <span class="chip-code">{{ rp.code }}</span>
              </span>
            </div>
          </div>
        </div>
        <div style="display:flex;gap:8px;justify-content:flex-end;margin-top:12px">
          <button class="btn btn-secondary" @click="permOpen=false">关闭</button>
          <button class="btn btn-primary" @click="savePerms">保存</button>
        </div>
      </div>
    </div>
    <div class="card">
      <div class="card-header" style="display:flex;justify-content:space-between;align-items:center">
        <h3>角色列表</h3>
        <div class="toolbar" style="display:flex;gap:8px;align-items:center">
          <input v-model="query" @input="onSearchInput" placeholder="搜索角色名称" class="glass" style="padding:8px 10px;width:240px" />
        </div>
      </div>
      <div class="roles-table">
        <table class="users-table">
          <thead><tr><th>ID</th><th>名称</th><th>操作</th></tr></thead>
          <tbody>
            <tr v-for="r in filteredRoles" :key="r.id">
              <td>{{ r.id }}</td>
              <td>{{ r.name }}</td>
              <td class="user-actions">
                <button class="btn btn-sm btn-info" @click="viewMembers(r)">成员</button>
                <button class="btn btn-sm btn-warning" @click="viewPerms(r)">权限</button>
                <button class="btn btn-sm btn-secondary" @click="openEdit(r)">编辑</button>
                <button class="btn btn-sm btn-danger" @click="confirmRemoveRole(r)">删除</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <div class="card" v-if="membersOpen">
      <div class="card-header"><h3>成员 - {{ currentRole?.name }}</h3></div>
      <div class="roles-table">
        <table class="users-table">
          <thead><tr><th>用户ID</th><th>角色ID</th><th>角色名</th></tr></thead>
          <tbody>
            <tr v-for="m in members" :key="m.user_id+'-'+m.role_id"><td>{{ m.user_id }}</td><td>{{ m.role_id }}</td><td>{{ m.role_name }}</td></tr>
          </tbody>
        </table>
      </div>
    </div>

    <div v-if="openCreate" class="modal" @click="openCreate=false">
      <div class="modal-content" @click.stop>
        <h3>创建角色</h3>
        <form @submit.prevent="createRole" class="user-form">
          <div class="form-group"><label>名称</label><input v-model="form.name" class="glass" placeholder="如 admin" /></div>
          <div class="form-actions" style="display:flex;gap:8px;justify-content:flex-end"><button type="button" class="btn btn-secondary" @click="openCreate=false">取消</button><button type="submit" class="btn btn-primary">创建</button></div>
        </form>
      </div>
    </div>

    <div v-if="openEditModal" class="modal" @click="openEditModal=false">
      <div class="modal-content" @click.stop>
        <h3>编辑角色</h3>
        <form @submit.prevent="saveRole" class="user-form">
          <div class="form-group"><label>名称</label><input v-model="editForm.name" class="glass" /></div>
          <div class="form-actions" style="display:flex;gap:8px;justify-content:flex-end"><button type="button" class="btn btn-secondary" @click="openEditModal=false">取消</button><button type="submit" class="btn btn-primary">保存</button></div>
        </form>
      </div>
    </div>
    <!-- Toast -->
    <div v-if="toast.show" :class="['toast', toast.type]">{{ toast.text }}</div>

    <!-- Confirm Modal -->
    <div v-if="confirm.open" class="modal" @click="closeConfirm">
      <div class="modal-content" @click.stop>
        <h3>{{ confirm.title }}</h3>
        <p style="margin:10px 0">{{ confirm.message }}</p>
        <div style="display:flex;gap:8px;justify-content:flex-end">
          <button class="btn btn-secondary" @click="closeConfirm">取消</button>
          <button class="btn btn-danger" @click="doConfirm">确定</button>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup>
import { ref, reactive, computed } from 'vue'
import { rolesApi } from '../utils/api.js'

const roles = ref([])
const members = ref([])
const membersOpen = ref(false)
const currentRole = ref(null)
const permOpen = ref(false)
const allPerms = ref([])
const rolePerms = ref([])
const selectedPermIds = ref([])
const query = ref('')
const filteredRoles = computed(() => {
  const q = query.value.trim().toLowerCase()
  if (!q) return roles.value
  return (roles.value||[]).filter(r => String(r.name||'').toLowerCase().includes(q))
})

const openCreate = ref(false)
const openEditModal = ref(false)
const form = reactive({ name: '' })
const editForm = reactive({ id: null, name: '' })
const toast = reactive({ show:false, text:'', type:'success' })
function showToast(text, type='success', ms=2000){ toast.text=text; toast.type=type; toast.show=true; setTimeout(()=>{ toast.show=false }, ms) }
const confirm = reactive({ open:false, title:'确认删除', message:'此操作不可恢复，确定删除该角色？', onYes:null })
function confirmRemoveRole(role){ confirm.title='确认删除'; confirm.message=`确定删除角色 ${role.name} ?`; confirm.onYes = async ()=>{ await rolesApi.deleteRole(role.id); showToast('删除成功'); await loadRoles() }; confirm.open=true }
function closeConfirm(){ confirm.open=false; confirm.onYes=null }
async function doConfirm(){ if (confirm.onYes) { try { await confirm.onYes() } catch(e){ showToast(e.message||'操作失败','error') } } confirm.open=false }
function onSearchInput(){ /* debounced search could be added */ }

async function loadRoles(){ const res = await rolesApi.getRoles(); roles.value = res?.data || res || [] }
function viewMembers(role){ currentRole.value = role; membersOpen.value = true; loadMembers(role.id) }
async function loadMembers(roleId){ const res = await rolesApi.getRoleMembers(roleId); members.value = res?.data || res || [] }
async function viewPerms(role){ currentRole.value = role; permOpen.value = true; await loadPerms(role.id) }
async function loadPerms(roleId){ const [all, rp] = await Promise.all([rolesApi.getPermissions(), rolesApi.getRolePermissions(roleId)]) ; allPerms.value = all?.data || all || []; rolePerms.value = rp?.data || rp || []; selectedPermIds.value = rolePerms.value.map(x=>x.perm_id) }
async function savePerms(){ await rolesApi.setRolePermissions(currentRole.value.id, selectedPermIds.value); showToast('权限已保存'); await loadPerms(currentRole.value.id) }
function selectAllPerms(){ selectedPermIds.value = (allPerms.value||[]).map(p=>p.id) }
function invertPerms(){ const set = new Set(selectedPermIds.value||[]); selectedPermIds.value = (allPerms.value||[]).map(p=>p.id).filter(id=>!set.has(id)) }
function clearPerms(){ selectedPermIds.value = [] }
async function createRole(){ if (!form.name) { showToast('请填写角色名称','error'); return } ; await rolesApi.createRole({ name: form.name }); openCreate.value=false; form.name=''; showToast('创建成功'); await loadRoles() }
function openEdit(role){ openEditModal.value=true; editForm.id=role.id; editForm.name=role.name }
async function saveRole(){ if (!editForm.name) { showToast('名称不能为空','error'); return } ; await rolesApi.updateRole(editForm.id, { name: editForm.name }); openEditModal.value=false; showToast('保存成功'); await loadRoles() }
async function syncPerms(){ const res = await rolesApi.syncPermissions(); const scanned = res?.data?.scanned ?? 0; const synced = res?.data?.synced ?? 0; showToast(`已扫描 ${scanned}，同步 ${synced}`,'success',3000) }

loadRoles()
</script>
<style scoped>
.roles-page { padding: 2rem; max-width: 1000px; margin: 0 auto }
.page-header { display:flex; align-items:center; justify-content:space-between; margin-bottom: 1rem }
.card { background: white; border-radius: 12px; box-shadow: 0 4px 12px rgba(0,0,0,0.1); margin-bottom: 1rem }
.card-header { padding: 1rem 1.5rem; border-bottom: 1px solid #eee }
.users-table { width: 100%; border-collapse: collapse }
.users-table th { background: #f8f9fa; padding: 0.75rem; text-align:left; border-bottom: 2px solid #eee }
.users-table td { padding: 0.75rem; border-bottom: 1px solid #eee }
.user-actions { display:flex; gap: 8px }
.modal { position: fixed; inset: 0; background: rgba(0,0,0,0.45); display: flex; align-items: center; justify-content: center; z-index: 1000 }
.modal-content { background: #fff; border-radius: 12px; width: 520px; max-width: 90vw; padding: 16px; box-shadow: 0 10px 30px rgba(0,0,0,0.15) }
.btn { padding: 0.5rem 1rem; border: none; border-radius: 6px; font-size: 0.9rem; font-weight: 500; cursor: pointer }
.btn-sm { padding: 0.35rem 0.6rem; font-size: 0.82rem }
.btn-outline { background: transparent; border: 1px solid #9ca3af; color: #374151 }
.btn-primary { background-color: #3b82f6; color: #fff }
.btn-secondary { background-color: #6b7280; color: #fff }
.btn-info { background-color: #06b6d4; color: #fff }
.btn-danger { background-color: #ef4444; color: #fff }
.btn-warning { background-color: #f59e0b; color: #fff }
.user-form { padding: 1rem }
.form-group { display:flex; flex-direction:column; gap:6px; margin-bottom: 10px }
.form-group input { padding: 0.6rem; border: 1px solid #ddd; border-radius: 6px }
.toast { position: fixed; right: 16px; bottom: 16px; padding: 10px 14px; border-radius: 8px; box-shadow: 0 4px 10px rgba(0,0,0,0.1); background: #eef2f7; color: #111827; z-index: 1100 }
.toast.success { background: #dcfce7; color: #166534 }
.toast.error { background: #fee2e2; color: #991b1b }
.chip { display:inline-block; padding: 6px 10px; border-radius: 10px; background:#e2e8f0; color:#111827; margin-right:8px; font-size:0.8rem; line-height:1.2 }
.chip .chip-title { display:block; font-weight:600 }
.chip .chip-code { display:block; font-size:0.75rem; opacity:0.8 }
</style>