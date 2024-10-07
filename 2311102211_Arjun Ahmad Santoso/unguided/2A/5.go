package main

import (
	"fmt"
)

func main() {

	array := [5]int{}
	array2 := [3]int{}
	fmt.Scanln( &array[0], &array[1], &array[2], &array[3], &array[4])
	fmt.Scanf("%c%c%c", &array2[0], &array2[1], &array2[2])
	fmt.Print("\n\n")
	for i := 0; i < 5; i++ {
		fmt.Printf("%c", array[i])
	}
	fmt.Print("\n")
	for i := 0; i < 3; i++ {
		fmt.Printf("%c", array2[i]+1)
	}
}