package main

import (
        "fmt"
        "strings"
)

func main() {
        var jumlahBunga int
        var bunga []string
        var pita string

        for {
                fmt.Print("Masukkan jumlah bunga: ")
                fmt.Scanln(&jumlahBunga)

                if jumlahBunga <= 0 {
                        fmt.Println("Jumlah bunga harus lebih dari 0")
                        continue
                }

                bunga = make([]string, jumlahBunga)
                for i := 0; i < jumlahBunga; i++ {
                        fmt.Printf("Bunga %d: ", i+1)
                        fmt.Scanln(&bunga[i])
                }

                pita = strings.Join(bunga, "-")
                fmt.Println("Pita:", pita)

                var lanjut string
                fmt.Print("Lanjutkan? (ya/tidak): ")
                fmt.Scanln(&lanjut)

                if strings.ToLower(lanjut) != "ya" {
                        break
                }
        }
}