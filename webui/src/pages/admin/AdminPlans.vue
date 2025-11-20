<template>
  <div>
    <h3>套餐计划</h3>
    <form @submit.prevent="create">
      <input v-model="name" placeholder="名称" />
      <input v-model.number="price" type="number" placeholder="价格" />
      <button :disabled="loading">创建</button>
    </form>
    <table border="1" cellspacing="0" cellpadding="6" style="margin-top:12px">
      <thead><tr><th>ID</th><th>名称</th><th>价格</th><th>操作</th></tr></thead>
      <tbody>
        <tr v-for="p in list" :key="p.id">
          <td>{{ p.id }}</td>
          <td><input v-model="p.name" /></td>
          <td><input v-model.number="p.price" type="number" /></td>
          <td>
            <button @click="update(p)">更新</button>
            <button @click="del(p)">删除</button>
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
const price = ref(0)
const list = ref([])
async function load(){ const { data } = await api.get('/api/admin/plans'); list.value = data?.data || [] }
onMounted(load)
async function create(){ loading.value=true; try{ await api.post('/api/admin/plans', { name: name.value, price: price.value }); name.value=''; price.value=0; await load() } finally{ loading.value=false } }
async function update(p){ await api.put(`/api/admin/plans/${p.id}`, { name: p.name, price: p.price }); await load() }
async function del(p){ await api.delete(`/api/admin/plans/${p.id}`); await load() }
</script>