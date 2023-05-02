<script setup lang="ts">
import { ref } from "vue";
const searchText = ref("");
const emit = defineEmits(["search"]);

const onSearch = () => {
  emit("search", searchText.value);
};
</script>

<template>
  <input
    type="text"
    class="text-black"
    v-model="searchText"
    @keyup.enter="onSearch"
  />
  <button @click="onSearch">Search</button>
</template>
