<template>
  <div class="card">
    <h3>OpenAPI</h3>
    <pre class="code">{{ json }}</pre>
  </div>
</template>
<script setup>
import { ref, onMounted } from 'vue'
import api from '../../../packages/api'
const json = ref(''); onMounted(async()=>{ const { data } = await api.get('/openapi.json'); json.value = JSON.stringify(data,null,2) })
</script>
<style>
.code{white-space:pre-wrap}
</style>