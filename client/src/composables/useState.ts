import { reactive } from 'vue';

const state = reactive({
    search: '',
    mailSelected: {
        message_id: '',
        from: '',
        to: '',
        content: '',
    },
});

export function useState() {

    function setSearch(value: string) {
        state.search = value;
    }

    function setMailSelected(value: any) {
        state.mailSelected = {...value}
    }


    return {
        setSearch,
        setMailSelected,
        state,
    };
}