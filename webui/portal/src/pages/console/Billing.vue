<template>
  <div class="grid">
    <div class="card">
      <h3>订单</h3>
      <div style="display:flex;gap:12px">
        <button class="btn" @click="loadOrders">刷新</button>
        <button class="btn" @click="exportOrders">导出 CSV</button>
      </div>
      <ul style="margin-top:10px"><li v-for="o in orders" :key="o.order_no">{{ o.order_no }} - {{ o.amount }} - {{ o.status }}</li></ul>
    </div>
    <div class="card">
      <h3>支付</h3>
      <div style="display:flex;gap:12px">
        <button class="btn" @click="loadPayments">刷新</button>
        <button class="btn" @click="exportPayments">导出 CSV</button>
      </div>
      <ul style="margin-top:10px"><li v-for="p in payments" :key="p.txn_id">{{ p.txn_id }} - {{ p.order_no }} - {{ p.amount }} - {{ p.status }}</li></ul>
    </div>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import api from '../../../../packages/api'
const orders = ref([]); const payments = ref([])
async function loadOrders(){ const { data } = await api.get('/api/admin/orders'); orders.value = data?.data || [] }
async function exportOrders(){ const { data } = await api.get('/api/admin/orders/export.csv'); const blob = new Blob([data]); const a = document.createElement('a'); a.href = URL.createObjectURL(blob); a.download = 'orders.csv'; a.click() }
async function loadPayments(){ const { data } = await api.get('/api/admin/payments'); payments.value = data?.data || [] }
async function exportPayments(){ const { data } = await api.get('/api/admin/payments/export.csv'); const blob = new Blob([data]); const a = document.createElement('a'); a.href = URL.createObjectURL(blob); a.download = 'payments.csv'; a.click() }
</script>