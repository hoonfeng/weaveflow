<template>
  <div class="card card-glass">
    <div class="section-header"><h3 class="section-title">租户用量（月度）</h3></div>
    <div style="display:flex;gap:12px;align-items:center;margin-bottom:8px">
      <input v-model.number="tenantId" type="number" placeholder="租户ID" class="glass" style="padding:10px" />
      <button class="btn" @click="load">加载</button>
    </div>
    <table class="table-modern" cellspacing="0" cellpadding="6" style="width:100%">
      <thead><tr><th>年月</th><th>调用数</th></tr></thead>
      <tbody>
        <tr v-for="x in list" :key="x.ym"><td>{{ x.ym }}</td><td>{{ x.cnt }}</td></tr>
      </tbody>
    </table>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import api from '../utils/api.js'
const tenantId = ref(0)
const list = ref([])
async function load(){ if (!tenantId.value) return; const { data } = await api.get(`/api/admin/tenants/${tenantId.value}/usage/monthly`); list.value = data?.data || [] }
</script>