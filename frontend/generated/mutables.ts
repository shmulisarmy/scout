import { create } from 'zustand'

export const useLive_ScoutsStore = create<{state: Live_Scout[]}>((set) => ({
	state: [{"links":{},"notes":"","as_of":"","scouted":false,"to_scout_for":"tell me when there is an ai truck driver in the state of florida"},{"links":{},"notes":"","as_of":"","scouted":false,"to_scout_for":"find me a good reliable car with a resonable at signing that i can rent for about 200-300 a month that has a decent insurence policy for 21-25 year olds."},{"links":{},"notes":"","as_of":"","scouted":false,"to_scout_for":"let me know if there are any water parks that open up in a place that's close a bar, this place could be anywhere that is in the vecinity of kosher food, anywhere in america"},{"links":{},"notes":"","as_of":"","scouted":false,"to_scout_for":"is there a new decent laptop that i can get for under 200$ with 16 gb of ram."}],
}))
if (typeof window !== 'undefined') {
	(window as any).Live_Scouts = useLive_ScoutsStore
}
		
		export type Live_Scout = {
  links: { [key: string]: any };
  notes: string;
  as_of: string;
  scouted: boolean;
  to_scout_for: string;
}


export const useLive_ScoutStore = create<{state: Live_Scout}>((set) => ({
	state: {"links":{},"notes":"","as_of":"","scouted":false,"to_scout_for":"tell me when there is an ai truck driver in the state of florida"},
}))
if (typeof window !== 'undefined') {
	(window as any).Live_Scout = useLive_ScoutStore
}
		
		
export const usePersonsStore = create<{state: Person[]}>((set) => ({
	state: [{"name":"shmuli","age":21,"email":"shmuli@shmuli.com"},{"name":"berel","age":25,"email":"berel@shmuli.com"}],
}))
if (typeof window !== 'undefined') {
	(window as any).Persons = usePersonsStore
}
		
		