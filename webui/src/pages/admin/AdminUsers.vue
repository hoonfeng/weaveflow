<template>
  <div class="grid">
    <div class="card">
      <h3>创建用户</h3>
      <form @submit.prevent="create" style="display:grid;gap:10px;margin-top:8px">
        <input v-model="form.username" placeholder="用户名" class="glass" style="padding:10px" />
        <input v-model="form.email" placeholder="邮箱" class="glass" style="padding:10px" />
        <input v-model="form.password" type="password" placeholder="密码" class="glass" style="padding:10px" />
        <button class="btn" :disabled="loading">创建</button>
      </form>
    </div>
    <div class="card">
      <h3>用户列表</h3>
      <div style="overflow:auto;" v-if="list.length">
        <table cellspacing="0" cellpadding="6" style="width:100%">
          <thead><tr><th>ID</th><th>用户名</th><th>邮箱</th><th>操作</th></tr></thead>
          <tbody>
            <tr v-for="u in list" :key="u.id">
              <td>{{ u.id }}</td>
              <td><input v-model="u.username" class="glass" style="padding:8px" /></td>
              <td><input v-model="u.email" class="glass" style="padding:8px" /></td>
              <td>
                <button class="btn" @click="update(u)">更新</button>
                <button class="btn" @click="del(u)">删除</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <div v-else>无数据</div>
    </div>
  </div>
</template>
<script setup>
import { ref, onMounted } from 'vue'
import api from '../../api'
const loading = ref(false)
const form = ref({ username:'', email:'', password:'' })
const list = ref([])
async function load(){
  const { data } = await api.get('/api/admin/users')
  list.value = data?.data || []
}
onMounted(load)
async function create(){
  loading.value = true
  try{
    await api.post('/api/admin/users', form.value)
    form.value = { username:'', email:'', password:'' }
    await load()
  } finally { loading.value = false }
}
async function update(u){
  await api.put(`/api/admin/users/${u.id}`, { username: u.username, email: u.email })
  await load()
}
async function del(u){
  await api.delete(`/api/admin/users/${u.id}`)
  await load()
}
</script>