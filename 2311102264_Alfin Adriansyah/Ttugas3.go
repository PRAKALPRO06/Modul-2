package main

import "fmt"

func main() {

	var urutanBenar = [4]string{"merah", "kuning", "hijau", "ungu"}

	var hasilPercobaan [5][4]string
	berhasil := true

	for i := 0; i < 5; i++ {
		fmt.Printf("Percobaan %d: ", i+1)
		for j := 0; j < 4; j++ {
			fmt.Scan(&hasilPercobaan[i][j])
		}
	}
	for i := 0; i < 5; i++ {
		for j := 0; j < 4; j++ {
			if hasilPercobaan[i][j] != urutanBenar[j] {
				berhasil = false
				break
			}
		}
		if !berhasil {
			break
		}
	}

	fmt.Printf("BERHASIL: %v\n", berhasil)
}
