<template>
  <div class="card card-glass">
    <div class="section-header"><h3 class="section-title">租户 API Keys</h3></div>
    <div style="display:flex;gap:12px;align-items:center;margin-bottom:8px">
      <input v-model.number="tenantId" type="number" placeholder="租户ID" class="glass" style="padding:10px" />
      <button class="btn" @click="load">加载</button>
    </div>
    <table class="table-modern" cellspacing="0" cellpadding="6" style="width:100%">
      <thead><tr><th>ID</th><th>名称</th><th>状态</th><th>过期</th><th>作用域</th><th>操作</th></tr></thead>
      <tbody>
        <tr v-for="k in list" :key="k.id">
          <td>{{ k.id }}</td>
          <td>{{ k.name }}</td>
          <td>{{ k.status }}</td>
          <td>{{ k.expired_at }}</td>
          <td>{{ k.scopes }}</td>
          <td style="display:flex;gap:8px">
            <button class="btn" @click="toggle(k, k.status !== 'active')">{{ k.status === 'active' ? '禁用' : '启用' }}</button>
            <button class="btn" @click="rotate(k)">轮换</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import api from '../utils/api.js'
const tenantId = ref(0)
const list = ref([])
async function load(){ if (!tenantId.value) return; const { data } = await api.get(`/api/admin/tenants/${tenantId.value}/apikeys`); list.value = data?.data || [] }
async function toggle(k, enabled){ if (!tenantId.value) return; await api.post('/api/admin/apikeys/toggle', { tenant_id: tenantId.value, name: k.name, enabled }); await load() }
async function rotate(k){ if (!tenantId.value) return; await api.post('/api/admin/apikeys/rotate', { tenant_id: tenantId.value, name: k.name }); await load() }
</script>