package main

import (
	"fmt"
	"p/calendar"
	"time"
)

func main() {
	fmt.Println("Календарь")
	now := time.Now()
	calendar.DrawCalendar(now)
}
