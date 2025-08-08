package apiglue

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

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

// //////
// sync headers
// ///////
func (this *ServerState[T]) Add_state_header(c *gin.Context) {
	header_json, _ := json.Marshal(MutableStateSender{
		Type: "mutable-state-sender",
		Key:  this.Key,
		Data: this.State,
	})
	c.Header("sync", string(header_json))
}

func (this *ServerState[T]) Add_update_header(c *gin.Context, path string, new_data any) {
	header_json, _ := json.Marshal(MutableUpdateMessage{
		Type:    "mutable-update",
		Key:     this.Key,
		Path:    path,
		NewData: new_data,
	})
	c.Header("sync", string(header_json))
}

func (this *ServerState[T]) Add_delete_header(c *gin.Context, path string) {
	header_json, _ := json.Marshal(MutableDeleteMessage{
		Type: "mutable-delete",
		Key:  this.Key,
		Path: path,
	})
	c.Header("sync", string(header_json))
}

func (this *ServerState[T]) Add_append_header(c *gin.Context, path string, new_data any) {
	header_json, _ := json.Marshal(MutableAppendMessage{
		Type:    "mutable-append",
		Key:     this.Key,
		Path:    path,
		NewData: new_data,
	})
	c.Header("sync", string(header_json))
}
