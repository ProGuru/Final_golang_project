package main

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/ProGuru/Final_golang_project/pkg/db"
	"github.com/ProGuru/Final_golang_project/pkg/server"
)

type Scheduler struct {
	id      int       // автоинкрементный идентификатор
	date    time.Time // дата задачи, которая будет храниться в формате YYYYMMDD или в Go-представлении 20060102
	title   string    // заголовок задачи
	comment string    // комментарий к задаче
	repeat  string    // строковое поле не более 128 символов, которое будет содержать правила повторений для задачи
}

const DB_FILE string = "scheduler.db" // название БД

func main() {
	// настройка подключения к БД
	var DBScheduler db.SchedulerStore
	err := DBScheduler.Init(DB_FILE)
	if err != nil {
		log.Fatal(err)
		return
	}

	// настройка логирования
	logFile, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	mylog := log.New(logFile, "serv ", log.LstdFlags|log.Lshortfile)

	baseDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// настройка и запуск сервера
	myServer := server.Run(mylog, filepath.Join(baseDir, "web"))
	log.Println("Сервер запускается на порту :7540")
	if err = myServer.Server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
