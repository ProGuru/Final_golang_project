package db

import "database/sql"

// параметры задачи
type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`    // дата задачи в формате 20060102
	Title   string `json:"title"`   // заголовок задачи. Обязательное поле
	Comment string `json:"comment"` // комментарий к задаче
	Repeat  string `json:"repeat"`  // правило повторения
}

/* Здесь же сразу можно написать функцию AddTask(task *Task) (int64, error), которая будет добавлять задачу в таблицу scheduler и возвращает идентификатор добавленной записи. Чтобы получить этот идентификатор, используйте значение типа Result, которое возвращают методы Exec() и ExecContext(), и вызовите у него метод LastInsertId().  */
func AddTask(task *Task) (int64, error) {
	var id int64
	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)`
	res, err := DB.Exec(query, sql.Named("date", task.Date), sql.Named("title", task.Title), sql.Named("comment", task.Comment), sql.Named("repeat", task.Repeat))

	if err == nil {
		id, err = res.LastInsertId()
	}
	return id, err
}

/* Вам нужно вспомнить, как отправляются запросы SELECT и как потом происходит сканирование результата с помощью методов Next() и Scan(). */
/* func Tasks(limit int) ([]*Task, error) {
	return
} */
