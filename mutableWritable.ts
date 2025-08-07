import { produce } from "immer";
import { writable } from "svelte/store";

export function mutableWritable<T>(initialValue: T) {
    const store = writable(initialValue);

    function updateWith(fn: (draft: T) => void) {
        store.update(produce(fn));
    }

    return {
        ...store,
        updateWith
    };
}