<template>
  <div class="card">
    <h3>文件上传</h3>
    <div style="display:flex;gap:12px;align-items:center;margin-top:8px">
      <input type="file" multiple @change="onSel" />
      <button class="btn" :disabled="loading || files.length===0" @click="send">上传</button>
    </div>
    <ul style="margin-top:10px">
      <li v-for="id in ids" :key="id">{{ id }}</li>
    </ul>
  </div>
</template>
<script setup>
import { ref } from 'vue'
const files = ref([])
const ids = ref([])
const loading = ref(false)
function onSel(e){
  files.value = Array.from(e.target.files || [])
}
async function send(){
  loading.value = true
  try{
    const fd = new FormData()
    files.value.forEach(f => fd.append('files', f))
    const resp = await fetch('/api/files/upload', { method: 'POST', body: fd })
    const data = await resp.json()
    ids.value = Array.isArray(data?.data) ? data.data : []
  } finally { loading.value = false }
}
</script>