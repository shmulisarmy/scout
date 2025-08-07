package apiglue

import (
	"context"
	"fmt"
	"gin-sevalla-app/api_glue/utils"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// type Recipe struct {
// 	Name         string   `json:"name"`
// 	Ingredients  []string `json:"ingredients"`
// 	Instructions []string `json:"instructions"`
// }

// var available_ingredients = ServerState[map[string]int]{
// 	State: map[string]int{
// 		"flour":           1,
// 		"sugar":           2,
// 		"eggs":            3,
// 		"milk":            4,
// 		"salt":            5,
// 		"oil":             6,
// 		"vanilla":         7,
// 		"butter":          8,
// 		"baking powder":   9,
// 		"vanilla extract": 10,
// 	},
// 	Key: "available_ingredients",
// }

const src_folder = "my-svelte-app/src/routes"

var html_base_page_string = load_from_file("templates/base_page.html")

func load_from_file(file_name string) string {
	file, err := os.ReadFile(file_name)
	if err != nil {
		panic(err)
	}
	return string(file)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]string)

func get_or_set_profile_id(c *gin.Context) string {
	var profile_id string
	if profile_id, _ = c.Cookie("profile_id"); profile_id == "" {
		fmt.Println("setting profile id")
		profile_id_int := redis_db.Incr(ctx, "profile_id_upto").Val()
		fmt.Println("profile_id_int", profile_id_int)
		profile_id = strconv.FormatInt(profile_id_int, 10)
		c.SetCookie("profile_id", profile_id, int(time.Hour*24), "/", "", false, true)
	} else {
		fmt.Println("using profile id", profile_id)
	}
	return profile_id
}

func ws_handler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()
	clients[conn] = get_or_set_profile_id(c)

	for {
		log.Println("Waiting for message...")
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		log.Printf("Received: %s", message)

		// Respond with "server received"
		if err := conn.WriteMessage(websocket.TextMessage, []byte("server received")); err != nil {
			log.Println("Write error:", err)
			break
		}
	}
	delete(clients, conn)
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

var to_add_to_mutable_ts_file string = "import { mutableWritable } from \"./mutableWritable.ts\"\n"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "5001"
	}

	converter := Ts_Type_Converter{
		parsed: map[string]bool{},
		queue:  []reflect.Type{},
		file:   src_folder + "/types.ts",
	}

	converter.Convert()

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "http://localhost:5173"}
	config.AllowCredentials = true
	r.Use(cors.New(config))
	r.LoadHTMLGlob("templates/*")

	os.WriteFile(src_folder+"/routes.ts", []byte(to_add_to_ts_file), 0644)
	os.WriteFile(src_folder+"/mutables.ts", []byte(to_add_to_mutable_ts_file), 0644)
}
