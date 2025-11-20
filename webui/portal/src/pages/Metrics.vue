<template>
  <div class="card">
    <h3>指标</h3>
    <pre class="code">{{ text }}</pre>
  </div>
</template>
<script setup>
import { ref, onMounted } from 'vue'
import api from '../../../packages/api'
const text = ref(''); onMounted(async()=>{ const { data } = await api.get('/metrics', { responseType:'text' }); text.value = typeof data==='string'?data:'' })
</script>
<style>.code{white-space:pre-wrap}</style>