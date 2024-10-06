package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for ulang := 1; ulang <= 2; ulang++ {
		yangHarusDiinputkan := []string{"merah", "kuning", "hijau", "ungu"}
		berhasil := true
		bacaInput := bufio.NewScanner(os.Stdin)
		
		fmt.Printf("=====================================\n")
		for i := 1; i <= 5; i++ {
			fmt.Printf("Percobaan %d: ", i)

			if !bacaInput.Scan() {
				fmt.Println("\nInput tidak cukup.")
				return
			}
			line := bacaInput.Text()
			warna := strings.Fields(line)

			if len(warna) != 4 {
				berhasil = false
				continue
			}
			for j, warna := range warna {
				if strings.ToLower(warna) != yangHarusDiinputkan[j] {
					berhasil = false
					break
				}
			}
		}
		fmt.Printf("BERHASIL: %v\n", berhasil)
	}
}