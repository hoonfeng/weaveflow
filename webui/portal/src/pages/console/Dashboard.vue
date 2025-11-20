<template>
  <div class="grid">
    <div class="card"><h3>用量</h3><p>查看日/月用量与报表</p><a class="link" href="/console/usage" @click.prevent="$router.push('/console/usage')">进入</a></div>
    <div class="card"><h3>API Keys</h3><p>生成与轮换</p><a class="link" href="/console/apikeys" @click.prevent="$router.push('/console/apikeys')">进入</a></div>
    <div class="card"><h3>Webhooks</h3><p>管理端点与任务</p><a class="link" href="/console/webhooks" @click.prevent="$router.push('/console/webhooks')">进入</a></div>
    <div class="card"><h3>账单</h3><p>订单与支付记录</p><a class="link" href="/console/billing" @click.prevent="$router.push('/console/billing')">进入</a></div>
  </div>
</template>
<script setup>
</script>