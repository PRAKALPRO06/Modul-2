// 2311102037 BRIAN FARREL EVANDHIKA IF 11 06

package main

import (
	"fmt"
)

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