<template>
  <label :for=props.label class="block mb-2 text-sm font-medium text-amber-500">{{ props.label }}</label>
  <input :type=props.type :id=props.id
    class="shadow-sm border border-amber-500 text-gray-900 text-sm rounded-lg block w-full p-2.5 focus:border-blue-500"
    :placeholder=props.placeholder :required="props.required" />
</template>

<script setup lang="ts">
const props = defineProps<{
  id: string
  name: string
  type: string
  label: string
  placeholder: string
  required: boolean
}>()
</script>
