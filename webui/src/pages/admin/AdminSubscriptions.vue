<template>
  <div>
    <h3>订阅管理</h3>
    <form @submit.prevent="create">
      <input v-model.number="tenant_id" type="number" placeholder="租户ID" />
      <input v-model.number="plan_id" type="number" placeholder="计划ID" />
      <button :disabled="loading">创建订阅</button>
    </form>
    <div style="margin-top:12px">
      <button @click="load">刷新当前订阅</button>
      <pre>{{ JSON.stringify(current,null,2) }}</pre>
    </div>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import api from '../../api'
const loading = ref(false)
const tenant_id = ref(0)
const plan_id = ref(0)
const current = ref({})
async function create(){ loading.value=true; try{ await api.post('/api/admin/subscriptions', { tenant_id: tenant_id.value, plan_id: plan_id.value }); await load() } finally{ loading.value=false } }
async function load(){ const { data } = await api.get('/api/admin/subscriptions/current', { params: { tenant_id: tenant_id.value } }); current.value = data?.data || {} }
</script>