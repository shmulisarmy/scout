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
	Main_Board.Add_update_header(c, "tasks."+strconv.Itoa(index)+".list", list)
	c.JSON(200, gin.H{
		"message": "task moved",
	})

}

func Add_comment(c *gin.Context, task_id int, author string, body string) {
	task_exists := list_find_index(Main_Board.State.Tasks, func(task Task) bool {
		return task.Id == task_id
	}) != -1
	if !task_exists {
		c.JSON(400, gin.H{
			"error": "task does not exist",
		})
		return
	}
	if author == "" {
		c.JSON(400, gin.H{
			"error": "author is required",
		})
		return
	}
	if body == "" {
		c.JSON(400, gin.H{
			"error": "no comment provided",
		})
		return
	}
	Main_Board.State.Comments = append(Main_Board.State.Comments, Comment{
		Id:     len(Main_Board.State.Comments) + 1,
		Body:   body,
		Author: author,
		TaskId: task_id,
	})
	Main_Board.Add_append_header(c, "comments", Comment{
		Id:     len(Main_Board.State.Comments) + 1,
		Body:   body,
		Author: author,
		TaskId: task_id,
	})
	c.JSON(200, gin.H{
		"message": "comment added",
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

	task_id_upto += 1
	Main_Board.State.Tasks = append(Main_Board.State.Tasks, Task{
		Title:    title,
		List:     list,
		Author:   author,
		Time:     time.Now().Format("2006-01-02 15:04:05"),
		Deadline: deadline,
		Id:       task_id_upto,
	})
	Main_Board.Add_append_header(c, "tasks", Task{
		Title:    title,
		List:     list,
		Author:   author,
		Time:     time.Now().Format("2006-01-02 15:04:05"),
		Deadline: deadline,
		Id:       task_id_upto,
	})
	c.JSON(200, gin.H{
		"message": "task created",
	})
}

func Delete_task(c *gin.Context, task_id int) {
	index := list_find_index(Main_Board.State.Tasks, func(task Task) bool {
		return task.Id == task_id
	})
	if index == -1 {
		c.JSON(400, gin.H{
			"error": "task does not exist",
		})
		return
	}
	Main_Board.State.Tasks = append(Main_Board.State.Tasks[:index], Main_Board.State.Tasks[index+1:]...)
	Main_Board.Add_delete_header(c, "tasks."+strconv.Itoa(index))
	c.JSON(200, gin.H{
		"message": "task deleted",
	})
}

var task_id_upto = 5

func Get_board(c *gin.Context) {
	Main_Board.Add_state_header(c)
	c.JSON(200, gin.H{
		"message": "board fetched",
	})
}
