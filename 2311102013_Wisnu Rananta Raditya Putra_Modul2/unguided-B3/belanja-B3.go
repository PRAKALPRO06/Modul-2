	package main

	import "fmt"
	import "math"

	func main() {
			var berat1_2311102013, berat2_2311102013 float64

			for {
					fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
					fmt.Scan(&berat1_2311102013, &berat2_2311102013)

					if berat1_2311102013 < 0 || berat2_2311102013 < 0 {
							fmt.Println("Berat tidak boleh negatif. Proses selesai.")
							break
					}

					if berat1_2311102013+berat2_2311102013 > 150 {
							fmt.Println("Total berat melebihi 150 kg. Proses selesai.")
							break
					}

					selisih := berat1_2311102013 - berat2_2311102013
					if math.Abs(selisih) >= 9 {
							fmt.Println("Sepeda motor Pak Andi akan oleng: true")
					} else {
							fmt.Println("Sepeda motor Pak Andi akan oleng: false")
					}
			}
	}