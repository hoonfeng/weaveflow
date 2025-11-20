<template>
  <div>
    <h2>OpenAPI</h2>
    <div v-if="loading">加载中...</div>
    <pre v-else class="code">{{ json }}</pre>
  </div>
</template>
<script setup>
import { ref, onMounted } from 'vue'
import api from '../api'
const loading = ref(false)
const json = ref('')
onMounted(async () => {
  loading.value = true
  try {
    const { data } = await api.get('/openapi.json')
    json.value = JSON.stringify(data, null, 2)
  } finally {
    loading.value = false
  }
})
</script>
<style>
.code{white-space:pre-wrap; background:#f7f7f7; padding:12px; border:1px solid #eee}
</style>