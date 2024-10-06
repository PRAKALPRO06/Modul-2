package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var bunga []string
	count := 0

	for {
		count++
		fmt.Printf("Bunga %d: ", count)
		scanner.Scan()
		input := scanner.Text()

		if strings.ToUpper(input) == "SELESAI" {
			break
		}

		bunga = append(bunga, input)
	}

	fmt.Print("\nPita: ")
	for _, nama := range bunga {
		fmt.Printf("%s - ", nama)
	}

	fmt.Printf("\nBunga: %d\n", len(bunga))
}
