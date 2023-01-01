
<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useState } from '../composables/useState'

const { state } = useState()

</script>
<template lang="">
    <div class="container mx-auto items-center bg-clip-padding">
        {{state.mailSelected.content || 'No mail selected'}}
    </div>
</template>