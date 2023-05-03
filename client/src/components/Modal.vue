<script setup lang="ts">
import type { SingleMail } from "../interfaces/Email";

defineProps<{
  showMail: boolean;
  singleMailContent: SingleMail;
  toggleMail: (mail?: string) => void;
}>();
</script>

<template>
  <div
    v-if="showMail"
    class="absolute w-full h-full bg-black bg-opacity-70 z-10 flex items-center justify-center"
    @click="toggleMail('close')"
  >
    <div class="w-[800px] h-[800px] bg-white rounded-md relative" @click.stop>
      <div class="absolute top-0 right-0 m-4">
        <button
          @click="toggleMail('close')"
          class="text-2xl text-black rounded-full bg-red-500 w-8 h-8"
        >
          &times;
        </button>
      </div>
      <div class="flex flex-col m-2 text-left p-4 overflow-y-auto h-full">
        <p class="text-black text-sm">From: {{ singleMailContent.from }}</p>
        <p class="text-black text-sm">
          To:
          <span v-if="singleMailContent.to.split(',').length > 2">
            {{ singleMailContent.to.split(",")[0].trim() }},
            {{ singleMailContent.to.split(",")[1].trim() }}
            <span class="text-gray-600">
              and {{ singleMailContent.to.split(",").length - 2 }} more...
            </span>
          </span>
          <span v-else>
            {{ singleMailContent.to }}
          </span>
        </p>
        <p class="text-black text-sm">Date: {{ singleMailContent.date }}</p>
        <p class="text-black text-sm">
          Subject: {{ singleMailContent.subject }}
        </p>
        <p class="text-black text-sm mt-4">Content:</p>
        <p class="text-black text-sm mt-2">
          {{ singleMailContent.content }}
        </p>
      </div>
    </div>
  </div>
</template>
