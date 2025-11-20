<template>
  <div>
    <h2>指标</h2>
    <div v-if="loading">加载中...</div>
    <pre v-else class="code">{{ text }}</pre>
  </div>
</template>
<script setup>
import { ref, onMounted } from 'vue'
import api from '../api'
const loading = ref(false)
const text = ref('')
onMounted(async () => {
  loading.value = true
  try {
    const { data } = await api.get('/metrics', { responseType: 'text' })
    text.value = typeof data === 'string' ? data : ''
  } finally {
    loading.value = false
  }
})
</script>
<style>
.code{white-space:pre-wrap; background:#f7f7f7; padding:12px; border:1px solid #eee}
</style>