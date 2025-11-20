<template>
  <div class="card">
    <h3>插件文本反转</h3>
    <div style="display:flex;gap:12px;align-items:center;margin-top:8px">
      <input v-model="text" placeholder="输入文本" class="glass" style="padding:10px;flex:1" />
      <button class="btn" :disabled="loading" @click="run">调用</button>
    </div>
    <p v-if="out" style="margin-top:8px">结果: {{ out }}</p>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import api from '../api'
const text = ref('Hello')
const out = ref('')
const loading = ref(false)
async function run(){
  loading.value = true
  try{
    const { data } = await api.post('/api/text/reverse', { text: text.value })
    out.value = data?.data || ''
  } finally {
    loading.value = false
  }
}
</script>