<template>
  <div class="card card-glass">
    <div class="section-header"><h3 class="section-title">支付通道</h3><button class="btn" @click="load">刷新</button></div>
    <div class="grid">
      <div class="card" v-for="it in providers" :key="it.provider">
        <div class="section-header"><div class="section-title">{{ it.provider || '未知' }}</div><span class="badge-modern">{{ it.cnt }}</span></div>
        <div style="display:flex;gap:8px">
          <button class="btn" @click="viewPayments(it.provider)">查看交易</button>
        </div>
      </div>
    </div>
    <div v-if="!providers.length" style="padding:12px;color:#9fb0c9">暂无通道数据</div>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { overviewApi } from '../utils/api.js'
const router = useRouter()
const providers = ref([])
async function load(){ const r = await overviewApi.paymentsStats(); const d = r?.data ?? r ?? {}; providers.value = d.by_provider || [] }
function viewPayments(provider){ router.push({ path: '/payments', query: { provider } }) }
load()
</script>
<style scoped>
</style>