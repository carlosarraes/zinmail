<script setup lang="ts">
import { ChevronLeft, ChevronRight } from "lucide-vue-next";

defineProps<{
  currentPage: number;
  totalPages: number;
  searched: boolean;
  nextPage: () => void;
  prevPage: () => void;
}>();
</script>

<template>
  <div class="w-1/3 flex justify-center items-center">
    <ChevronLeft
      class="text-xl cursor-pointer self-center"
      :class="{
        'text-gray-300': currentPage <= 1,
        'text-white': currentPage > 1,
      }"
      @click="prevPage"
      :style="{ opacity: currentPage > 1 ? 1 : 0.5 }"
      v-show="searched"
    />
    <span
      class="text-xl text-white cursor-pointer self-center"
      v-show="searched"
    >
      Page {{ currentPage }} of {{ totalPages }}
    </span>
    <ChevronRight
      class="text-xl cursor-pointer self-center"
      :class="{
        'text-gray-300': currentPage >= totalPages,
        'text-white': currentPage < totalPages,
      }"
      @click="nextPage"
      :style="{ opacity: currentPage < totalPages ? 1 : 0.5 }"
      v-show="searched"
    />
  </div>
</template>
