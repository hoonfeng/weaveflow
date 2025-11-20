<template>
  <section class="grid">
    <div class="card">
      <h2>接口配置化平台</h2>
      <p>以配置驱动 API，与插件、数据源和文档一体化的现代化框架</p>
      <div style="display:flex;gap:12px;margin-top:8px">
        <button class="btn" @click="$router.push('/docs')">查看文档</button>
        <button class="btn" @click="$router.push('/openapi')">OpenAPI</button>
      </div>
    </div>
    <div class="card"><h3>控制台</h3><p>管理用量、APIKeys、Webhooks。</p><a class="link" href="/console" @click.prevent="$router.push('/console')">进入</a></div>
    <div class="card"><h3>指标</h3><p>查看请求指标，支持 Prometheus。</p><a class="link" href="/metrics" @click.prevent="$router.push('/metrics')">进入</a></div>
    <div class="card"><h3>登录</h3><p>获取 JWT，用于访问控制台。</p><a class="link" href="/login" @click.prevent="$router.push('/login')">进入</a></div>
  </section>
</template>
<script setup>
</script>