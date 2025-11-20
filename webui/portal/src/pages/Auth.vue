<template>
  <div class="card" style="max-width:420px">
    <h3>登录</h3>
    <form @submit.prevent="submit" style="display:grid;gap:10px;margin-top:8px">
      <input v-model="user" placeholder="用户名" class="glass" style="padding:10px" />
      <input type="password" v-model="pass" placeholder="密码" class="glass" style="padding:10px" />
      <button class="btn" :disabled="loading">提交</button>
    </form>
    <p v-if="token" style="margin-top:8px;word-break:break-all">Token: {{ token }}</p>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import api from '../../../packages/api'
const user = ref('admin'); const pass = ref('admin'); const loading = ref(false); const token = ref('')
async function submit(){ loading.value=true; try{ const { data } = await api.post('/api/auth/login', { username: user.value, password: pass.value }); const t = data?.data?.token || data?.token || ''; token.value = t; if(t){ localStorage.setItem('jwt', t); location.href = '/console' } } finally{ loading.value=false } }
</script>