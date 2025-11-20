<template>
  <span :class="classes">{{ text }}</span>
  
</template>

<script setup>
const props = defineProps({
  kind: { type: String, default: 'ok' },
  text: { type: String, default: '' }
})
const classes = computed(() => ['tag', props.kind === 'ok' ? 'tag-ok' : 'tag-off'])
</script>

<style scoped>
</style>