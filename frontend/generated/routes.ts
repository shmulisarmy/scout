import { handle_server_sync } from "../apiglue/zustand_sync";
export function api_get_todos(){
	//LINK /Users/shmuli/repositories/scout/main.go:111
	fetch(`http://localhost:8080/api/get_todos`, {credentials: 'include'})
	.then(response => {
	if (response.headers.get("sync")){
		handle_server_sync(JSON.parse(response.headers.get("sync")))
	}
	return response.json()})
	.then(data => console.log(data))
}
export function api_add_todo(_0: string, ){
	//LINK /Users/shmuli/repositories/scout/main.go:118
	fetch(`http://localhost:8080/api/add_todo/${_0}`, {credentials: 'include'})
	.then(response => {
	if (response.headers.get("sync")){
		handle_server_sync(JSON.parse(response.headers.get("sync")))
	}
	return response.json()})
	.then(data => console.log(data))
}
export function api_ws(){
	//LINK /Users/shmuli/repositories/scout/main.go:126
	fetch(`http://localhost:8080/api/ws`, {credentials: 'include'})
	.then(response => {
	if (response.headers.get("sync")){
		handle_server_sync(JSON.parse(response.headers.get("sync")))
	}
	return response.json()})
	.then(data => console.log(data))
}
export function api_delete_todo(_0: number, ){
	//LINK /Users/shmuli/repositories/scout/main.go:128
	fetch(`http://localhost:8080/api/delete_todo/${_0}`, {credentials: 'include'})
	.then(response => {
	if (response.headers.get("sync")){
		handle_server_sync(JSON.parse(response.headers.get("sync")))
	}
	return response.json()})
	.then(data => console.log(data))
}
export function api_kanban_move_task(_0: number, _1: string, ){
	//LINK /Users/shmuli/repositories/scout/main.go:142
	fetch(`http://localhost:8080/api/kanban/move_task/${_0}/${_1}`, {credentials: 'include'})
	.then(response => {
	if (response.headers.get("sync")){
		handle_server_sync(JSON.parse(response.headers.get("sync")))
	}
	return response.json()})
	.then(data => console.log(data))
}
export function api_kanban_create_task(_0: string, _1: string, _2: string, _3: string, ){
	//LINK /Users/shmuli/repositories/scout/main.go:143
	fetch(`http://localhost:8080/api/kanban/create_task/${_0}/${_1}/${_2}/${_3}`, {credentials: 'include'})
	.then(response => {
	if (response.headers.get("sync")){
		handle_server_sync(JSON.parse(response.headers.get("sync")))
	}
	return response.json()})
	.then(data => console.log(data))
}
export function api_kanban_get_board(){
	//LINK /Users/shmuli/repositories/scout/main.go:144
	fetch(`http://localhost:8080/api/kanban/get_board`, {credentials: 'include'})
	.then(response => {
	if (response.headers.get("sync")){
		handle_server_sync(JSON.parse(response.headers.get("sync")))
	}
	return response.json()})
	.then(data => console.log(data))
}
