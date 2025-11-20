<template>
  <div>
    <h2 style="margin-bottom:8px">接口文档</h2>
    <div v-if="loading" class="card">加载中...</div>
    <div v-else class="grid">
      <div class="card" v-for="e in list" :key="e.module + e.method + e.path">
        <div style="display:flex;justify-content:space-between;align-items:center">
          <strong>{{ e.title || e.endpoint }}</strong>
          <span class="tag" :style="tagStyle(e.method)">{{ e.method }}</span>
        </div>
        <div style="color:var(--muted);margin:6px 0">{{ e.path }}</div>
        <div style="display:flex;gap:12px;color:var(--muted)">
          <span>模块: {{ e.module }}</span>
          <span>认证: {{ e.auth || 'none' }}</span>
        </div>
      </div>
    </div>
  </div>
  </template>
<script setup>
import { ref, onMounted } from 'vue'
import api from '../api'
const loading = ref(false)
const list = ref([])
onMounted(async () => {
  loading.value = true
  try {
    const { data } = await api.get('/docs')
    list.value = Array.isArray(data) ? data : []
  } finally {
    loading.value = false
  }
})
function tagStyle(m){
  const mm = String(m||'').toUpperCase()
  const c = mm==='GET'? '#22c55e' : mm==='POST'? '#3b82f6' : mm==='PUT'? '#f59e0b' : '#ef4444'
  return { padding:'4px 8px', borderRadius:'10px', background:c, color:'#fff', fontSize:'12px' }
}
</script>