import { create } from 'zustand'


export const useCounterStore = create<{state: number}>((set) => {
  return {state: 0};
})


if (typeof window !== 'undefined') {
    (window as any).useCounterStore = useCounterStore;
}