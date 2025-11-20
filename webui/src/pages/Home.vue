<template>
  <div>
    <section class="hero glass" :class="{ show }">
      <h2>接口配置化平台</h2>
      <p>以配置驱动 API，与插件、数据源和文档一体化的现代化框架</p>
      <div class="actions">
        <button class="btn" @click="$router.push('/docs')">查看文档</button>
        <button class="btn" @click="$router.push('/openapi')">OpenAPI</button>
      </div>
    </section>
    <section class="grid" style="margin-top:20px">
      <div class="card">
        <h3>接口文档</h3>
        <p>自动生成接口清单，快速联调。</p>
        <a class="link" href="/docs" @click.prevent="$router.push('/docs')">进入</a>
      </div>
      <div class="card">
        <h3>指标导出</h3>
        <p>内置请求指标，支持 Prometheus。</p>
        <a class="link" href="/metrics" @click.prevent="$router.push('/metrics')">进入</a>
      </div>
      <div class="card">
        <h3>登录</h3>
        <p>获取 JWT，访问受保护的管理接口。</p>
        <a class="link" href="/login" @click.prevent="$router.push('/login')">进入</a>
      </div>
      <div class="card">
        <h3>插件示例</h3>
        <p>调用 goproc 插件，体验扩展能力。</p>
        <a class="link" href="/reverse" @click.prevent="$router.push('/reverse')">进入</a>
      </div>
      <div class="card">
        <h3>文件上传</h3>
        <p>直传到本地存储，返回文件 ID。</p>
        <a class="link" href="/upload" @click.prevent="$router.push('/upload')">进入</a>
      </div>
      <div class="card">
        <h3>后台管理</h3>
        <p>用户、角色、租户、订单、支付等。</p>
        <a class="link" href="/admin" @click.prevent="$router.push('/admin')">进入</a>
      </div>
    </section>
  </div>
  </template>
<script setup>
import { ref, onMounted } from 'vue'
const show = ref(false)
onMounted(()=>{ setTimeout(()=> show.value = true, 180) })
</script>
<style>
.hero{padding:24px;margin:12px;border-radius:20px;transform:translateY(8px);opacity:0;transition:all .45s ease}
.hero.show{transform:none;opacity:1}
.hero h2{margin:0 0 8px}
.hero .actions{display:flex;gap:12px;margin-top:8px}
</style>