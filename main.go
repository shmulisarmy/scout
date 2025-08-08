package main

import (
	"fmt"
	apiglue "gin-sevalla-app/api_glue"
	"log"

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

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "http://localhost:5173"}
	config.AllowCredentials = true
	r.Use(cors.New(config))
	// r.LoadHTMLGlob("templates/*")

	// r.GET("/scouters", func(c *gin.Context) {
	// 	for _, live_scout := range live_scouts {
	// 		report := live_scout.Scout()
	// 		d := DisplayStructDetailed(report)
	// 		fmt.Println(d)
	// 		d = DisplayStructDetailed(live_scout)
	// 		fmt.Println(d)
	// 	}
	// })
	apiglue.Config.Src_folder = "./frontend/generated"
	apiglue.Config.Port = "8080"
	apiglue.Config.Framework = "zustand"
	apiglue.OnConfigSet()
	// converter := apiglue.Ts_Type_Converter{
	// 	Parsed: map[string]bool{},
	// 	Queue: []reflect.Type{
	// 		reflect.TypeOf(live_scouts[0]),
	// 		reflect.TypeOf(live_scout.State),
	// 	},
	// 	File: apiglue.Config.Src_folder + "/types.ts",
	// }
	// converter.Convert()
	// scouts := apiglue.NewServerState(live_scouts)
	live_scouts.Add_to_ts()
	live_scout.Add_to_ts()
	apiglue.Make_route(r, "ws", ws_handler)

	apiglue.Make_route(r, "scouters", func(c *gin.Context) {
		fmt.Println("scouters")
		for _, live_scout := range live_scouts.State {
			report := live_scout.Scout()
			d := DisplayStructDetailed(report)
			fmt.Println(d)
			d = DisplayStructDetailed(live_scout)
			fmt.Println(d)
		}
	})

	people.Add_to_ts()

	apiglue.Gen()

	r.Run("localhost:" + apiglue.Config.Port)
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
