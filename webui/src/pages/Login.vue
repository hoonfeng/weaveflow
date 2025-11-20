<template>
  <div class="grid">
    <div class="card" style="max-width:420px">
      <h3>登录</h3>
      <form @submit.prevent="submit" style="display:grid;gap:12px;margin-top:8px">
        <input v-model="user" placeholder="用户名" class="glass" style="padding:10px" />
        <input type="password" v-model="pass" placeholder="密码" class="glass" style="padding:10px" />
        <button class="btn" :disabled="loading">提交</button>
      </form>
      <p v-if="token" style="margin-top:8px;word-break:break-all">Token: {{ token }}</p>
    </div>
    <div class="card">
      <h3>说明</h3>
      <p>登录后将自动在请求头附加 JWT 用于访问受保护的管理接口。</p>
    </div>
  </div>
  </template>
<script setup>
import { ref } from 'vue'
import api from '../api'
const user = ref('admin')
const pass = ref('admin')
const loading = ref(false)
const token = ref('')
async function submit(){
  loading.value = true
  try{
    const { data } = await api.post('/api/auth/login', { username: user.value, password: pass.value })
    const t = data?.data || data?.token || ''
    token.value = t
    if (t) localStorage.setItem('jwt', t)
  } finally {
    loading.value = false
  }
}
</script>