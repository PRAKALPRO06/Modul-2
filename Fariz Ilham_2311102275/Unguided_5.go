package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Berat() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()

		weights := strings.Fields(input)
		if len(weights) != 2 {
			fmt.Println("Input tidak valid.")
			continue
		}

		weight1, err1 := strconv.ParseFloat(weights[0], 64)
		weight2, err2 := strconv.ParseFloat(weights[1], 64)

		if err1 != nil || err2 != nil {
			fmt.Println("Input tidak valid.")
			continue
		}

		fmt.Print("Sepeda motor pak Andi akan oleng: ")

		difference := weight1 - weight2
		if difference < 0 {
			difference = -difference
		}

		if difference >= 9 {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}

		totalWeight := weight1 + weight2
		if totalWeight > 150 {
			fmt.Println("Proses selesai.")
			break
		}
	}
}

func main() {
	Berat()
}
