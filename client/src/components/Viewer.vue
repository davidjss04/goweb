
<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useState } from '../composables/useState'

const { state } = useState()

// const mail = ref({})

// const loadMail = async (id: string) => {

//     if (id === '') {
//         mail.value = []
//         return
//     }

//     const res = await fetch(`http://localhost:3000/mail/${id}`)
//     const data = await res.json()

//     if (data.length === 0) return

//     mail.value = data.map((mail: any) => {
//         return {
//             message_id: mail.message_id,
//             from: mail.from,
//             to: mail.to.split(',')[0],
//             content: mail.content
//         }
//     })


//     console.log(mail.value)
// }


// onMounted(() => {
//     loadMail(state.mailSelected.message_id)
// })

</script>
<template lang="">
    <div class="container mx-auto items-center bg-clip-padding">
        {{state.mailSelected.content || 'No mail selected'}}
    </div>
</template>