<template>
  <div>
    <h3>工具与运维</h3>
    <div style="display:flex; gap:12px; flex-wrap:wrap">
      <button @click="reload">热重载</button>
      <button @click="lint">接口 Lint</button>
      <button @click="queue">队列状态</button>
      <button @click="purgeNonces">清理 Nonces</button>
      <button @click="pluginsEnable">启用插件</button>
      <button @click="pluginsDisable">禁用插件</button>
    </div>
    <pre style="margin-top:12px">{{ out }}</pre>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import api from '../../api'
const out = ref('')
async function reload(){ const { data } = await api.post('/api/admin/reload'); out.value = JSON.stringify(data?.data||data,null,2) }
async function lint(){ const { data } = await api.get('/api/admin/lint'); out.value = JSON.stringify(data?.data||data,null,2) }
async function queue(){ const { data } = await api.get('/api/admin/queue/status'); out.value = JSON.stringify(data?.data||data,null,2) }
async function purgeNonces(){ const { data } = await api.post('/api/admin/nonces/purge'); out.value = JSON.stringify(data?.data||data,null,2) }
async function pluginsEnable(){ const { data } = await api.post('/api/admin/plugins/enable', { names: ['metrics'], enabled: true }); out.value = JSON.stringify(data?.data||data,null,2) }
async function pluginsDisable(){ const { data } = await api.post('/api/admin/plugins/disable', { names: ['metrics'], enabled: false }); out.value = JSON.stringify(data?.data||data,null,2) }
</script>