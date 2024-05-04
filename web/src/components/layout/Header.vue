<template>
  <v-app-bar color="#6A1B9A" density="compact" :elevation="0">
    <v-app-bar-title>Todolist</v-app-bar-title>

    <template v-slot:append>
      <v-btn icon="mdi-magnify"></v-btn>
      <v-btn icon="mdi-dots-vertical"></v-btn>
    </template>
  </v-app-bar>
</template>

<script setup lang="ts">
</script>
