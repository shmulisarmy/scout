import { writable } from "svelte/store";

const ws = new WebSocket('ws://localhost:5001/ws');

ws.onopen = function (event) {
    console.log('WebSocket connection opened.');
};


export let notificationMessages = writable<{messages: string[], last_show_time: number, viewing_as_last_index: number}>({messages: [], last_show_time: 0, viewing_as_last_index: 0})

ws.onmessage = function (event) {
    console.log('Received:', event.data);
    const j = JSON.parse(event.data);
    if (j.type.startsWith("mutable")) {
        handle_server_sync(j);
    } else if (j.type === "notification") {
        notificationMessages.update((draft) => {
            draft.messages.push(j.message);
            // draft.last_show_time = Date.now();
            draft.viewing_as_last_index = draft.messages.length-1;
            return draft;
        });
    }
};

// --- Deep update helper (already built into updateWith via immer)
export function handle_server_sync(j: any) {
console.log('in handle_server_sync', j);

switch (j.type) {
    case 'populate-slot': {
        document.getElementById('slot')!.innerHTML = j.html;
        break;
    }
    case 'mutable-state-sender': {
        if (typeof window !== 'undefined') {
            window[j.key].update(() => j.data);
        }

        break;
    }
    case 'mutable-append':
        console.log('handling mutable-append', j);
        if (typeof window !== 'undefined') {
            window[j.key].updateWith((draft) => {
                const split_path = j.path.split('.');
                const last = split_path.pop();
                let target = draft as any;

                for (const part of split_path) {
                    if (!(part in target)) target[part] = {};
                    target = target[part];
                }

                target[last].push(j.new_data);
            });
        break;
    }

    case 'mutable-update': {
        console.log('handling mutable-update', j);
        if (typeof window !== 'undefined') {
            window[j.key].updateWith((draft) => {
                const split_path = j.path.split('.');
                const last = split_path.pop();
                let target = draft as any;

                for (const part of split_path) {
                if (!(part in target)) target[part] = {};
                target = target[part];
            }

            target[last!] = j.new_data;
        });
        break;
    }
}
}
}