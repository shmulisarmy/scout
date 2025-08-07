package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var todos = []Todo{
	{
		ID:    1,
		Title: "Todo 1",
		Done:  false,
	},
	{
		ID:    2,
		Title: "Todo 2",
		Done:  false,
	},
}
var todo_upto int = 2

func main() {
	r := gin.Default()
	r.GET("/todos", func(c *gin.Context) {
		c.JSON(http.StatusOK, todos)
	})
	r.GET("/todos/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid id",
			})
			return
		}
		for i, todo := range todos {
			if todo.ID == id {
				c.JSON(http.StatusOK, todos[i])
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{
			"message": "todo not found",
		})
	})
	r.GET("/create-todo", func(c *gin.Context) {
		new_todo := Todo{
			ID:    todo_upto + 1,
			Title: "Todo " + strconv.Itoa(todo_upto+1),
			Done:  false,
		}
		todos = append(todos, new_todo)
		todo_upto++
		c.JSON(http.StatusOK, gin.H{
			"message": new_todo,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
