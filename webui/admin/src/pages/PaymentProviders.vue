<template>
  <div class="card card-glass">
    <div class="section-header"><h3 class="section-title">支付通道配置</h3><div style="display:flex;gap:12px"><button class="btn" @click="load">刷新</button></div></div>
    <div style="display:flex;gap:12px;align-items:center;margin-bottom:8px">
      <input v-model="form.code" placeholder="代码(如 alipay/wechat)" class="glass" style="padding:10px" />
      <input v-model="form.name" placeholder="名称" class="glass" style="padding:10px" />
      <label style="display:flex;align-items:center;gap:6px"><input type="checkbox" v-model="form.enabled" /> 启用</label>
      <textarea v-model="form.config_json" placeholder="JSON 配置" class="glass" style="padding:10px;flex:1;height:80px"></textarea>
      <button class="btn" @click="create">创建</button>
    </div>
    <table class="table-modern" cellspacing="0" cellpadding="6" style="width:100%">
      <thead><tr><th>ID</th><th>代码</th><th>名称</th><th>启用</th><th>配置</th><th>操作</th></tr></thead>
      <tbody>
        <tr v-for="p in list" :key="p.id">
          <td>{{ p.id }}</td>
          <td>{{ p.code }}</td>
          <td><input v-model="p.name" class="glass" style="padding:8px" /></td>
          <td>
            <label style="display:flex;align-items:center;gap:6px"><input type="checkbox" :checked="!!p.enabled" @change="toggle(p, $event.target.checked)" /> 启用</label>
          </td>
          <td><textarea v-model="p.config_json" class="glass" style="padding:8px;width:100%;height:90px"></textarea></td>
          <td style="display:flex;gap:8px">
            <button class="btn" @click="update(p)">保存</button>
            <button class="btn" @click="remove(p)">删除</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
<script setup>
import { ref } from 'vue'
import { payProvidersApi } from '../utils/api.js'
const list = ref([])
const form = ref({ code:'', name:'', enabled:true, config_json:'' })
async function load(){ const r = await payProvidersApi.list(); const d = r?.data ?? r ?? []; list.value = d }
async function create(){ if (!form.value.code || !form.value.name) return; await payProvidersApi.create(form.value); form.value = { code:'', name:'', enabled:true, config_json:'' }; await load() }
async function update(p){ await payProvidersApi.update(p.id, { name: p.name, enabled: !!p.enabled, config_json: p.config_json }); await load() }
async function toggle(p, enabled){ await payProvidersApi.toggle(p.id, !!enabled); await load() }
async function remove(p){ if (!confirm('确定删除该通道？')) return; await payProvidersApi.remove(p.id); await load() }
load()
</script>
<style scoped>
</style>