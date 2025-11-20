<template>
  <div>
    <h2>后台管理</h2>
    <ul>
      <li><router-link to="/admin/users">用户管理</router-link></li>
      <li><router-link to="/admin/roles">角色管理</router-link></li>
      <li><router-link to="/admin/tenants">租户管理</router-link></li>
      <li><router-link to="/admin/plans">套餐计划</router-link></li>
      <li><router-link to="/admin/subscriptions">订阅管理</router-link></li>
      <li><router-link to="/admin/apikeys">API Key 管理</router-link></li>
      <li><router-link to="/admin/orders">订单管理</router-link></li>
      <li><router-link to="/admin/payments">支付管理</router-link></li>
      <li><router-link to="/admin/tools">工具与运维</router-link></li>
    </ul>
  </div>
</template>
<script setup>
</script>