package calendar

import (
	"fmt"
	"time"
)

func isLeap(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}
func howDaysInMonyh(m int, flag bool) int {
	switch m {
	case 1:
		return 31
	case 2:
		if flag {
			return 29
		}
		return 28
	case 3:
		return 31
	case 4:
		return 30
	case 5:
		return 31
	case 6:
		return 30
	case 7:
		return 31
	case 8:
		return 31
	case 9:
		return 30
	case 10:
		return 31
	case 11:
		return 30
	}
	return 31
}
func meet(firstDay time.Time, numDays int) []string {
	w := int(firstDay.Weekday())
	res := make([]string, 0, numDays)
	for i := 0; i < w; i++ {
		res = append(res, "  ")
	}
	for i := 1; i <= numDays; i++ {
		if i < 10 {
			res = append(res, fmt.Sprintf("0%d", i))
			continue
		}
		res = append(res, fmt.Sprintf("%d", i))
	}
	return res
}
func DrawCalendar(t time.Time) {
	y := t.Year()
	m := t.Month()
	firstDayMonth := time.Date(y, m, 1, 0, 0, 0, 1, time.UTC)
	days := meet(firstDayMonth, howDaysInMonyh(int(m), isLeap(y)))
	fmt.Printf("\n\n\n--------%d--------\n", y)
	fmt.Println(m, "\n")
	fmt.Println("Su Mo Tu We Th Fr Sa")
	for i, day := range days {
		fmt.Print(day)
		if i%7 == 0 {
			fmt.Print("\n")
			continue
		}
		fmt.Print(" ")
	}

}
