package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ProGuru/Final_golang_project/pkg/db"
)

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	// данные нужно проверять так же, как при добавлении задачи
	var task db.Task
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		writeErrorJson(w, map[string]any{"error": "невозможно прочитать тело запроса"})
		return
	}

	// Нужно десериализовать полученный в запросе JSON в переменную var task db.Task.
	err = json.Unmarshal(body, &task)
	if err != nil {
		writeErrorJson(w, map[string]any{"error": "ошибка десериализации JSON"})
		return
	}

	// Проверить, что поле task.Title не пустое
	if task.Title == "" {
		writeErrorJson(w, map[string]any{"error": "не указан заголовок задачи"})
		return
	}

	// Проверить на корректность полученное значение
	if err := checkDate(&task); err != nil {
		writeErrorJson(w, map[string]any{"error": err.Error()})
		return
	}

	// Обновляем задачу в БД
	err = db.UpdateTask(&task)
	if err != nil {
		writeErrorJson(w, map[string]any{"error": "произошла ошибка при обновлении задачи в БД"})
		return
	}

	// В случае успешного изменения должен возвращаться пустой JSON {}
	writeJson(w, json.NewEncoder(w).Encode(nil))
}
