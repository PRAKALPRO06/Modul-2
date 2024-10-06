package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	pita := ""
	BunganyaAda := 0 

	i := 1
	for {
		fmt.Printf("Bunga %d: ", i)
		scanner.Scan()
		bunga := scanner.Text()

		if strings.ToUpper(bunga) == "SELESAI" {
			break
		}

		pita += bunga + " - "
		BunganyaAda++
		i++
	}

	if pita == "" {
		fmt.Println("Pita: ")
	} else {
		fmt.Printf("Pita: %s\n", strings.TrimSpace(pita))
	}

	fmt.Printf("Bunga: %d\n", BunganyaAda)
}