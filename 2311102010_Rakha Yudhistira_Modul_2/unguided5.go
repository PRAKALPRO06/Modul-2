package main

import (
	"fmt"
)

func main() {
	var kantongKanan, KantongKiri float64

	for {
		fmt.Print("masukkan berat belanjaan di kedua kantong = ")
		fmt.Scan(&kantongKanan, &KantongKiri)

		if kantongKanan >= 9 || KantongKiri >= 9 {
			fmt.Println("sepeda motor pak andi akan oleng : true")
			break
		} else {
			fmt.Println("sepeda motor pak andi akan oleng : false")
		}
	}
}