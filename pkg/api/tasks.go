package api

import (
	"fmt"
	"net/http"

	"github.com/ProGuru/Final_golang_project/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.Tasks(50) // в параметре максимальное количество записей
	if err != nil {
		// здесь вызываете функцию, которая возвращает ошибку в JSON
		// её желательно было реализовать на предыдущем шаге
		writeErrorJson(w, map[string]any{"error": fmt.Errorf("ошибка при сканировании результата запроса: %w", err)})
		return
	}
	writeJson(w, TasksResp{
		Tasks: tasks,
	})
}
