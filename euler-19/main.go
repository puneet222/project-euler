package main

import "fmt"

var daysInMonth = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func isLeapYear(year int) bool {
	if year % 100 == 0 && year % 400 != 0 {
		return false
	}
	if year % 4 == 0 {
		return true
	}
	return false
}

func countSundays(year int, start int, limit int) int {
	count := 0
	for y := year; y < year+limit; y++ {
		daysInMonth[1] = 28
		if isLeapYear(y) {
			daysInMonth[1] = 29
		}
		for m := 0; m < 12; m++ {
			if start == 6 { // sunday
				count++
			}
			remain := daysInMonth[m]%7
			start = (start+remain)%7
		}
	}
	return count
}

func main() {
	fmt.Println("Total Sundays of 1st of every month:", countSundays(1901, 1, 100))
}
