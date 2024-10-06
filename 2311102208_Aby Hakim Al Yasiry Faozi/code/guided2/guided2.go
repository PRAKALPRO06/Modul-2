package main

import (
	"fmt"
)

func isLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	}
	if year%100 == 0 {
		return false
	}
	if year%4 == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(isLeapYear(2016)) // Output: true
	fmt.Println(isLeapYear(2000)) // Output: true
	fmt.Println(isLeapYear(2018)) // Output: false
}
