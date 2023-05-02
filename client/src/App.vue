<script setup lang="ts">
import { ref } from "vue";
import SearchBar from "./components/SearchBar.vue";
import type { Email } from "./interfaces/Email";
import MailList from "./components/MailList.vue";

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
  <h1 class="header">Hello</h1>
  <SearchBar @search="handleSearch" />
  <MailList :mails="mails" />
</template>

<style>
body {
  overflow-y: hidden;
}
</style>
