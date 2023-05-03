<script setup lang="ts">
import type { Email, SingleMail } from "../interfaces/Email";
import Modal from "./Modal.vue";
import MatchedContent from "./MatchedContent.vue";
import { Eye } from "lucide-vue-next";
import { ref } from "vue";

const props = defineProps<{
  mails: Email[];
}>();

const showMail = ref(false);
const singleMailContent = ref<SingleMail>({
  from: "",
  to: "",
  date: "",
  subject: "",
  content: "",
});

const toggleMail = (mail?: string) => {
  showMail.value = !showMail.value;
  const mailContent = props.mails.find((m) => m.id === mail);
  if (mailContent) {
    singleMailContent.value = mailContent;
  }
};
</script>

<template>
  <div class="modal-wrapper" v-show="showMail">
    <Modal
      :showMail="showMail"
      :singleMailContent="singleMailContent"
      :toggleMail="toggleMail"
    />
  </div>
  <section class="w-full mx-auto table-wrapper">
    <table class="w-[1200px] divide-y divide-gray-200">
      <thead class="bg-gray-50 min-w-full sticky top-0 z-0 shadow-md">
        <tr>
          <th
            class="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider"
          >
            From
          </th>
          <th
            class="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider"
          >
            To
          </th>
          <th
            class="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider"
          >
            Subject
          </th>
          <th
            class="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider"
          >
            Matched Content
          </th>
          <th
            class="px-6 py-3 text-center text-xs font-medium text-gray-500 uppercase tracking-wider"
          >
            View
          </th>
        </tr>
      </thead>
      <tbody class="bg-slate-50 min-w-full divide-y divide-gray-200">
        <tr v-for="mail in mails" :key="mail.id">
          <td class="px-6 py-4">
            <div class="text-sm text-left text-gray-900 truncate w-40">
              {{ mail.from }}
            </div>
          </td>
          <td class="px-6 py-4">
            <div class="text-sm text-left text-gray-900 truncate w-40">
              {{ mail.to === "" ? "No recipient" : mail.to }}
            </div>
          </td>
          <td class="px-6 py-4 whitespace-nowrap">
            <div class="text-sm text-left text-gray-900 truncate w-44">
              {{ mail.subject }}
            </div>
          </td>
          <td class="px-6 py-4 cursor-text">
            <MatchedContent :highlight="mail.highlight" />
          </td>
          <td class="px-6 py-4">
            <div
              class="flex justify-center items-center text-sm text-left text-gray-900 w-8"
            >
              <Eye
                class="text-2xl text-black cursor-pointer self-center hover:scale-110 duration-150"
                @click="toggleMail(mail.id)"
              />
            </div>
          </td>
        </tr>
      </tbody>
    </table>
  </section>
</template>

<style scoped>
.table-wrapper {
  max-height: calc(100vh - 4rem);
  overflow-y: scroll;
  overflow-x: hidden;
  padding-bottom: 2rem;
}

.modal-wrapper {
  position: fixed;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  z-index: 1;
}
</style>
