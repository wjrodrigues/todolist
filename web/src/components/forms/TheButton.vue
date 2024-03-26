<template>
  <button class="text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline hover:opacity-90"
    :class="props.color" :type="props.type" :disabled="props.disable">
    {{props.label}}
  </button>
</template>

<script setup lang="ts">
const props = defineProps<{
  type: "button" | "submit" | "reset"
  label: string
  color: string
  disable: boolean
}>()
</script>
