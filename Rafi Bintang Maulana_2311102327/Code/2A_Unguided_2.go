package main
import (
	"fmt"
)
func main() {
	var BilanganInteger [5]int
	var Karakter string
	fmt.Println("===============================================")
	fmt.Println("Masukkan 5 data integer (antara 32 sampai 127):")
	for i := 0; i < 5; i++ {
		fmt.Scan(&BilanganInteger[i])
	}
	fmt.Println("Masukkan 3 karakter:")
	fmt.Scan(&Karakter)
	fmt.Println("===============================================")
	fmt.Print("Hasil karakter dari data integer: ")
	for _, value := range BilanganInteger {
		if value >= 32 && value <= 127 {
			fmt.Printf("%c", value)
		}
	}
	fmt.Println()

	if len(Karakter) == 3 {
		fmt.Print("3 karakter berikutnya: ")
		for _, ch := range Karakter {
			fmt.Printf("%c", ch+1) 
		}
	fmt.Println("\n===============================================")
	} else {
		fmt.Println("Input harus terdiri dari tepat 3 karakter.")
	}
}