<script setup lang="ts">
import { ref } from "vue";
import SearchBar from "./components/SearchBar.vue";
import type { Email } from "./interfaces/Email";
import MailList from "./components/MailList.vue";
import { MailOpen } from "lucide-vue-next";

const mails = ref<Email[]>([]);

const handleSearch = async (searchTerm: string) => {
  const response = await fetch("http://localhost:8080/search", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ searchTerm }),
  });

  const { emails } = await response.json();
  mails.value = emails;
};
</script>

<template>
  <header class="flex gap-4 py-2">
    <MailOpen class="text-3xl self-center" />
    <h1 class="text-3xl">Zin Mail</h1>
  </header>
  <SearchBar @search="handleSearch" />
  <MailList :mails="mails" />
</template>

<style>
body {
  overflow-y: hidden;
}
</style>
