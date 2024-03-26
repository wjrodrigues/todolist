<template>
  <p :class="props.class">{{ props.title }}</p>
</template>

<script setup lang="ts">
const props = defineProps<{
  title: string
  class: string
}>()
</script>
