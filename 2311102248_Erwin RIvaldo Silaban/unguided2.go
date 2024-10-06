
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Masukkan 5 data integer (32-127) dipisahkan oleh spasi:")
	intLine, _ := reader.ReadString('\n')
	intStrs := strings.Fields(strings.TrimSpace(intLine))

	if len(intStrs) != 5 {
		fmt.Println("Error: Harap masukkan tepat 5 angka.")
		return
	}

	var intChars string
	for _, s := range intStrs {
		num, err := strconv.Atoi(s)
		if err != nil || num < 32 || num > 127 {
			fmt.Println("Error: Masukkan harus berupa angka antara 32 dan 127.")
			return
		}
		intChars += string(rune(num))
	}

	fmt.Println("Masukkan 3 karakter (tanpa spasi):")
	charLine, _ := reader.ReadString('\n')
	chars := strings.TrimSpace(charLine)

	if len(chars) != 3 {
		fmt.Println("Error: Harap masukkan tepat 3 karakter.")
		return
	}

	fmt.Println("Keluaran:")
	fmt.Println(intChars)

	var nextChars string
	for _, ch := range chars {
		nextChars += string(ch + 1)
	}
	fmt.Println(nextChars)
}
