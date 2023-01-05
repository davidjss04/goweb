import { reactive, readonly } from 'vue';

const state = reactive({
    search: '',
    mailSelected: {
        message_id: '',
        from: '',
        to: '',
        content: '',
    },
    mails: [] as any,
});

export function useState() {
    async function loadData (mail: string)  {
        if (mail === '') {
            state.mails = [];
            return;
        }

        const res = await fetch(`http://localhost:3000/mails/emails/${mail}`);
        const data = await res.json();

        if (data === null) {
            state.mails = [];
            return;
        }

        state.mails = data.map((mail: any) => {
            return {
                message_id: mail.message_id,
                from: mail.from,
                to: mail.to.split(',')[0],
                content: mail.content,
            };
        });

        console.log(state.mails);
    };

    function setSearch(value: string) {
        loadData(value);
        state.search = value;
    }

    function setMailSelected(value: any) {
        state.mailSelected = { ...value };
    }

    function setMails(value: any) {
        state.mails = [...value];
    }

    return {
        setSearch,
        setMailSelected,
        setMails,
        state: readonly(state),
    };
}
