package api

import (
	"net/http"
	"path/filepath"
)

// Для регистрации API обработчиков
func Init(webDir string) {
	http.Handle("/", http.FileServer(http.Dir(filepath.Clean(webDir)))) // Регистрируем обработчик для корневого пути, который будет отдавать статические файлы из папки web
	http.HandleFunc("/api/nextdate", nextDayHandler)                    // Регистрируем обработчик для пути /api/nextdate, который будет обрабатывать запросы на получение следующей даты задачи
	http.HandleFunc("/api/task", taskHandler)                           // Добавляем задачу
	//http.HandleFunc("/api/tasks", tasksHandler)                         // Получаем список ближайших задач
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// обработка других методов будет добавлена на следующих шагах
	case http.MethodPost:
		addTaskHandler(w, r)
	}
}
