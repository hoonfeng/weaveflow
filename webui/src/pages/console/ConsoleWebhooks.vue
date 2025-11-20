<template>
  <div>
    <h3>Webhook 管理</h3>
    <div>
      <input v-model.number="tenant_id" type="number" placeholder="租户ID" />
      <input v-model="url" placeholder="回调URL" />
      <input v-model="secret" placeholder="签名密钥(可选)" />
      <button @click="add">添加端点</button>
      <button @click="listEndpoints">查询端点</button>
    </div>
    <div style="margin-top:12px">
      <h4>任务</h4>
      <button @click="listTasks">查询任务</button>
      <input v-model.number="taskId" type="number" placeholder="任务ID" />
      <button @click="retry">重试任务</button>
    </div>
    <pre style="margin-top:12px">{{ out }}</pre>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import api from '../../api'
const tenant_id = ref(0)
const url = ref('')
const secret = ref('')
const taskId = ref(0)
const out = ref('')
async function add(){ const { data } = await api.post('/api/webhooks', { tenant_id: tenant_id.value, url: url.value, secret: secret.value }); out.value = JSON.stringify(data?.data||data,null,2) }
async function listEndpoints(){ const { data } = await api.get('/api/webhooks', { params: { tenant_id: tenant_id.value } }); out.value = JSON.stringify(data?.data||data,null,2) }
async function listTasks(){ const { data } = await api.get('/api/webhooks/tasks', { params: { tenant_id: tenant_id.value } }); out.value = JSON.stringify(data?.data||data,null,2) }
async function retry(){ const { data } = await api.post('/api/webhooks/tasks/retry', { id: taskId.value }); out.value = JSON.stringify(data?.data||data,null,2) }
</script>