<script setup lang="ts">
import { ref, computed } from "vue";
import SearchBar from "./components/SearchBar.vue";
import HandlePages from "./components/HandlePages.vue";
import MailList from "./components/MailList.vue";
import type { Email } from "./interfaces/Email";
import { MailOpen } from "lucide-vue-next";

const mails = ref<Email[]>([]);
const pageSize = ref(20);
const currentPage = ref(1);
const searched = ref(false);

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
  searched.value = true;
};

const totalPages = computed(() =>
  Math.ceil(mails.value.length / pageSize.value)
);

const paginatedMails = computed(() => {
  const start = (currentPage.value - 1) * pageSize.value;
  const end = start + pageSize.value;
  return mails.value.slice(start, end);
});

const nextPage = () => {
  if (currentPage.value < totalPages.value) currentPage.value += 1;
};

const prevPage = () => {
  if (currentPage.value > 1) currentPage.value -= 1;
};
</script>

<template>
  <header class="flex gap-4 py-2">
    <MailOpen class="text-3xl self-center" />
    <h1 class="text-3xl">Zin Mail</h1>
  </header>
  <section class="flex w-full items-center justify-evenly">
    <HandlePages
      :currentPage="currentPage"
      :totalPages="totalPages"
      :searched="searched"
      :nextPage="nextPage"
      :prevPage="prevPage"
    />
    <div class="w-2/3">
      <SearchBar @search="handleSearch" />
    </div>
  </section>
  <MailList :mails="paginatedMails" />
</template>

<style>
body {
  overflow-y: hidden;
}
</style>
