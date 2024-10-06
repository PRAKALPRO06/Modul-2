package main
import (
	"fmt"
)

func main() {
	var beratBebanKiri, beratBebanKanan float64

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&beratBebanKiri, &beratBebanKanan)

		totalBerat := beratBebanKiri + beratBebanKanan
		if beratBebanKiri < 0 || beratBebanKanan < 0 || totalBerat > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		selisih := beratBebanKiri - beratBebanKanan
		if selisih < 0 {
			selisih = -selisih
		}

		if selisih >= 9 {
			fmt.Println("Sepeda motor Pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor Pak Andi akan oleng: false")
		}
	}
}