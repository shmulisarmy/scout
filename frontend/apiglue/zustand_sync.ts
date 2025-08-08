import { create } from "zustand";

// --- WebSocket setup
const ws = new WebSocket('ws://localhost:5001/ws');

ws.onopen = function () {
    console.log('WebSocket connection opened.');
};

// --- Zustand store (no middleware)
export const notificationMessages = create<{
    messages: string[];
    last_show_time: number;
    viewing_as_last_index: number;
    addMessage: (message: string) => void;
}>((set, get) => ({
    messages: [],
    last_show_time: 0,
    viewing_as_last_index: 0,

    addMessage: (message: string) => {
        const { messages } = get();
        const newMessages = [...messages, message];
        set({
            messages: newMessages,
            viewing_as_last_index: newMessages.length - 1,
            // last_show_time: Date.now() // optionally include this
        });
    },
}));

// --- WebSocket message handler
ws.onmessage = function (event) {
    console.log('Received:', event.data);
    const j = JSON.parse(event.data);

    if (j.type.startsWith("mutable")) {
        handle_server_sync(j);
    } else if (j.type === "notification") {
        notificationMessages.getState().addMessage(j.message);
    }
};

// --- Handle mutable updates from server
export function handle_server_sync(j: any) {
    console.log('in handle_server_sync', j);

    switch (j.type) {
        case 'populate-slot': {
            document.getElementById('slot')!.innerHTML = j.html;
            break;
        }

        case 'mutable-state-sender': {
            if (typeof window !== 'undefined') {
                window[j.key].setState((prev: any) => {return {...prev, state: j.data}});
            }
            break;
        }

        case 'mutable-append': {
            console.log('handling mutable-append', j);
            if (typeof window !== 'undefined') {
                window[j.key].setState((prev: any) => {
                    if (j.path === '') {
                        return {state: [...prev.state, j.new_data]}
                    }
                    const split_path = j.path.split('.');
                    const last = split_path.pop();
                    let target = { ...prev } as any;
                    let cursor = target.state;

                    for (const part of split_path) {
                        if (!(part in cursor)) cursor[part] = {};
                        else cursor[part] = { ...cursor[part] };
                        cursor = cursor[part];
                    }

                    const existing = cursor[last] ?? [];
                    cursor[last] = [...existing, j.new_data];

                    return target;
                });
            }
            break;
        }

        case 'mutable-update': {
            console.log('handling mutable-update', j);
            if (typeof window !== 'undefined') {
                window[j.key].setState((prev: any) => {
                    if (j.path === '') {
                        return {state: j.new_data}
                    }
                    const split_path = j.path.split('.');
                    const last = split_path.pop();
                    let target = { ...prev } as any;
                    let cursor = target.state;

                    for (const part of split_path) {
                        if (!(part in cursor)) cursor[part] = {};
                        else cursor[part] = { ...cursor[part] };
                        cursor = cursor[part];
                    }

                    cursor[last!] = j.new_data;

                    return target;
                });
            }
            break;
        }
    }
}
