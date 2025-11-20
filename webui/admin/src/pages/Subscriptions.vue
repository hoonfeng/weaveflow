<template>
  <div class="card card-glass">
    <div class="section-header"><h3 class="section-title">订阅管理</h3></div>
    <form @submit.prevent="create" style="display:flex;gap:12px;align-items:center;margin-top:8px">
      <input v-model.number="tenant_id" type="number" placeholder="租户ID" class="glass" style="padding:10px" />
      <input v-model.number="plan_id" type="number" placeholder="计划ID" class="glass" style="padding:10px" />
      <button class="btn" :disabled="loading">创建订阅</button>
      <button class="btn" type="button" @click="load">查询当前订阅</button>
    </form>
    <pre style="margin-top:10px">{{ JSON.stringify(current,null,2) }}</pre>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import api from '../utils/api.js'
const loading = ref(false)
const tenant_id = ref(0)
const plan_id = ref(0)
const current = ref({})
async function create(){ loading.value=true; try{ await api.post('/api/admin/subscriptions', { tenant_id: tenant_id.value, plan_id: plan_id.value }); await load() } finally{ loading.value=false } }
async function load(){ const { data } = await api.get('/api/admin/subscriptions/current', { params: { tenant_id: tenant_id.value } }); current.value = data?.data || {} }
</script>
