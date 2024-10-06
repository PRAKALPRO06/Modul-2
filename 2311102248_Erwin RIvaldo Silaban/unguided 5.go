package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	totalBerat := 0.0

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		weights := strings.Split(input, " ")
		if len(weights) != 2 {
			fmt.Println("Masukkan dua angka dipisahkan oleh spasi.")
			continue
		}

		weight1, err1 := strconv.ParseFloat(weights[0], 64)
		weight2, err2 := strconv.ParseFloat(weights[1], 64)

		if err1 != nil || err2 != nil || weight1 < 0 || weight2 < 0 {
			fmt.Println("Masukkan dua bilangan real positif.")
			continue
		}

		totalBerat += weight1 + weight2
		selisihBerat := math.Abs(weight1 - weight2)

		if weight1 >= 9 || weight2 >= 9 || totalBerat > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		if selisihBerat >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}
	}
}
