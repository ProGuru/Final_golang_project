package api

import (
	"net/http"
	"strconv"

	"github.com/ProGuru/Final_golang_project/pkg/db"
)

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		writeErrorJson(w, map[string]any{"error": "id задачи не указан"})
		return
	}

	// Проверить на корректность полученное значение
	_, err := strconv.Atoi(id)
	if err != nil {
		writeErrorJson(w, map[string]any{"error": "ошибка при конвертации id в целое число"})
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		writeErrorJson(w, map[string]any{"error": "произошла ошибка при поиске задачи в БД"})
		return
	}

	writeJson(w, task)
}
