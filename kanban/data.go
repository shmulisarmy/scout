package kanban

import apiglue "gin-sevalla-app/api_glue"

var board_state = Kanban{
	Lists: []string{"To Do", "In Progress", "Done"},
	Tasks: []Task{
		{
			Title:    "Task 1",
			List:     "To Do",
			Author:   "Shmuli",
			Time:     "2025-08-09 21:13:09",
			Deadline: "2025-08-09 21:13:09",
			Id:       1,
		},
		{
			Title:    "Task 2",
			List:     "In Progress",
			Author:   "Shmuli",
			Time:     "2025-08-09 21:13:09",
			Deadline: "2025-08-09 21:13:09",
			Id:       2,
		},
		{
			Title:    "Task 3",
			List:     "Done",
			Author:   "yosef",
			Time:     "2025-08-09 21:13:09",
			Deadline: "2025-08-09 21:13:09",
			Id:       3,
		},
	},
	Comments: []Comment{
		{
			Id:     1,
			Body:   "Comment 1",
			Author: "Shmuli",
			TaskId: 1,
		},
		{
			Id:     2,
			Body:   "Comment 2",
			Author: "yosef",
			TaskId: 3,
		},
	},
	Users: []string{"Shmuli", "yosef"},
}

var Main_Board = apiglue.NewServerState(board_state)

type Kanban struct {
	Lists    []string  `json:"lists"`
	Tasks    []Task    `json:"tasks"`
	Comments []Comment `json:"comments"`
	Users    []string  `json:"users"`
}

type Task struct {
	Title    string `json:"title"`
	List     string `json:"list"`
	Author   string `json:"author"`
	Time     string `json:"time"`
	Deadline string `json:"deadline"`
	Id       int    `json:"id"`
}

type Comment struct {
	Id     int    `json:"id"`
	Body   string `json:"body"`
	Author string `json:"author"`
	TaskId int    `json:"task_id"`
}
