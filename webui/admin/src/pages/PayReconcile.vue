<template>
  <div class="card">
    <div class="section-header"><div class="section-title">支付对账</div></div>
    <div style="display:flex;gap:12px;align-items:center;margin-bottom:12px">
      <input v-model="start" placeholder="开始 YYYY-MM-DD" style="padding:8px;border:1px solid #cbd5e1;border-radius:8px" />
      <input v-model="end" placeholder="结束 YYYY-MM-DD" style="padding:8px;border:1px solid #cbd5e1;border-radius:8px" />
      <button class="btn" @click="go">对账</button>
    </div>
    <table class="table-modern" v-if="rows.length">
      <thead><tr><th>日期</th><th>渠道</th><th>数量</th></tr></thead>
      <tbody>
        <tr v-for="(r,idx) in rows" :key="idx">
          <td>{{ r.d || r.date }}</td>
          <td>{{ r.provider }}</td>
          <td>{{ r.cnt }}</td>
        </tr>
      </tbody>
    </table>
    <div v-else>暂无数据</div>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import { payApi } from '../utils/api.js'
const today = () => new Date().toISOString().slice(0,10)
const start = ref(today())
const end = ref(today())
const rows = ref([])
async function go(){ const r = await payApi.reconcile({ start: start.value, end: end.value }); rows.value = r?.data ?? r ?? [] }
</script>
<style scoped>
</style>