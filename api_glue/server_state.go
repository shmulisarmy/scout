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
	default:
		panic("unknown framework")
	}

}
