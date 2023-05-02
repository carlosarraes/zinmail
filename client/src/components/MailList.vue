<script setup lang="ts">
import type { Email } from "../interfaces/Email";
import { ref } from "vue";
const props = defineProps<{
  mails: Email[];
}>();

const showMail = ref(false);
const singleMailContent = ref({
  from: "",
  to: "",
  subject: "",
  content: "",
});

const toggleMail = (mail?: string) => {
  showMail.value = !showMail.value;
  const mailContent = props.mails.find((m) => m.id === mail);
  if (mailContent) {
    singleMailContent.value = mailContent;
  }
  console.log(singleMailContent.value);
};
</script>

<template>
  <div
    v-if="showMail"
    class="absolute w-full h-full bg-black bg-opacity-70 z-10 flex items-center justify-center"
  >
    <div class="w-8/12 bg-white rounded-md relative">
      <div class="absolute top-0 right-0 m-4">
        <button
          @click="toggleMail('close')"
          class="text-2xl text-black rounded-full bg-red-500 w-8 h-8"
        >
          &times;
        </button>
      </div>
      <div class="flex flex-col m-2 text-left p-4 overflow-y-auto max-h-96">
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
  <div class="w-full mx-auto table-wrapper">
    <table class="w-[800px] divide-y divide-gray-200">
      <thead class="bg-gray-50 min-w-full">
        <tr>
          <th
            scope="col"
            class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
          >
            From
          </th>
          <th
            scope="col"
            class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
          >
            To
          </th>
          <th
            scope="col"
            class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
          >
            Subject
          </th>
          <th
            scope="col"
            class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
          >
            Highlight
          </th>
        </tr>
      </thead>
      <tbody
        class="bg-slate-50 min-w-full divide-y divide-gray-200 cursor-pointer"
      >
        <tr v-for="mail in mails" :key="mail.id" @click="toggleMail(mail.id)">
          <td class="px-6 py-4 whitespace-nowrap">
            <div class="text-sm text-left text-gray-900 truncate w-52">
              {{ mail.from }}
            </div>
          </td>
          <td class="px-6 py-4">
            <div class="text-sm text-left text-gray-900 truncate w-52">
              {{ mail.to }}
            </div>
          </td>
          <td class="px-6 py-4 whitespace-nowrap">
            <div class="text-sm text-left text-gray-900 truncate w-52">
              {{ mail.subject }}
            </div>
          </td>
          <td class="px-6 py-4">
            <div
              class="text-sm text-left text-gray-900 w-96"
              v-html="mail.highlight[0]"
            ></div>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<style scoped>
.table-wrapper {
  max-height: calc(100vh - 4rem);
  overflow-y: scroll;
}
</style>
