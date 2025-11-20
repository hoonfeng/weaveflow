<template>
  <div class="card card-glass">
    <div class="section-header"><h3 class="section-title">租户订阅</h3></div>
    <div style="display:flex;gap:12px;align-items:center;margin-bottom:8px">
      <input v-model.number="tenantId" type="number" placeholder="租户ID" class="glass" style="padding:10px" />
      <button class="btn" @click="load">加载</button>
    </div>
    <div style="display:flex;gap:12px;align-items:center;margin-bottom:8px">
      <input v-model.number="planId" type="number" placeholder="计划ID" class="glass" style="padding:10px" />
      <button class="btn" @click="create">创建订阅</button>
      <button class="btn" @click="changePlan">变更计划</button>
      <button class="btn" @click="cancel">取消订阅</button>
    </div>
    <h4 class="section-title" style="margin:8px 0">当前订阅</h4>
    <pre style="margin-top:6px">{{ JSON.stringify(current,null,2) }}</pre>
    <h4 class="section-title" style="margin:12px 0">历史订阅</h4>
    <table class="table-modern" cellspacing="0" cellpadding="6" style="width:100%">
      <thead><tr><th>ID</th><th>计划</th><th>状态</th><th>开始</th><th>结束</th></tr></thead>
      <tbody>
        <tr v-for="s in list" :key="s.id"><td>{{ s.id }}</td><td>{{ s.plan_id }}</td><td>{{ s.status }}</td><td>{{ s.start_at }}</td><td>{{ s.end_at }}</td></tr>
      </tbody>
    </table>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import api from '../utils/api.js'
import { subscriptionsApi } from '../utils/api.js'
const tenantId = ref(0)
const planId = ref(0)
const current = ref({})
const list = ref([])
async function load(){ if (!tenantId.value) return; const cur = await subscriptionsApi.current(tenantId.value); current.value = cur?.data || cur || {}; const { data } = await api.get(`/api/admin/tenants/${tenantId.value}/subscriptions`); list.value = data?.data || [] }
async function create(){ if (!tenantId.value || !planId.value) return; await subscriptionsApi.create({ tenant_id: tenantId.value, plan_id: planId.value }); await load() }
async function changePlan(){ if (!tenantId.value || !planId.value) return; await subscriptionsApi.change({ tenant_id: tenantId.value, plan_id: planId.value }); await load() }
async function cancel(){ if (!tenantId.value) return; await subscriptionsApi.cancel({ tenant_id: tenantId.value }); await load() }
</script>