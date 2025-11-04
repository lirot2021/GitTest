package main

import (
	"fmt"
	"p/calendar"
	"time"
)

func main() {
	fmt.Println("Календарь")
	now := time.Now()
	for i := 0; i < 12; i++{
		calendar.DrawCalendar(now.Add(time.Month*i))
	}
}
