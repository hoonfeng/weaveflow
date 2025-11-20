<template>
  <div class="shell">
    <div class="top glass">
      <div class="brand" @click="$router.push('/')">IFaceConf Portal</div>
      <div class="nav">
        <a class="link" href="/" @click.prevent="$router.push('/')">首页</a>
        <a class="link" href="/docs" @click.prevent="$router.push('/docs')">文档</a>
        <a class="link" href="/openapi" @click.prevent="$router.push('/openapi')">OpenAPI</a>
        <a class="link" href="/metrics" @click.prevent="$router.push('/metrics')">指标</a>
        <a class="link" href="/console" @click.prevent="$router.push('/console')">控制台</a>
        <a class="link" href="/login" @click.prevent="$router.push('/login')">登录</a>
      </div>
    </div>
    <transition name="fade">
      <main class="content">
        <router-view />
      </main>
    </transition>
  </div>
</template>
<script setup>
</script>
<style>
.shell{min-height:100vh;display:flex;flex-direction:column}
.top{position:sticky;top:0;display:flex;align-items:center;justify-content:space-between;padding:12px 20px;margin:12px;border-radius:16px}
.brand{font-weight:700;letter-spacing:.8px;color:#fff}
.nav{display:flex;gap:14px}
.content{padding:20px}
</style>