<template>
  <div class="card card-glass">
    <div class="section-header"><h3 class="section-title">支付管理</h3><div style="display:flex;gap:12px"><button class="btn" @click="load">刷新</button><button class="btn" @click="exportCsv">导出 CSV</button></div></div>
    <div style="display:flex;gap:12px;align-items:center;margin-bottom:8px">
      <input v-model="provider" placeholder="渠道(如 alipay/wechat)" class="glass" style="padding:10px" />
      <select v-model="status" class="glass" style="padding:10px">
        <option value="">全部状态</option>
        <option value="succeeded">succeeded</option>
        <option value="failed">failed</option>
        <option value="refunded">refunded</option>
      </select>
      <button class="btn" @click="load">筛选</button>
    </div>
    <table class="table-modern" cellspacing="0" cellpadding="6" style="width:100%; margin-top:4px">
      <thead><tr><th>交易ID</th><th>订单号</th><th>渠道</th><th>状态</th><th>时间</th><th>操作</th></tr></thead>
      <tbody>
        <tr v-for="p in list" :key="p.txn_id">
          <td>{{ p.txn_id }}</td>
          <td>{{ p.order_no }}</td>
          <td>{{ p.provider }}</td>
          <td>{{ p.status }}</td>
          <td>{{ p.ts }}</td>
          <td style="display:flex;gap:8px">
            <button class="btn" @click="showDetail(p.txn_id)">详情</button>
            <button class="btn" @click="openRefund(p)" :disabled="p.status==='refunded'">退款</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
  <div v-if="detailOpen" class="modal" @click="closeDetail">
    <div class="modal-content" @click.stop>
      <h3>支付详情</h3>
      <pre style="margin-top:10px">{{ JSON.stringify(detail,null,2) }}</pre>
      <div style="display:flex;gap:8px;justify-content:flex-end"><button class="btn" @click="closeDetail">关闭</button></div>
    </div>
  </div>
  <div v-if="refundOpen" class="modal" @click="closeRefund">
    <div class="modal-content" @click.stop>
      <h3>退款</h3>
      <form @submit.prevent="doRefund" style="display:flex;flex-direction:column;gap:8px">
        <input v-model.number="refundAmount" type="number" placeholder="金额" class="glass" style="padding:10px" />
        <input v-model="refundReason" placeholder="原因（可选）" class="glass" style="padding:10px" />
        <div style="display:flex;gap:8px;justify-content:flex-end"><button type="button" class="btn" @click="closeRefund">取消</button><button type="submit" class="btn">确定</button></div>
      </form>
    </div>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import api from '../utils/api.js'
import { paymentsApi } from '../utils/api.js'
const list = ref([])
const refundOpen = ref(false)
const currentPay = ref(null)
const refundAmount = ref(0)
const refundReason = ref('')
const detailOpen = ref(false)
const detail = ref(null)
const route = useRoute()
const router = useRouter()
const provider = ref(String(route.query.provider||''))
const status = ref(String(route.query.status||''))
async function load(){ const res = await paymentsApi.getPayments({ provider: provider.value||undefined, status: status.value||undefined }); const d = res?.data ?? res ?? {}; list.value = d.data || d || [] }
async function exportCsv(){ const res = await paymentsApi.exportPayments({ provider: provider.value||undefined, status: status.value||undefined }); const blob = new Blob([res]); const a = document.createElement('a'); a.href = URL.createObjectURL(blob); a.download = 'payments.csv'; a.click() }
async function showDetail(txnId){ const r = await paymentsApi.getPayment(txnId); const d = r?.data ?? r ?? {}; detail.value = (Array.isArray(d)? d[0] : d) || d; detailOpen.value = true }
function closeDetail(){ detailOpen.value = false; detail.value = null }
function openRefund(p){ currentPay.value = p; refundOpen.value = true; refundAmount.value = Number(p.amount)||0; refundReason.value = '' }
function closeRefund(){ refundOpen.value = false; currentPay.value = null; refundAmount.value = 0; refundReason.value = '' }
async function doRefund(){ if(!currentPay.value) return; await paymentsApi.refund(currentPay.value.txn_id, { amount: refundAmount.value, reason: refundReason.value }); closeRefund(); await load() }
load()
</script>
<style scoped>
.modal { position: fixed; inset: 0; background: rgba(0,0,0,0.45); display: flex; align-items: center; justify-content: center; z-index: 1000 }
.modal-content { background: #fff; border-radius: 12px; width: 520px; max-width: 90vw; padding: 16px; box-shadow: 0 10px 30px rgba(0,0,0,0.15) }
</style>