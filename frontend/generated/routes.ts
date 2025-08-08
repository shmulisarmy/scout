import { handle_server_sync } from "../apiglue/zustand_sync";
export function get_todos(){
	//LINK /Users/shmuli/repositories/scout/main.go:110
	fetch(`http://localhost:8080/get_todos`, {credentials: 'include'})
	.then(response => {
	if (response.headers.get("sync")){
		handle_server_sync(JSON.parse(response.headers.get("sync")))
	}
	return response.json()})
	.then(data => console.log(data))
}
export function add_todo(_0: string, ){
	//LINK /Users/shmuli/repositories/scout/main.go:122
	fetch(`http://localhost:8080/add_todo/${_0}`, {credentials: 'include'})
	.then(response => {
	if (response.headers.get("sync")){
		handle_server_sync(JSON.parse(response.headers.get("sync")))
	}
	return response.json()})
	.then(data => console.log(data))
}
export function ws(){
	//LINK /Users/shmuli/repositories/scout/main.go:137
	fetch(`http://localhost:8080/ws`, {credentials: 'include'})
	.then(response => {
	if (response.headers.get("sync")){
		handle_server_sync(JSON.parse(response.headers.get("sync")))
	}
	return response.json()})
	.then(data => console.log(data))
}
export function delete_todo(_0: number, ){
	//LINK /Users/shmuli/repositories/scout/main.go:139
	fetch(`http://localhost:8080/delete_todo/${_0}`, {credentials: 'include'})
	.then(response => {
	if (response.headers.get("sync")){
		handle_server_sync(JSON.parse(response.headers.get("sync")))
	}
	return response.json()})
	.then(data => console.log(data))
}
