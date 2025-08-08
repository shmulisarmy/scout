import { handle_server_sync } from "../apiglue/sync";
export function ws(){
	//LINK /Users/shmuli/repositories/scout/main.go:88
	fetch(`http://localhost:5001/ws`, {credentials: 'include'})
	.then(response => {
	if (response.headers.get("sync")){
		handle_server_sync(JSON.parse(response.headers.get("sync")))
	}
	return response.json()})
	.then(data => console.log(data))
}
export function scouters(){
	//LINK /Users/shmuli/repositories/scout/main.go:90
	fetch(`http://localhost:5001/scouters`, {credentials: 'include'})
	.then(response => {
	if (response.headers.get("sync")){
		handle_server_sync(JSON.parse(response.headers.get("sync")))
	}
	return response.json()})
	.then(data => console.log(data))
}
