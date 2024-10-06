package main

import (
	"fmt"
)

func main() {
	var berat1_2311102011, berat2 float64

	for {
		fmt.Print("masukkan berat belanjaan di kedua kantong = ")
		fmt.Scan(&berat1_2311102011, &berat2)

		if berat1_2311102011 >= 9 || berat2 >= 9 {
			fmt.Println("sepeda motor pak andi akan oleng : true")
		} else {
			fmt.Println("sepeda motor pak andi akan oleng : false")
		}
	}
}
