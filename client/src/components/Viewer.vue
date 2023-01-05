<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { useState } from '../composables/useState';

const { state } = useState();
</script>
<template lang="">
    <div class="container mx-auto items-center bg-clip-padding">
        <div class="overflow-y-scroll p-10" style="height: 85vh;">
            {{ state.mailSelected.content || 'No mail selected' }}
        </div>
    </div>
</template>
