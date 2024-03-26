<template>
  <main>
    <div class="h-screen flex items-center justify-center">
      <div class="w-full max-w-xs">
        <form class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-16">
          <div class="mb-4">
            <TheTitle title="Cadastro" class="text-xl font-medium text-center	" />
          </div>

          <div class="mb-4">
            <TheInput id="name" name="name" type="text" label="Nome" placeholder="" :required="true" />
          </div>

          <div class="mb-4">
            <TheInput id="email" name="email" type="email" label="Email" placeholder="user@email.com"
              :required="true" />
          </div>

          <div class="mb-4">
            <TheInput id="password" name="password" type="password" label="Senha" placeholder="*******"
              :required="true" />
          </div>

          <div class="mb-4">
            <div class="flex justify-end">
              <TheBotton type="submit" label="Salvar" color="bg-blue-600" :disable="true" />
            </div>
          </div>
        </form>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import TheTitle from "@/components/forms/TheTitle.vue"
import TheInput from "@/components/forms/TheInput.vue"
import TheBotton from "@/components/forms/TheButton.vue"
</script>