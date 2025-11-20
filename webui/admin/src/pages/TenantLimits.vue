<template>
  <div class="card card-glass">
    <div class="section-header"><h3 class="section-title">租户限流</h3></div>
    <div style="display:flex;gap:12px;align-items:center;margin-bottom:8px">
      <input v-model.number="tenantId" type="number" placeholder="租户ID" class="glass" style="padding:10px" />
      <button class="btn" @click="load">查询</button>
    </div>
    <div style="display:flex;gap:12px;align-items:center;margin-bottom:8px">
      <input v-model.number="rps" type="number" step="0.1" placeholder="rps" class="glass" style="padding:10px" />
      <input v-model.number="burst" type="number" placeholder="burst" class="glass" style="padding:10px" />
      <button class="btn" @click="save">保存</button>
    </div>
    <pre style="margin-top:6px">{{ JSON.stringify(current,null,2) }}</pre>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import api from '../utils/api.js'
const tenantId = ref(0)
const rps = ref(0)
const burst = ref(0)
const current = ref([])
async function load(){ if (!tenantId.value) return; const { data } = await api.get('/api/admin/tenant_limits', { params: { tenant_id: tenantId.value } }); current.value = data?.data || []; const cur = current.value[0] || {}; rps.value = Number(cur.rps)||0; burst.value = Number(cur.burst)||0 }
async function save(){ if (!tenantId.value) return; await api.post('/api/admin/tenant_limits/upsert', { tenant_id: tenantId.value, rps: rps.value, burst: burst.value }); await load() }
</script>