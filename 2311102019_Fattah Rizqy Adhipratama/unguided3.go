package main

import "fmt"

func main() {
	maxF := 100
	f0 := 0
	f1 := 1
	f2 := 1

	fmt.Println("Bilangan pertama:", f1)

	for f2 <= maxF {
		f0 = f1
		f1 = f2
		f2 = f1 + f0
		fmt.Println("Bilangan berikutnya:", f1)
	}
}