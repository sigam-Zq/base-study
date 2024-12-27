package main

import (
	"log"
	"time"
)

func main() {
	log.Println(initHoliday(2025))
}

func initHoliday(yearNumber int) []string {

	dayOffList := make([]string, 0)
	for i := time.Date(yearNumber, time.January, 1, 1, 0, 0, 0, time.Local); i.Before(time.Date(yearNumber+1, time.January, 0, 0, 0, 0, 0, time.Local)); i = i.Add(24 * time.Hour) {
		if week := i.Weekday(); week == time.Sunday || week == time.Saturday {
			dayOffList = append(dayOffList, i.Format(time.DateOnly))
		}
	}

	return dayOffList
}
