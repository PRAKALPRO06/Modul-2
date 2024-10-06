package main

import (
	"fmt"
)


//reza alvonzo 2311102026 IF 11 06
func main() {
	var berat1, berat2 float64

	for {
		fmt.Print("masukkan berat belanjaan di kedua kantong = ")
		fmt.Scan(&berat1, &berat2)

		if berat1 >= 9 || berat2 >= 9 {
			fmt.Println("sepeda motor pak andi akan oleng : true")
		} else {
			fmt.Println("sepeda motor pak andi akan oleng : false")
		}
	}
}
