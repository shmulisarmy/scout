"use client"

import { useRef } from "react";
import { useCounterStore } from "./stores";
import { handle_server_sync } from "../../apiglue/zustand_sync";
import { usePersonsStore } from "../../generated/mutables";






export default function Counter(){
    const {state: count} = useCounterStore();
    const {state: people} = usePersonsStore();
    let form_ref = useRef<HTMLFormElement>(null);
    
    return (
        <div>
            <h1>Counter</h1>
            <p>{count}</p>
            <p>{count}</p>
            <p>{count}</p>
            <p>{count}</p>
            <p>{count}</p>
            <p>{count}</p>
            <button onClick={() => {
                handle_server_sync({type: 'mutable-update', path: '', new_data: count + 1, key: 'useCounterStore'})
            }}>new update</button>
            <h1>People</h1>
            <ul>
                {people.map((person, index) => (
                    <li key={index}>{person.name}</li>
                ))}
            </ul>
            <button onClick={() => {
                handle_server_sync({type: 'mutable-append', path: '', new_data: {name: 'new person', age: 21, email: 'newperson@newperson.com'}, key: 'Persons'})
            }}>add person</button>

        </div>
    )   
}