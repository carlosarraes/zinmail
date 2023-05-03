<script setup lang="ts">
import { MailOpen, Github, Code } from "lucide-vue-next";
</script>

<template>
  <header class="flex justify-between py-2">
    <div class="flex gap-4">
      <MailOpen class="text-3xl self-center" />
      <h1 class="text-3xl">Zin Mail</h1>
    </div>
    <div class="flex justify-center gap-4 items-center">
      <a
        href="https://github.com/carlosarraes/"
        target="_blank"
        rel="noopener noreferrer"
        class="flex gap-2 items-center hover:scale-105 duration-200"
      >
        <Github :size="32" />
        <span>/carlosarraes</span>
      </a>
      <a
        href="https://github.com/carlosarraes/zinmail"
        target="_blank"
        rel="noopener noreferrer"
        class="flex gap-2 items-center hover:scale-105 duration-200"
      >
        <Code :size="32" />
        <span>Source</span>
      </a>
    </div>
  </header>
</template>
