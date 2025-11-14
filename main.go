package main

import (
	"fmt"
	"p/calendar"
	"time"
)

func main() {
	fmt.Println("Календарь")
	now := time.Now()
	for i := 0; i < 12; i++ {
		calendar.DrawCalendar(now.AddDate(0, i, 0))
	}
	fmt.Println(":^)")
	fmt.Println("Я лосось поющий, на пляжу ляжущий")
//этот текст для проверки fetch
}
