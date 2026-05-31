package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const YYYYMMDD string = "20060102"

// Чтобы не принимать во внимание время, лучше написать специальную функцию afterNow(date, now time.Time) bool , которая будет возвращать true, если первая дата больше второй
func afterNow(date, now time.Time) bool {
	if date.Year() != now.Year() {
		return date.Year() > now.Year()
	}
	if date.Month() != now.Month() {
		return date.Month() > now.Month()
	}
	return date.Day() > now.Day()
}

// now — время, от которого ищется ближайшая дата
// dstart — исходное время в формате 20060102, от которого начинается отсчёт повторений
// repeat — правило повторения в описанном выше формате
func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if repeat == "" {
		return "", fmt.Errorf("пустая строка в параметре repeat")
	}
	repeatParts := strings.Split(repeat, " ")
	if (repeatParts[0] != "d") && (repeatParts[0] != "y") {
		return "", fmt.Errorf("неизвестный символ, обозначающий интервал переноса задачи")
	}
	if (len(repeatParts) != 2) && (repeatParts[0] == "d") {
		return "", fmt.Errorf("при указании периода переноса задачи введён посторонний символ")
	}
	if (len(repeatParts) != 1) && (repeatParts[0] == "y") {
		return "", fmt.Errorf("лишний параметр при указании ежегодного периода")
	}

	parseDstart, err := time.Parse(YYYYMMDD, dstart)
	if err != nil {
		return "", fmt.Errorf("ошибка конвертации исходного времени в дату типа формата \"ГГГГММДД\"")
	}

	/* 	if afterNow(parseDstart, now) {
		return "", fmt.Errorf("возвращаемая дата должна быть больше даты, указанной в переменной now")
	} */

	expirationTime := parseDstart
	if repeatParts[0] == "d" {
		dayPeriod, err := strconv.Atoi(repeatParts[1])
		if err != nil {
			return "", fmt.Errorf("ошибка конвертации числа, обозначающего период повтора задачи")
		}

		if dayPeriod > 400 {
			return "", fmt.Errorf("превышено максимально допустимое число, обозначающее период повтора в днях")
		}

		for {
			expirationTime = expirationTime.AddDate(0, 0, dayPeriod)
			if afterNow(expirationTime, now) {
				break
			}
		}
	} else if repeatParts[0] == "y" {
		for {
			expirationTime = expirationTime.AddDate(1, 0, 0)
			if afterNow(expirationTime, now) {
				break
			}
		}
	}

	return expirationTime.Format(YYYYMMDD), nil
}

// написать GET-обработчик запросов для api/nextdate.
// который должен принимать запросы в таком формате: "/api/nextdate?now=<20060102>&date=<20060102>&repeat=<правило>"
// Например: "api/nextdate?now=20240126&date=20240126&repeat=y"

// в обработчике нужно будет вызвать функцию NextDate
// и вернуть дату следующего выполнения в формате 20060102 или текст ошибки
// Если параметр now не определён, то тогда следует брать текущую дату

func nextDayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Сервер не поддерживает "+r.Method, http.StatusMethodNotAllowed)
		return
	}

	nowRequestStr := r.FormValue("now")
	dateRequestStr := r.FormValue("date")
	repeatRequestStr := r.FormValue("repeat")

	var now time.Time
	var err error
	if nowRequestStr == "" {
		now = time.Now()
	} else {
		now, err = time.Parse(YYYYMMDD, nowRequestStr)
		if err != nil {
			http.Error(w, "ошибка конвертации параметра now в дату типа формата \"ГГГГММДД\"", http.StatusBadRequest)
			return
		}
	}

	nextDate, err := NextDate(now, dateRequestStr, repeatRequestStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// временная заглушка к коду
	fmt.Fprintln(w, nextDate)
	fmt.Print(nextDate)
}
