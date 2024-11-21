package main

import (
	"log"
	"strconv"
	"time"
)

func main() {
	GetDayOff("2025")
}

func GetDayOff(year string) {
	yearNumber, err := strconv.Atoi(year)
	if err != nil {
		panic(err)
	}
	dayOffList := make([]string, 0)
	for i := time.Date(yearNumber, time.January, 0, 0, 0, 0, 0, time.Local); i.Before(time.Date(yearNumber+1, time.January, 0, 0, 0, 0, 0, time.Local)); i = i.Add(24 * time.Hour) {
		if week := i.Weekday(); week == time.Sunday || week == time.Saturday {
			dayOffList = append(dayOffList, i.Format(time.DateOnly))
		}
	}

	log.Println(len(dayOffList))
	log.Println("00000")
	log.Println(dayOffList)
	log.Println("00000")
}
