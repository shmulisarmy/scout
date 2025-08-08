package apiglue

import (
	"context"
	"fmt"
	"gin-sevalla-app/api_glue/utils"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var Config struct {
	Port       string
	Src_folder string
	Framework  string
}

// var html_base_page_string = utils.Load_from_file("templates/base_page.html")

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var ctx = context.Background()

func NewServerState[T any](state T) ServerState[T] {
	key := utils.CreateNameFromType(state)
	fmt.Println("key", key)
	return ServerState[T]{
		State: state,
		Key:   key,
	}
}

var to_add_to_mutable_ts_file_beginning = map[string]string{
	"zustand": "import { create } from 'zustand'\n",
	"svelte":  "import { mutableWritable } from \"../apiglue/mutableWritable.ts\"\n",
}

var to_add_to_mutable_ts_file = ""

func OnConfigSet() {
	to_add_to_mutable_ts_file = to_add_to_mutable_ts_file_beginning[Config.Framework]
}

var running_status = struct {
	gen_ran       bool
	converter_ran bool
	route_made    bool
}{
	gen_ran:       false,
	converter_ran: false,
	route_made:    false,
}

func Gen() {
	if !running_status.converter_ran && !running_status.route_made {
		panic("you should first run the converter or make a route so there is what to generate")
	}
	if running_status.gen_ran {
		panic("you should only run the generator once")
	}

	os.WriteFile(Config.Src_folder+"/routes.ts", []byte(to_add_to_ts_file), 0644)
	os.WriteFile(Config.Src_folder+"/mutables.ts", []byte(to_add_to_mutable_ts_file), 0644)
}
