<template>
  <div>
    <h3>API Key 管理</h3>
    <div>
      <input v-model="tenant_id" type="number" placeholder="租户ID" />
      <button @click="rotate">旋转密钥</button>
      <button @click="toggle">启用/禁用</button>
    </div>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import api from '../../api'
const tenant_id = ref(0)
async function rotate(){ await api.post('/api/admin/apikeys/rotate', { tenant_id: tenant_id.value }) }
async function toggle(){ await api.post('/api/admin/apikeys/toggle', { tenant_id: tenant_id.value }) }
</script>