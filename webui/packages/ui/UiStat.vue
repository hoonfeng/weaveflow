<template>
  <div class="stat-card">
    <div class="stat-icon"><slot name="icon" /></div>
    <div class="stat-content">
      <div class="stat-value">{{ value }}</div>
      <div class="stat-label">{{ label }}</div>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  label: { type: String, default: '' },
  value: { type: [String, Number], default: '' }
})
</script>

<style scoped>
</style>