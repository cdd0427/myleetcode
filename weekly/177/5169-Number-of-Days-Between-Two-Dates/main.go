package main

import "time"

func daysBetweenDates(date1 string, date2 string) int {
	layout := "2006-01-02 15:04:05"
	a, _ := time.Parse(layout, date1+" 00:00:00")
	b, _ := time.Parse(layout, date2+" 00:00:00")
	d := b.Sub(a)
	res := int(d.Hours() / 24)
	if res > 0 {
		return res
	}
	return -res
}
