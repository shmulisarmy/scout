import { handle_server_sync } from "../apiglue/zustand_sync";
export function api_get_todos(){
	//LINK /Users/shmuli/repositories/scout/main.go:111
	fetch(`http://localhost:8080/api/get_todos`)
	.then(response => {
	if (response.headers.get("sync")){
		handle_server_sync(JSON.parse(response.headers.get("sync")))
	}
	return response.json()})
	.then(data => console.log(data))
}
export function api_add_todo(_0: string, ){
	//LINK /Users/shmuli/repositories/scout/main.go:118
	fetch(`http://localhost:8080/api/add_todo/${_0}`)
	.then(response => {
	if (response.headers.get("sync")){
		handle_server_sync(JSON.parse(response.headers.get("sync")))
	}
	return response.json()})
	.then(data => console.log(data))
}
export function api_ws(){
	//LINK /Users/shmuli/repositories/scout/main.go:126
	fetch(`http://localhost:8080/api/ws`)
	.then(response => {
	if (response.headers.get("sync")){
		handle_server_sync(JSON.parse(response.headers.get("sync")))
	}
	return response.json()})
	.then(data => console.log(data))
}
export function api_delete_todo(_0: number, ){
	//LINK /Users/shmuli/repositories/scout/main.go:128
	fetch(`http://localhost:8080/api/delete_todo/${_0}`)
	.then(response => {
	if (response.headers.get("sync")){
		handle_server_sync(JSON.parse(response.headers.get("sync")))
	}
	return response.json()})
	.then(data => console.log(data))
}
export function api_kanban_move_task(_0: number, _1: string, ){
	//LINK /Users/shmuli/repositories/scout/main.go:142
	fetch(`http://localhost:8080/api/kanban/move_task/${_0}/${_1}`)
	.then(response => {
	if (response.headers.get("sync")){
		handle_server_sync(JSON.parse(response.headers.get("sync")))
	}
	return response.json()})
	.then(data => console.log(data))
}
export function api_kanban_create_task(_0: string, _1: string, _2: string, _3: string, ){
	//LINK /Users/shmuli/repositories/scout/main.go:143
	fetch(`http://localhost:8080/api/kanban/create_task/${_0}/${_1}/${_2}/${_3}`)
	.then(response => {
	if (response.headers.get("sync")){
		handle_server_sync(JSON.parse(response.headers.get("sync")))
	}
	return response.json()})
	.then(data => console.log(data))
}
export function api_kanban_get_board(){
	//LINK /Users/shmuli/repositories/scout/main.go:144
	fetch(`http://localhost:8080/api/kanban/get_board`)
	.then(response => {
	if (response.headers.get("sync")){
		handle_server_sync(JSON.parse(response.headers.get("sync")))
	}
	return response.json()})
	.then(data => console.log(data))
}
export function api_kanban_add_comment(_0: number, _1: string, _2: string, ){
	//LINK /Users/shmuli/repositories/scout/main.go:145
	fetch(`http://localhost:8080/api/kanban/add_comment/${_0}/${_1}/${_2}`)
	.then(response => {
	if (response.headers.get("sync")){
		handle_server_sync(JSON.parse(response.headers.get("sync")))
	}
	return response.json()})
	.then(data => console.log(data))
}
export function api_kanban_delete_task(_0: number, ){
	//LINK /Users/shmuli/repositories/scout/main.go:146
	fetch(`http://localhost:8080/api/kanban/delete_task/${_0}`)
	.then(response => {
	if (response.headers.get("sync")){
		handle_server_sync(JSON.parse(response.headers.get("sync")))
	}
	return response.json()})
	.then(data => console.log(data))
}
