package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var pita string
	var jumlahBunga int

	for {
		fmt.Print("Jumlah Bunga : ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "SELESAI" {
			fmt.Println("Pita:", pita)
			fmt.Printf("Bunga: %d\n", jumlahBunga)
			return
		}

		n, err := strconv.Atoi(input)
		if err != nil || n <= 0 {
			fmt.Println("Masukkan Jumlah Bunga : ")
			continue
		}

		for i := 1; i <= n; i++ {
			fmt.Printf("Bunga %d: ", i)
			bunga, _ := reader.ReadString('\n')
			bunga = strings.TrimSpace(bunga)

			if bunga == "SELESAI" {
				fmt.Println("Pita:", pita)
				fmt.Printf("Bunga: %d\n", jumlahBunga)
				return
			}

			if pita == "" {
				pita = bunga
			} else {
				pita += " - " + bunga
			}
			jumlahBunga++
		}

		fmt.Println("Pita:", pita)
		fmt.Printf("Bunga: %d\n", jumlahBunga)
	}
}
