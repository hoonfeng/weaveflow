<template>
  <div class="card">
    <h3>Webhook 管理</h3>
    <div style="display:flex;gap:12px;align-items:center;margin-top:8px">
      <input v-model.number="tenantId" type="number" placeholder="租户ID" class="glass" style="padding:10px" />
      <input v-model="url" placeholder="回调URL" class="glass" style="padding:10px;flex:1" />
      <input v-model="secret" placeholder="签名密钥(可选)" class="glass" style="padding:10px" />
      <button class="btn" @click="add">添加端点</button>
      <button class="btn" @click="listEndpoints">查询端点</button>
    </div>
    <div style="display:flex;gap:12px;align-items:center;margin-top:8px">
      <input v-model.number="taskId" type="number" placeholder="任务ID" class="glass" style="padding:10px" />
      <button class="btn" @click="listTasks">查询任务</button>
      <button class="btn" @click="retry">重试任务</button>
    </div>
    <pre style="margin-top:10px">{{ out }}</pre>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import api, { setTenant } from '../../../../packages/api'
const tenantId = ref(0); const url = ref(''); const secret = ref(''); const taskId = ref(0); const out = ref('')
function apply(){ if(tenantId.value) setTenant(tenantId.value) }
async function add(){ apply(); const { data } = await api.post('/api/webhooks', { tenant_id: tenantId.value, url: url.value, secret: secret.value }); out.value = JSON.stringify(data?.data||data,null,2) }
async function listEndpoints(){ apply(); const { data } = await api.get('/api/webhooks', { params: { tenant_id: tenantId.value } }); out.value = JSON.stringify(data?.data||data,null,2) }
async function listTasks(){ apply(); const { data } = await api.get('/api/webhooks/tasks', { params: { tenant_id: tenantId.value } }); out.value = JSON.stringify(data?.data||data,null,2) }
async function retry(){ apply(); const { data } = await api.post('/api/webhooks/tasks/retry', { id: taskId.value }); out.value = JSON.stringify(data?.data||data,null,2) }
</script>