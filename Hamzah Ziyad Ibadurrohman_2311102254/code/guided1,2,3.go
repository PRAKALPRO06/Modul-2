package main

import (
	"fmt"
	"math"
)

func hitungbola() {
	var jarijari float64
	fmt.Print("Masukkan jari-jari bola : ")
	fmt.Scanln(&jarijari)

	volume := (4.0 / 3.0) * math.Pi * math.Pow(jarijari, 3)
	luasPermukaan := 4 * math.Pi * math.Pow(jarijari, 2)

	fmt.Printf("Volume bola dengan jari-jari %.2f adalah %.2f\n", jarijari, volume)
	fmt.Printf("Luas permukaan bola dengan jari-jari %.2f adalah %.2f\n", jarijari, luasPermukaan)
}

func cekTahunKabisat() {
	var tahun int
	fmt.Print("Masukkan sebuah tahun: ")
	fmt.Scanln(&tahun)

	if (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0) {
		fmt.Println(tahun, "adalah tahun kabisat: true")
	} else {
		fmt.Println(tahun, "bukan tahun kabisat: false")
	}
}

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)
	fmt.Println("Masukkan Input string: ")
	fmt.Scanln(&satu)
	fmt.Println("Masukkan Input string: ")
	fmt.Scanln(&dua)
	fmt.Println("Masukkan Input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("output awal =" + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("output akhir =" + satu + " " + dua + " " + tiga)

	cekTahunKabisat()

	hitungbola()
}
