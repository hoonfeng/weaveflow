<template>
  <div class="grid">
    <div class="card">
      <div class="section-header"><div class="section-title">接口规范检查</div><button class="btn" @click="runLint">检查</button></div>
      <table class="table-modern" v-if="lintIssues.length">
        <thead><tr><th>级别</th><th>描述</th><th>位置</th></tr></thead>
        <tbody>
          <tr v-for="(it,idx) in lintIssues" :key="idx">
            <td><span class="badge-modern">{{ it.level }}</span></td>
            <td>{{ it.message }}</td>
            <td>{{ it.path }}</td>
          </tr>
        </tbody>
      </table>
      <div v-else>暂无问题</div>
    </div>

    <div class="card">
      <div class="section-header"><div class="section-title">重载配置</div><button class="btn" @click="doReload">重载</button></div>
      <pre style="margin-top:10px">{{ json(reloadDiff) }}</pre>
    </div>

    <div class="card">
      <div class="section-header"><div class="section-title">模型迁移</div><button class="btn" @click="doMigrate">迁移</button></div>
      <ul>
        <li v-for="(t,idx) in migrateTables" :key="idx">{{ t }}</li>
      </ul>
    </div>

    <div class="card">
      <div class="section-header"><div class="section-title">队列状态</div><button class="btn" @click="queue">查看</button></div>
      <pre style="margin-top:10px">{{ json(queueStatus) }}</pre>
    </div>

    <div class="card">
      <div class="section-header"><div class="section-title">清理 Nonces</div>
        <div style="display:flex;gap:8px;align-items:center">
          <input type="number" v-model.number="ttlSeconds" min="60" step="60" style="padding:8px;border-radius:8px;border:1px solid #cbd5e1;width:120px" />
          <button class="btn" @click="purge">清理</button>
        </div>
      </div>
      <pre style="margin-top:10px">{{ json(purgeResult) }}</pre>
    </div>

    <div class="card">
      <div class="section-header"><div class="section-title">远程调试</div><button class="btn" @click="remoteDebug">查看</button></div>
      <pre style="margin-top:10px">{{ json(remoteDebugState) }}</pre>
    </div>

    <div class="card">
      <div class="section-header"><div class="section-title">远程调用测试</div><button class="btn" @click="remoteTest">触发</button></div>
      <pre style="margin-top:10px">{{ json(remoteTestResult) }}</pre>
    </div>
  </div>
  
  
</template>
<script setup>
import { ref } from 'vue'
import { systemApi } from '../utils/api.js'

const lintIssues = ref([])
const reloadDiff = ref({})
const migrateTables = ref([])
const queueStatus = ref({})
const ttlSeconds = ref(900)
const purgeResult = ref({})
const remoteDebugState = ref({})
const remoteTestResult = ref({})

const json = (v) => JSON.stringify(v, null, 2)

async function runLint(){ const r = await systemApi.lint(); lintIssues.value = (r?.data) || r || [] }
async function doReload(){ const r = await systemApi.reload(); reloadDiff.value = r?.data || r || {} }
async function doMigrate(){ const r = await systemApi.migrate(); migrateTables.value = r?.data || r || [] }
async function queue(){ const r = await systemApi.queueStatus(); queueStatus.value = r?.data || r || {} }
async function purge(){ const r = await systemApi.purgeNonces(ttlSeconds.value); purgeResult.value = r?.data || r || {} }
async function remoteDebug(){ const r = await systemApi.remoteDebug(); remoteDebugState.value = r?.data || r || {} }
async function remoteTest(){ const r = await systemApi.remoteTest(); remoteTestResult.value = r?.data || r || {} }
</script>