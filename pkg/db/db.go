package db

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

const schema = `CREATE TABLE scheduler (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	date CHAR(8) NOT NULL DEFAULT "",
	title VARCHAR(64) NOT NULL DEFAULT "",
	comment TEXT NOT NULL DEFAULT "",
	repeat VARCHAR(128) NOT NULL DEFAULT "")`

var DB *sql.DB

func Init(dbFileName string) error {
	// проверяем, существует ли файл БД. Если нет - то создаём
	_, err := os.Stat(dbFileName)

	var install bool // по умолчанию false
	if err != nil {
		install = true
	}

	db, err := sql.Open("sqlite", dbFileName)
	if err != nil {
		return err
	}
	DB = db

	// если install равен true, после открытия БД требуется выполнить
	// sql-запрос с CREATE TABLE и CREATE INDEX
	if install {
		_, err := DB.Exec(schema)
		if err != nil {
			return err
		}
	}

	return nil
}
