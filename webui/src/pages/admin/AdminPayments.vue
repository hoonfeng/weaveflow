<template>
  <div>
    <h3>支付管理</h3>
    <div>
      <button @click="load">刷新</button>
      <button @click="exportCsv">导出 CSV</button>
    </div>
    <table border="1" cellspacing="0" cellpadding="6" style="margin-top:12px">
      <thead><tr><th>交易ID</th><th>订单号</th><th>金额</th><th>状态</th></tr></thead>
      <tbody>
        <tr v-for="p in list" :key="p.txn_id">
          <td>{{ p.txn_id }}</td>
          <td>{{ p.order_no }}</td>
          <td>{{ p.amount }}</td>
          <td>{{ p.status }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import api from '../../api'
const list = ref([])
async function load(){ const { data } = await api.get('/api/admin/payments'); list.value = data?.data || [] }
async function exportCsv(){ const { data } = await api.get('/api/admin/payments/export.csv'); const blob = new Blob([data]); const a = document.createElement('a'); a.href = URL.createObjectURL(blob); a.download = 'payments.csv'; a.click() }
</script>