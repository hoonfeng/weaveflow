<template>
  <div class="card card-glass">
    <div class="section-header"><h3 class="section-title">权限管理</h3><div style="display:flex;gap:12px"><button class="btn" @click="load">刷新</button></div></div>
    <div style="display:flex;gap:12px;align-items:center;margin-bottom:8px">
      <input v-model="form.code" placeholder="代码(如 admin.users.list)" class="glass" style="padding:10px" />
      <input v-model="form.name" placeholder="名称" class="glass" style="padding:10px" />
      <input v-model="form.description" placeholder="描述" class="glass" style="padding:10px" />
      <button class="btn" @click="create">创建</button>
    </div>
    <table class="table-modern" cellspacing="0" cellpadding="6" style="width:100%">
      <thead><tr><th>ID</th><th>代码</th><th>名称</th><th>描述</th><th>操作</th></tr></thead>
      <tbody>
        <tr v-for="p in list" :key="p.id">
          <td>{{ p.id }}</td>
          <td><input v-model="p.code" class="glass" style="padding:8px" /></td>
          <td><input v-model="p.name" class="glass" style="padding:8px" /></td>
          <td><input v-model="p.description" class="glass" style="padding:8px" /></td>
          <td><button class="btn" @click="update(p)">更新</button><button class="btn" @click="del(p)">删除</button></td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import { rolesApi } from '../utils/api.js'
const list = ref([])
const form = ref({ code:'', name:'', description:'' })
async function load(){ const r = await rolesApi.getPermissions(); const d = r?.data ?? r ?? []; list.value = d }
async function create(){ if (!form.value.code) return; await rolesApi.createPermission({ code: form.value.code, name: form.value.name, description: form.value.description }); form.value = { code:'', name:'', description:'' }; await load() }
async function update(p){ await rolesApi.updatePermission(p.id, { code: p.code, name: p.name, description: p.description }); await load() }
async function del(p){ await rolesApi.deletePermission(p.id); await load() }
load()
</script>
<style scoped>
</style>