<template>
  <div class="card card-glass">
    <div class="section-header"><h3 class="section-title">订单管理</h3><div style="display:flex;gap:12px"><button class="btn" @click="load">刷新</button><button class="btn" @click="exportCsv">导出 CSV</button><button class="btn" @click="exportJson">导出 JSON</button></div></div>
    <table class="table-modern" cellspacing="0" cellpadding="6" style="width:100%; margin-top:4px">
      <thead><tr><th>订单号</th><th>金额</th><th>状态</th><th>操作</th></tr></thead>
      <tbody>
        <tr v-for="o in list" :key="o.order_no">
          <td>{{ o.order_no }}</td>
          <td>{{ o.amount }}</td>
          <td>{{ o.status }}</td>
          <td><button class="btn" @click="showDetail(o.order_no)">详情</button></td>
        </tr>
      </tbody>
    </table>
  </div>
  <div v-if="detailOpen" class="modal" @click="closeDetail">
    <div class="modal-content" @click.stop>
      <h3>订单详情</h3>
      <pre style="margin-top:10px">{{ JSON.stringify(detail,null,2) }}</pre>
      <div style="display:flex;gap:8px;justify-content:flex-end"><button class="btn" @click="closeDetail">关闭</button></div>
    </div>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import api from '../utils/api.js'
const list = ref([])
async function load(){ const { data } = await api.get('/api/admin/orders'); list.value = data?.data || [] }
async function exportCsv(){ const { data } = await api.get('/api/admin/orders/export.csv'); const blob = new Blob([data]); const a = document.createElement('a'); a.href = URL.createObjectURL(blob); a.download = 'orders.csv'; a.click() }
async function exportJson(){ const { data } = await api.get('/api/admin/orders/export'); const blob = new Blob([JSON.stringify(data)], { type: 'application/json' }); const a = document.createElement('a'); a.href = URL.createObjectURL(blob); a.download = 'orders.json'; a.click() }
const detailOpen = ref(false)
const detail = ref(null)
async function showDetail(orderNo){ const { data } = await api.get(`/api/admin/orders/${orderNo}`); detail.value = (data?.data||data||[])[0] || {}; detailOpen.value = true }
function closeDetail(){ detailOpen.value = false; detail.value = null }
</script>
<style scoped>
.modal { position: fixed; inset: 0; background: rgba(0,0,0,0.45); display: flex; align-items: center; justify-content: center; z-index: 1000 }
.modal-content { background: #fff; border-radius: 12px; width: 640px; max-width: 90vw; padding: 16px; box-shadow: 0 10px 30px rgba(0,0,0,0.15) }
</style>