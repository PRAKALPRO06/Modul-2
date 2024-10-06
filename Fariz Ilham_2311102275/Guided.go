package main

import (
	"fmt"
	"math"
)

// fungsi untuk menukar 3 string
func main() {
	var (
		satu, dua, tiga string
		temp            string
	)
	fmt.Println("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Println("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Println("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)

	fmt.Println()
	cek_Tahun_kabisat()
}

// fungsi untuk memeriksa tahun kabisat
func cek_Tahun_kabisat() {
	var tahun int
	fmt.Println("Masukkan sebuah tahun: ")
	fmt.Scanln(&tahun)
	fmt.Println()

	if (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0) {
		fmt.Println(tahun, "adalah tahun kabiat: true")
	} else {
		fmt.Println(tahun, "bukan tahun kabisat: false")
	}

	fmt.Println()
	hitung_bola()
}

// fungsi untuk menghitung volume dan luas permukaan bola
func hitung_bola() {
	var jarijari float64
	fmt.Println("Masukkan jari-jari bola: ")
	fmt.Scanln(&jarijari)

	volume := (4.0 / 3.0) * math.Pi * math.Pow(jarijari, 3)
	luas_permukaan := 4 * math.Pi * math.Pow(jarijari, 2)

	fmt.Println("Volume bola dengan jari-jari %.2f adalah %.2f\n", jarijari, volume)
	fmt.Println()
	fmt.Println("luas_permukaan bola dengan jari-jari %.2f adalah %.2f\n", jarijari, luas_permukaan)
}
