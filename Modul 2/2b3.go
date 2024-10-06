package main

import "fmt"

func main() {
        const (
                maxSelisihBerat = 9
                maxTotalBerat   = 150
        )

        var beratKiri, beratKanan float64
        var oleng bool

        for {
                fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
                fmt.Scan(&beratKiri, &beratKanan)

                if beratKiri < 0 || beratKanan < 0 {
                        fmt.Println("Berat tidak boleh negatif.")
                        break
                }

                if beratKiri+beratKanan > maxTotalBerat {
                        fmt.Println("Total berat melebihi batas.")
                        break
                }

                oleng = beratKiri-beratKanan >= maxSelisihBerat || beratKanan-beratKiri >= maxSelisihBerat
                fmt.Printf("Sepeda motor pak Andi akan oleng: %t\n", oleng)
        }

        fmt.Println("Proses selesai.")
}