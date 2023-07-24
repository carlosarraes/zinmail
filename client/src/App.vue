<script setup lang="ts">
import { ref, computed } from "vue";
import SearchBar from "./components/SearchBar.vue";
import HandlePages from "./components/HandlePages.vue";
import MailList from "./components/MailList.vue";
import Header from "./components/Header.vue";
import Spinner from "./components/Spinner.vue";
import type { Email } from "./interfaces/Email";

const mails = ref<Email[]>([]);
const pageSize = ref(20);
const currentPage = ref(1);
const searched = ref(false);
const loading = ref(false);
const url = "http://127.0.0.1:8080";

let controller: AbortController | null = null;

const handleSearch = async (searchTerm: string) => {
  if (controller) controller.abort();

  loading.value = true;

  try {
    controller = new AbortController();

    const response = await fetch(`${url}/search`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ searchTerm }),
      signal: controller.signal,
    });

    if (!response.ok) throw new Error("Something went wrong");

    const { emails } = await response.json();
    mails.value = emails;
    searched.value = true;
    currentPage.value = 1;
  } catch (e) {
    console.error(e);
    if (e instanceof Error) {
      if (e.name === "AbortError") return;
    }
  } finally {
    loading.value = false;
    controller = null;
  }
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
  <Header />
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
  <div class="flex flex-col justify-center items-center">
    <MailList :mails="paginatedMails" />
    <Spinner v-show="loading" class="mt-4" />
  </div>
</template>

<style>
body {
  overflow-y: hidden;
}
</style>
