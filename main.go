package main

import (
	apiglue "gin-sevalla-app/api_glue"
	"gin-sevalla-app/kanban"
	"log"
	"reflect"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// func get_or_set_profile_id(c *gin.Context) string {
// 	var profile_id string
// 	if profile_id, _ = c.Cookie("profile_id"); profile_id == "" {
// 		fmt.Println("setting profile id")
// 		profile_id_int := redis_db.Incr(ctx, "profile_id_upto").Val()
// 		fmt.Println("profile_id_int", profile_id_int)
// 		profile_id = strconv.FormatInt(profile_id_int, 10)
// 		c.SetCookie("profile_id", profile_id, int(time.Hour*24), "/", "", false, true)
// 	} else {
// 		fmt.Println("using profile id", profile_id)
// 	}
// 	return profile_id
// }

var clients = make(map[*websocket.Conn]string)

func ws_handler(c *gin.Context) {
	conn, err := websocket.Upgrade(c.Writer, c.Request, nil, 1024, 1024)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()
	// clients[conn] = get_or_set_profile_id(c)

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

var todo_id_upto = 0

type Todo struct {
	Title          string `json:"title"`
	Done           bool   `json:"done"`
	Id             int    `json:"id"`
	Estimated_time string `json:"estimated_time"`
	Created_at     string `json:"created_at"`
}

func new_todo(title string) Todo {
	todo_id_upto++
	return Todo{
		Title:          title,
		Done:           false,
		Id:             todo_id_upto,
		Estimated_time: "",
		Created_at:     time.Now().Format("2006-01-02 15:04:05"),
	}
}

var todos = apiglue.NewServerState([]Todo{
	new_todo("todo 1"),
	new_todo("todo 2"),
})

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"sync"},
		AllowCredentials: false, // must be false when AllowOrigins = ["*"]
		MaxAge:           12 * time.Hour,
	}))

	apiglue.Config.Src_folder = "./frontend/generated"
	apiglue.Config.Port = "8080"
	apiglue.Config.Framework = "zustand"
	apiglue.OnConfigSet()

	converter := apiglue.Ts_Type_Converter{
		Parsed: map[string]bool{},
		Queue: []reflect.Type{
			reflect.TypeOf(Todo{}),
		},
		File: "./frontend/generated/types.ts",
	}
	converter.Convert()

	apiglue.Make_route(r, "api/get_todos", func(c *gin.Context) {
		todos.Add_state_header(c)
		c.JSON(200, gin.H{
			"message": "todos fetched",
		})
	})

	apiglue.Make_route(r, "api/add_todo", func(c *gin.Context, new_todo_name string) {
		todos.State = append(todos.State, new_todo(new_todo_name))
		todos.Add_append_header(c, "", new_todo(new_todo_name))
		c.JSON(200, gin.H{
			"message": "todo added",
		})
	})

	apiglue.Make_route(r, "api/ws", ws_handler)

	apiglue.Make_route(r, "api/delete_todo", func(c *gin.Context, id int) {
		for i := range todos.State {
			if todos.State[i].Id == id {
				todos.State = append(todos.State[:i], todos.State[i+1:]...)
				break
			}
		}
		todos.Add_state_header(c)
	})

	people.Add_to_ts()
	todos.Add_to_ts()

	//kanban
	apiglue.Make_route(r, "api/kanban/move_task", kanban.Move_task)
	apiglue.Make_route(r, "api/kanban/create_task", kanban.Create_task)
	apiglue.Make_route(r, "api/kanban/get_board", kanban.Get_board)
	apiglue.Make_route(r, "api/kanban/add_comment", kanban.Add_comment)
	apiglue.Make_route(r, "api/kanban/delete_task", kanban.Delete_task)
	kanban.Main_Board.Add_to_ts()
	//kanban

	apiglue.Gen()

	address := "0.0.0.0:" + apiglue.Config.Port
	log.Println("Starting server on", address)

	if err := r.Run(address); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

var people = apiglue.NewServerState([]Person{
	{
		Name:  "shmuli",
		Age:   21,
		Email: "shmuli@shmuli.com",
	},
	{
		Name:  "berel",
		Age:   25,
		Email: "berel@shmuli.com",
	},
})

var live_scout = apiglue.NewServerState(Live_Scout{
	Links:        make(map[string]struct{}),
	To_Scout_For: "tell me when there is an ai truck driver in the state of florida",
})
