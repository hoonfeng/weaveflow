<template>
  <div class="card">
    <h3>内置插件</h3>
    <button class="btn" @click="load">刷新</button>
    <div style="margin-top:10px">
      <table class="glass" style="width:100%">
        <thead><tr><th>名称</th><th>端点</th></tr></thead>
        <tbody>
          <tr v-for="it in list" :key="it.name">
            <td>{{ it.name }}</td>
            <td>{{ (it.endpoints||[]).join(', ') }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import api from '../utils/api.js'
const list = ref([])
async function load(){ const res = await api.get('/api/admin/plugins/internal'); list.value = res?.data || res }
load()
</script>