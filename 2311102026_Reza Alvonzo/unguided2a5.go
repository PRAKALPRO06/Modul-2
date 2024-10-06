package main

import (
	"bufio"
	"fmt"
	"os"
)

//reza alvonzo 2311102026 IF 11 06
func main() {
	var a, b, c, d, e int
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Masukkan 5 buah angka (nilai antara 32 hingga 127):")
	fmt.Scanf("%d %d %d %d %d", &a, &b, &c, &d, &e)

	fmt.Printf("Hasil karakter dari angka yang diinput: %c%c%c%c%c\n", a, b, c, d, e)

	fmt.Println("Masukkan karakter (contoh input: 123):")
	inputStr, _ := reader.ReadString('\n') 

	
	shiftedStr := ""
	for _, char := range inputStr {
		if char != '\n' && char != '\r' {
			shiftedStr += string(char + 1) 
		}
	}
	fmt.Printf("Hasil karakter dari inputan angka yang digeser: %s\n", shiftedStr)
}