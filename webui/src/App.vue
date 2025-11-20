<template>
  <div class="shell">
    <div class="top glass" :class="{ show: show }">
      <div class="brand" @click="$router.push('/')">IFaceConf</div>
      <div class="nav">
        <a class="link" href="/admin" @click.prevent="$router.push('/admin')">后台</a>
        <a class="link" href="/admin/users" @click.prevent="$router.push('/admin/users')">用户</a>
        <a class="link" href="/admin/roles" @click.prevent="$router.push('/admin/roles')">角色</a>
        <a class="link" href="/admin/tenants" @click.prevent="$router.push('/admin/tenants')">租户</a>
        <a class="link" href="/admin/orders" @click.prevent="$router.push('/admin/orders')">订单</a>
        <a class="link" href="/admin/payments" @click.prevent="$router.push('/admin/payments')">支付</a>
        <a class="link" href="/admin/tools" @click.prevent="$router.push('/admin/tools')">运维</a>
        <a class="link" href="/console" @click.prevent="$router.push('/console')">控制台</a>
        <a class="link" href="/login" @click.prevent="$router.push('/login')">登录</a>
      </div>
      <button class="btn" @click="toggle">切换主题</button>
    </div>
    <transition name="fade">
      <main class="content">
        <router-view />
      </main>
    </transition>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
const show = ref(false)
function toggle(){ document.body.classList.toggle('dark') }
onMounted(()=>{ setTimeout(()=>{ show.value = true }, 120) })
</script>

<style>
.shell{min-height:100vh;display:flex;flex-direction:column;font-family:system-ui,Arial,Helvetica,sans-serif}
.top{position:sticky;top:0;display:flex;align-items:center;justify-content:space-between;padding:12px 20px;margin:12px;border-radius:16px;opacity:0;transform:translateY(-8px);transition:all .35s ease}
.top.show{opacity:1;transform:none}
.brand{font-weight:700;letter-spacing:.8px;color:#fff}
.nav{display:flex;gap:14px}
.content{padding:20px}
body.dark{--bg:#0f172a;--fg:#e5e7eb}
</style>