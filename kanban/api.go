package kanban

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Move_task(c *gin.Context, task_id int, list string) {
	fmt.Printf("in move_task params are task_id: %d, list: %s\n", task_id, list)
	if !list_contains(Main_Board.State.Lists, list) {
		c.JSON(400, gin.H{
			"error": "list must be one of " + strings.Join(Main_Board.State.Lists, ", "),
		})
		return
	}
	index := list_find_index(Main_Board.State.Tasks, func(task Task) bool {
		return task.Id == task_id
	})
	Main_Board.State.Tasks[index].List = list
	Main_Board.Add_update_header(c, "tasks."+strconv.Itoa(task_id)+".list", list)
	c.JSON(200, gin.H{
		"message": "task moved",
	})

}

func Create_task(c *gin.Context, title string, list string, author string, deadline string) {
	//basic bouncer
	fmt.Printf("in create_task params are title: %s, list: %s, author: %s, deadline: %s\n", title, list, author, deadline)
	if title == "" {
		c.JSON(400, gin.H{
			"error": "title is required",
		})
		return
	}
	if list == "" {
		c.JSON(400, gin.H{
			"error": "list is required",
		})
		return
	}
	if author == "" {
		c.JSON(400, gin.H{
			"error": "author is required",
		})
		return
	}
	if deadline == "" {
		c.JSON(400, gin.H{
			"error": "deadline is required",
		})
		return
	}
	//end basic bouncer
	//bouncer
	if !list_contains(Main_Board.State.Users, author) {
		c.JSON(400, gin.H{
			"error": "author must be one of " + strings.Join(Main_Board.State.Users, ", "),
		})
		return
	}
	if !list_contains(Main_Board.State.Lists, list) {
		c.JSON(400, gin.H{
			"error": "list must be one of " + strings.Join(Main_Board.State.Lists, ", "),
		})
		return
	}
	//end bouncer

	Main_Board.State.Tasks = append(Main_Board.State.Tasks, Task{
		Title:    title,
		List:     list,
		Author:   author,
		Time:     time.Now().Format("2006-01-02 15:04:05"),
		Deadline: deadline,
		Id:       len(Main_Board.State.Tasks) + 1,
	})
	Main_Board.Add_append_header(c, "tasks", Task{
		Title:    title,
		List:     list,
		Author:   author,
		Time:     time.Now().Format("2006-01-02 15:04:05"),
		Deadline: deadline,
		Id:       len(Main_Board.State.Tasks) + 1,
	})
	c.JSON(200, gin.H{
		"message": "task created",
	})
}

func Get_board(c *gin.Context) {
	Main_Board.Add_state_header(c)
	c.JSON(200, gin.H{
		"message": "board fetched",
	})
}
