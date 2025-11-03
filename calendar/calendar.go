package calendar

import (
	"fmt"
	"time"
)

func DrawCalendar(t time.Time) {
	y := t.Year()
	m := t.Month()
	fmt.Printf("----%d----\n", y)
	fmt.Println(m)
}
