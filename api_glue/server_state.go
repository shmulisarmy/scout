package apiglue

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/gorilla/websocket"
)

type MutableStateSender struct {
	Type string `json:"type" default:"mutable-state-sender"`
	Key  string `json:"key"`
	Data any    `json:"data"`
}

type ServerState[T any] struct {
	State       T      `json:"state"`
	Key         string `json:"key"`
	Client_List []*websocket.Conn
}

func (this *ServerState[T]) Add_to_ts() {
	converter := Ts_Type_Converter{
		Parsed: map[string]bool{},
		Queue: []reflect.Type{
			reflect.TypeOf(this.State),
		},
		File: Config.Src_folder + "/mutables.ts",
	}
	j, _ := json.Marshal(this.State)
	to_add_to_mutable_ts_file += converter.Convert()
	state_type_name := typeToTSType(reflect.TypeOf(this.State), &converter.Queue, converter.Parsed)
	switch Config.Framework {
	case "zustand":
		to_add_to_mutable_ts_file += fmt.Sprintf(`
export const use%sStore = create<{state: %s}>((set) => ({
	state: %v,
}))
if (typeof window !== 'undefined') {
	(window as any).%s = use%sStore
}
		
		`, this.Key, state_type_name, string(j), this.Key, this.Key)
	case "svelte":
		to_add_to_mutable_ts_file += fmt.Sprintf(`
			export const %s = mutableWritable<%s>(%v)
			if (typeof window !== 'undefined') {
				(window as any).%s = %s
			}
		`, this.Key, state_type_name, string(j), this.Key, this.Key)
	}

}

func (this *ServerState[T]) Send_state() {
	for _, client := range this.Client_List {
		client.WriteJSON(MutableStateSender{
			Type: "mutable-state-sender",
			Key:  this.Key,
			Data: this.State,
		})
	}
}

func (this *ServerState[T]) Send_update(path string, new_data any) {
	for _, client := range this.Client_List {
		client.WriteJSON(MutableUpdateMessage{
			Type:    "mutable-update",
			Key:     this.Key,
			Path:    path,
			NewData: new_data,
		})
	}
}

func (this *ServerState[T]) Send_delete(path string) {
	for _, client := range this.Client_List {
		client.WriteJSON(MutableDeleteMessage{
			Type: "mutable-delete",
			Key:  this.Key,
			Path: path,
		})
	}
}

func (this *ServerState[T]) Onboard_client(client *websocket.Conn) {
	this.Client_List = append(this.Client_List, client)
	this.Send_state()
}

func (this *ServerState[T]) Send_append(path string, new_data any) {
	for _, client := range this.Client_List {
		client.WriteJSON(MutableAppendMessage{
			Type:    "mutable-append",
			Key:     this.Key,
			Path:    path,
			NewData: new_data,
		})
	}
}
