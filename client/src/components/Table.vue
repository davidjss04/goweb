<script setup lang="ts">
import { onMounted, onUnmounted, onUpdated, ref, watchEffect } from 'vue';
import { useState } from '../composables/useState';

const { state, setMailSelected } = useState();

const selectMail = (mail: any) => {
    setMailSelected(mail);
};

onUpdated(() => {
    console.log('updated');
});

watchEffect(() => {
    console.log('watchEffect table');
    console.log('desde table', state.mails);
});

</script>

<template lang="">
    <div class="relative overflow-x-auto shadow-md sm:rounded-lg">
        <table
            class="w-full text-sm text-left text-gray-500 dark:text-gray-400"
        >
            <thead
                class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400"
            >
                <tr>
                    <th scope="col" class="px-6 py-3">#</th>
                    <th scope="col" class="px-6 py-3">Message ID</th>
                    <th scope="col" class="px-6 py-3">From</th>
                    <th scope="col" class="px-6 py-3">To</th>
                    <th scope="col" class="px-6 py-3">
                        <span class="sr-only">View</span>
                    </th>
                </tr>
            </thead>
            <tbody>
                <tr
                    class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600"
                    v-for="(mail, index) in state.mails"
                >
                    <td class="px-6 py-4">{{ index }}</td>
                    <th
                        scope="row"
                        class="px-6 py-4 font-medium text-gray-900 dark:text-white whitespace-nowrap"
                    >
                        {{ mail.message_id.split('.')[1] }}
                    </th>
                    <td class="px-6 py-4">{{ mail.from  }}</td>
                    <td class="px-6 py-4">{{ mail.to }}</td>
                    <td class="px-6 py-4 text-right">
                        <a
                            href="#"
                            class="font-medium text-blue-600 dark:text-blue-500 hover:underline"
                            @click="selectMail(mail)"
                            >View</a
                        >
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</template>
