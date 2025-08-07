package apiglue

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type MutableStateSender struct {
	Type string `json:"type" default:"mutable-state-sender"`
	Key  string `json:"key"`
	Data any    `json:"data"`
}

type ServerState[T any] struct {
	State T      `json:"state"`
	Key   string `json:"key"`
}

func (this ServerState[T]) add_to_ts() {
	converter := Ts_Type_Converter{
		parsed: map[string]bool{},
		queue: []reflect.Type{
			reflect.TypeOf(this.State),
		},
		file: src_folder + "/mutables.ts",
	}
	j, _ := json.Marshal(this.State)
	to_add_to_mutable_ts_file += converter.Convert()
	state_type_name := typeToTSType(reflect.TypeOf(this.State), &converter.queue, converter.parsed)
	to_add_to_mutable_ts_file += fmt.Sprintf(`
	export const %s = mutableWritable<%s>(%v)
	if (typeof window !== 'undefined') {
		window.%s = %s
	}
	`, this.Key, state_type_name, string(j), this.Key, this.Key)

}

func (this *ServerState[T]) send_state() {
	for client := range clients {
		client.WriteJSON(MutableStateSender{
			Type: "mutable-state-sender",
			Key:  this.Key,
			Data: this.State,
		})
	}
}

func (this *ServerState[T]) send_update(path string, new_data any) {
	for client := range clients {
		client.WriteJSON(MutableUpdateMessage{
			Type:    "mutable-update",
			Key:     this.Key,
			Path:    path,
			NewData: new_data,
		})
	}
}

func (this *ServerState[T]) send_delete(path string) {
	for client := range clients {
		client.WriteJSON(MutableDeleteMessage{
			Type: "mutable-delete",
			Key:  this.Key,
			Path: path,
		})
	}
}

func (this *ServerState[T]) send_append(path string, new_data any) {
	for client := range clients {
		client.WriteJSON(MutableAppendMessage{
			Type:    "mutable-append",
			Key:     this.Key,
			Path:    path,
			NewData: new_data,
		})
	}
}
