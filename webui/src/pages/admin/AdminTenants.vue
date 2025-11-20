<template>
  <div>
    <h3>租户管理</h3>
    <form @submit.prevent="create">
      <input v-model="name" placeholder="租户名" />
      <button :disabled="loading">创建</button>
    </form>
    <table border="1" cellspacing="0" cellpadding="6" style="margin-top:12px">
      <thead><tr><th>ID</th><th>名称</th><th>操作</th></tr></thead>
      <tbody>
        <tr v-for="t in list" :key="t.id">
          <td>{{ t.id }}</td>
          <td><input v-model="t.name" /></td>
          <td>
            <button @click="update(t)">更新</button>
            <button @click="del(t)">删除</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
<script setup>
import { ref, onMounted } from 'vue'
import api from '../../api'
const loading = ref(false)
const name = ref('')
const list = ref([])
async function load(){ const { data } = await api.get('/api/admin/tenants'); list.value = data?.data || [] }
onMounted(load)
async function create(){ loading.value=true; try{ await api.post('/api/admin/tenants', { name: name.value }); name.value=''; await load() } finally{ loading.value=false } }
async function update(t){ await api.put(`/api/admin/tenants/${t.id}`, { name: t.name }); await load() }
async function del(t){ await api.delete(`/api/admin/tenants/${t.id}`); await load() }
</script>