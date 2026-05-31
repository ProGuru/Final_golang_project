package api

import (
	"github.com/ProGuru/Final_golang_project/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

/* func tasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.Tasks(50) // в параметре максимальное количество записей
	if err != nil {
		// здесь вызываете функцию, которая возвращает ошибку в JSON
		// её желательно было реализовать на предыдущем шаге
		// ...
		return
	}
	writeJson(w, TasksResp{
		Tasks: tasks,
	})
} */
