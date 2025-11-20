<template>
  <div class="card">
    <h3>用量统计</h3>
    <div style="display:flex;gap:12px;align-items:center;margin-top:8px">
      <input v-model.number="tenantId" type="number" placeholder="租户ID" class="glass" style="padding:10px" />
      <button class="btn" @click="daily">查询日用量</button>
      <button class="btn" @click="monthly">查询月用量</button>
      <button class="btn" @click="reports">聚合报表</button>
    </div>
    <pre style="margin-top:10px">{{ out }}</pre>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import api, { setTenant } from '../../../../packages/api'
const tenantId = ref(0); const out = ref('')
function apply(){ if(tenantId.value) setTenant(tenantId.value) }
async function daily(){ apply(); const { data } = await api.get('/api/usage/daily', { params:{ tenant_id: tenantId.value } }); out.value = JSON.stringify(data?.data||data,null,2) }
async function monthly(){ apply(); const { data } = await api.get('/api/usage/monthly', { params:{ tenant_id: tenantId.value } }); out.value = JSON.stringify(data?.data||data,null,2) }
async function reports(){ apply(); const { data } = await api.post('/api/usage/reports', { tenant_id: tenantId.value }); out.value = JSON.stringify(data?.data||data,null,2) }
</script>