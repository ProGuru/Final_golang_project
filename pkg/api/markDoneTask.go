package api

import (
	"net/http"
	"time"

	"github.com/ProGuru/Final_golang_project/pkg/db"
)

func markDoneTaskHandler(w http.ResponseWriter, r *http.Request) {
	// обработчик для POST-запроса /api/task/done, который делает задачу выполненной
	id := r.URL.Query().Get("id")
	if id == "" {
		writeErrorJson(w, map[string]any{"error": "id задачи не указан"})
		return
	}

	// Необходимо получить параметры задачи по её id. Для этого уже есть функция db.GetTask()
	task, err := db.GetTask(id)
	if err != nil {
		writeErrorJson(w, map[string]any{"error": "не удалось получить параметры задачи по её id"})
		return
	}

	// Если правило для повторения отсутствует, то необходимо удалить задачу.
	if task.Repeat == "" {
		err = db.DeleteTask(id)
		if err != nil {
			writeErrorJson(w, map[string]any{"error": "ошибка при удалении задачи из БД"})
			return
		}
	} else {
		// Если задача периодическая, то необходимо вызвать NextDate() для получения следующей даты
		// В этом случае нужно обновить дату у задачи в базе данных
		nextDate, err := NextDate(time.Now(), task.Date, task.Repeat)
		if err != nil {
			writeErrorJson(w, map[string]any{"error": "ошибка при вычислении следующей даты для периодической задачи"})
			return
		}

		// И обновить дату задачи в БД на полученное значение
		err = db.UpdateDate(nextDate, id)
		if err != nil {
			writeErrorJson(w, map[string]any{"error": "ошибка при обновлении задачи в БД"})
			return
		}
	}

	// В случае успешного изменения должен возвращаться пустой JSON {}
	writeJson(w, map[string]any{})
}
