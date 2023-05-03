<script setup lang="ts">
import { ref } from "vue";
import { ChevronLeft, ChevronRight } from "lucide-vue-next";

const index = ref(0);

defineProps<{
  highlight: Array<string>;
}>();
</script>

<template>
  <section class="flex justify-center items-center gap-2">
    <ChevronLeft
      class="text-2xl text-black cursor-pointer self-center"
      @click="index > 0 ? index-- : index"
      v-show="highlight.length >= 2"
      :style="{ opacity: index > 0 ? 1 : 0.2 }"
    />
    <div
      class="text-sm text-left text-gray-900 w-96"
      v-html="highlight[index]"
    ></div>
    <ChevronRight
      class="text-2xl text-black cursor-pointer self-center"
      @click="index < highlight.length - 1 ? index++ : index"
      v-show="highlight.length >= 2"
      :style="{ opacity: index < highlight.length - 1 ? 1 : 0.2 }"
    />
  </section>
</template>
