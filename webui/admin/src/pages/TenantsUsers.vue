<template>
  <div class="card card-glass">
    <div class="section-header"><h3 class="section-title">租户成员</h3></div>
    <div style="display:flex;gap:12px;align-items:center;margin-bottom:8px">
      <input v-model.number="tenantId" type="number" placeholder="租户ID" class="glass" style="padding:10px" />
      <button class="btn" @click="load">加载</button>
    </div>
    <table class="table-modern" cellspacing="0" cellpadding="6" style="width:100%">
      <thead><tr><th>用户ID</th><th>用户名</th><th>邮箱</th><th>角色</th><th>操作</th></tr></thead>
      <tbody>
        <tr v-for="m in members" :key="m.uid">
          <td>{{ m.uid }}</td>
          <td>{{ m.username }}</td>
          <td>{{ m.email }}</td>
          <td>{{ m.role }}</td>
          <td>
            <button class="btn" @click="unbind(m.uid)">解绑</button>
          </td>
        </tr>
      </tbody>
    </table>
    <div style="display:flex;gap:8px;align-items:center;margin-top:8px">
      <input v-model.number="bindUserId" type="number" placeholder="用户ID" class="glass" style="padding:10px" />
      <input v-model="bindRole" placeholder="角色" class="glass" style="padding:10px" />
      <button class="btn" @click="bind">绑定</button>
    </div>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import api from '../utils/api.js'
const tenantId = ref(0)
const members = ref([])
async function load(){ if (!tenantId.value) return; const { data } = await api.get(`/api/admin/tenants/${tenantId.value}/users`); members.value = data?.data || [] }
const bindUserId = ref(0)
const bindRole = ref('member')
async function bind(){ if (!tenantId.value || !bindUserId.value) return; await api.post(`/api/admin/tenants/${tenantId.value}/users`, { user_id: bindUserId.value, role: bindRole.value }); await load() }
async function unbind(uid){ if (!tenantId.value || !uid) return; await api.delete(`/api/admin/tenants/${tenantId.value}/users/${uid}`); await load() }
</script>