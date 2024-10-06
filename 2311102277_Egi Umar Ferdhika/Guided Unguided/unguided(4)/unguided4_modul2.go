package main

import (
	"fmt"
)

func main() {
	ribbon := ""
	flowerCount := 0

	for {
		flowerCount++
		var flower string
		fmt.Printf("Bunga %d: ", flowerCount)
		fmt.Scan(&flower)

		if flower == "selesai" {
			break
		}

		if flowerCount == 1 {
			ribbon = flower
		} else {
			ribbon = ribbon + " - " + flower
		}
	}
	fmt.Println("Pita:", ribbon)
}
