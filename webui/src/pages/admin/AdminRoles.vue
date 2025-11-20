<template>
  <div>
    <h3>角色管理</h3>
    <form @submit.prevent="create">
      <input v-model="name" placeholder="角色名" />
      <button :disabled="loading">创建</button>
    </form>
    <table border="1" cellspacing="0" cellpadding="6" style="margin-top:12px">
      <thead><tr><th>ID</th><th>名称</th><th>操作</th></tr></thead>
      <tbody>
        <tr v-for="r in list" :key="r.id">
          <td>{{ r.id }}</td>
          <td><input v-model="r.name" /></td>
          <td>
            <button @click="update(r)">更新</button>
            <button @click="del(r)">删除</button>
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
async function load(){ const { data } = await api.get('/api/admin/roles'); list.value = data?.data || [] }
onMounted(load)
async function create(){ loading.value=true; try{ await api.post('/api/admin/roles', { name: name.value }); name.value=''; await load() } finally{ loading.value=false } }
async function update(r){ await api.put(`/api/admin/roles/${r.id}`, { name: r.name }); await load() }
async function del(r){ await api.delete(`/api/admin/roles/${r.id}`); await load() }
</script>