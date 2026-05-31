package db

import (
	"database/sql"
	"fmt"
)

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
	res, err := db.Exec(query, sql.Named("date", task.Date), sql.Named("title", task.Title), sql.Named("comment", task.Comment), sql.Named("repeat", task.Repeat))

	if err == nil {
		id, err = res.LastInsertId()
	}
	return id, err
}

/* Вам нужно вспомнить, как отправляются запросы SELECT и как потом происходит сканирование результата с помощью методов Next() и Scan(). */
func Tasks(limit int) ([]*Task, error) {
	// выбираем все поля в порядке увеличения по дате
	query := `SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date ASC`
	rows, err := db.Query(query)

	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса к БД: %w", err)
	}
	defer rows.Close()

	var res []*Task
	var count int
	for rows.Next() {
		p := Task{}

		err := rows.Scan(&p.ID, &p.Date, &p.Title, &p.Comment, &p.Repeat)
		if err != nil {
			return nil, fmt.Errorf("ошибка при сканировании результата запроса: %w", err)
		}

		res = append(res, &p)

		count++
		if count > limit {
			break
		}
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при сканировании результата запроса: %w", err)
	}
	if res == nil {
		res = make([]*Task, 0) // возвращаем пустой слайс, если в БД нет задач
	}

	return res, nil
}

// функция по указанному id возвратит структуру или указатель на структуру Task
func GetTask(id string) (*Task, error) {
	var task Task
	err := db.QueryRow(`SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?`, id).Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса к БД: %w", err)
	}

	return &task, nil
}

// функция для сохранения результатов редактирования
func UpdateTask(task *Task) error {
	// параметры пропущены, не забудьте указать WHERE
	query := `UPDATE scheduler SET date = :date, title = :title, comment = :comment, repeat = :repeat WHERE id = :id`
	res, err := db.Exec(query, sql.Named("date", task.Date), sql.Named("title", task.Title), sql.Named("comment", task.Comment), sql.Named("repeat", task.Repeat), sql.Named("id", task.ID))

	if err != nil {
		return fmt.Errorf("ошибка при выполнении запроса к БД: %w", err)
	}
	// метод RowsAffected() возвращает количество записей к которым
	// была применена SQL команда
	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка при выполнении запроса к БД: %w", err)
	}

	if count == 0 {
		return fmt.Errorf(`incorrect id for updating task`)
	}
	return nil
}

func UpdateDate(next string, id string) error {
	// по сути, отличается от db.UpdateTask() только запросом, так как нужно изменить только колонку date, а не все поля
	query := `UPDATE scheduler SET date = :date WHERE id = :id`
	res, err := db.Exec(query, sql.Named("date", next), sql.Named("id", id))

	if err != nil {
		return fmt.Errorf("ошибка при выполнении запроса к БД: %w", err)
	}
	// метод RowsAffected() возвращает количество записей к которым
	// была применена SQL команда
	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка при выполнении запроса к БД: %w", err)
	}

	if count == 0 {
		return fmt.Errorf(`incorrect id for updating task`)
	}
	return nil
}

func DeleteTask(id string) error {
	// нужно удалить из таблицы scheduler задачу с указанным идентификатором
	query := `DELETE FROM scheduler WHERE id = ?`
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("ошибка при выполнении запроса к БД: %w", err)
	}
	return nil
}
