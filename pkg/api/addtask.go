package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ProGuru/Final_golang_project/pkg/db"
)

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Нужно десериализовать полученный в запросе JSON в переменную var task db.Task.
	var task db.Task
	body, err := io.ReadAll(r.Body)
	if err != nil {
		writeErrorJson(w, map[string]any{"error": "невозможно прочитать тело запроса"})
		return
	}
	defer r.Body.Close()

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

	// Пришла очередь вызвать функцию db.AddTask(task), чтобы добавить задачу в базу данных
	id, err := db.AddTask(&task)
	if err != nil {
		writeErrorJson(w, map[string]any{"error": "произошла ошибка при добавлении задачи в БД"})
		return
	}

	// Осталось возвратить идентификатор добавленной задачи в виде JSON
	writeJson(w, map[string]any{"id": id})
}

func checkDate(task *db.Task) error {
	// берём текущее время
	now := time.Now()

	// если task.Date пустая строка, то присваиваем ему текущее время now.Format("20060102")
	if task.Date == "" {
		task.Date = now.Format("20060102")
	}

	// проверяем, что в task.Date указана корректная дата
	t, err := time.Parse("20060102", task.Date)
	if err != nil {
		return fmt.Errorf("дата представлена в формате, отличном от \"20060102\"")
	}

	// если определён task.Repeat, то проверяем корректность правила
	var next string
	if task.Repeat != "" {
		next, err = NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return fmt.Errorf("правило повторения указано в неправильном формате")
		}
	}

	// остаётся проверить, что дата t больше now
	// если сегодня (now) больше task.Date (t)
	if afterNow(now, t) {
		if len(task.Repeat) == 0 {
			// если правила повторения нет, то берём сегодняшнее число
			task.Date = now.Format("20060102")
		} else {
			// в противном случае, берём вычисленную ранее следующую дату
			task.Date = next
		}
	}

	return nil
}

func writeJson(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// сериализовать данные в JSON
	bytes, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "не удалось сериализовать JSON", http.StatusInternalServerError)
		return
	}

	w.Write(bytes)
}

func writeErrorJson(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)

	// сериализовать данные в JSON
	bytes, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "не удалось сериализовать JSON", http.StatusInternalServerError)
		return
	}

	w.Write(bytes)
}
