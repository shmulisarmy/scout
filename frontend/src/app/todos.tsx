"use client"

import { useEffect } from "react";
import { handle_server_sync } from "../../apiglue/zustand_sync";
import { useTodosStore } from "../../generated/mutables";
import { add_todo, delete_todo, get_todos } from "../../generated/routes";




export default function Todos(){
    const {state: todos} = useTodosStore();
    useEffect(() => {
        get_todos()
    }, [])
    return (
        <div>
            <h1>Todos</h1>
            <ul>
                {todos.map((todo, index) => (
                    <li key={index}>
                        <p>
                            {todo.title}
                        </p>
                        <button onClick={() => {
                            delete_todo(todo.id)
                        }}>delete</button>
                    </li>
                ))}
            </ul>
            <form onSubmit={(e) => {
                e.preventDefault();
                add_todo((e.target as HTMLFormElement).todo.value)
            }}>
                <input type="text" name="todo" />
                <button type="submit">add todo</button>
            </form>
            <button onClick={() => {
                get_todos()
            }}>refresh todos</button>
        </div>
    )
}