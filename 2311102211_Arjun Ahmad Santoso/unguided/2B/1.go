package main

import "fmt"

func main() {

	berhasil := true
	array := [4] string {}
	for i := 1; i<=5; i++ {
		fmt.Printf("Percobaan %d: ", i)
		for j := 0; j<4; j++ {
			fmt.Scan(&array[j])
		}
		if (array[0] != "merah"|| 
		array[1] != "kuning"|| 
		array[2] != "hijau"|| 
		array[3] != "ungu" ) {
			berhasil = false
		}
	}
	fmt.Println("Berhasil: ", berhasil)

}