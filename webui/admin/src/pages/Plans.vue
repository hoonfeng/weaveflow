<template>
  <div class="grid">
    <div class="card">
      <h3>创建计划</h3>
      <form @submit.prevent="create" style="display:grid;gap:10px;margin-top:8px">
        <input v-model="name" placeholder="名称" class="glass" style="padding:10px" />
        <input v-model.number="price" type="number" placeholder="价格" class="glass" style="padding:10px" />
        <button class="btn" :disabled="loading">创建</button>
      </form>
    </div>
    <div class="card card-glass">
      <div class="section-header"><h3 class="section-title">计划列表</h3></div>
      <div v-if="list.length">
        <table class="table-modern" cellspacing="0" cellpadding="6" style="width:100%">
          <thead><tr><th>ID</th><th>名称</th><th>价格</th><th>操作</th></tr></thead>
          <tbody>
            <tr v-for="p in list" :key="p.id">
              <td>{{ p.id }}</td>
              <td><input v-model="p.name" class="glass" style="padding:8px" /></td>
              <td><input v-model.number="p.price" type="number" class="glass" style="padding:8px" /></td>
              <td>
                <button class="btn" @click="update(p)">更新</button>
                <button class="btn" @click="del(p)">删除</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <div v-else>无数据</div>
    </div>
  </div>
</template>
<script setup>
import { ref, onMounted } from 'vue'
import api from '../utils/api.js'
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