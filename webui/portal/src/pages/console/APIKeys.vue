<template>
  <div class="card">
    <h3>API Keys</h3>
    <div style="display:flex;gap:12px;align-items:center;margin-top:8px">
      <input v-model.number="tenantId" type="number" placeholder="租户ID" class="glass" style="padding:10px" />
      <button class="btn" @click="rotate">旋转密钥</button>
      <button class="btn" @click="toggle">启用/禁用</button>
    </div>
    <pre style="margin-top:10px">{{ out }}</pre>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import api, { setTenant } from '../../../../packages/api'
const tenantId = ref(0); const out = ref('')
function apply(){ if(tenantId.value) setTenant(tenantId.value) }
async function rotate(){ apply(); const { data } = await api.post('/api/admin/apikeys/rotate', { tenant_id: tenantId.value }); out.value = JSON.stringify(data?.data||data,null,2) }
async function toggle(){ apply(); const { data } = await api.post('/api/admin/apikeys/toggle', { tenant_id: tenantId.value }); out.value = JSON.stringify(data?.data||data,null,2) }
</script>