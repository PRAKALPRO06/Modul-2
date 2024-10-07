package main

import (
    "fmt"
    "math"
)

func main() {
    for {
        var kantong1, kantong2 float64

        fmt.Print("Masukan berat belanjaan di kedua kantong: ")
        fmt.Scan(&kantong1, &kantong2)

        if kantong1 < 0 || kantong2 < 0 {
            fmt.Println("Proses selesai.")
            break
        }

        if kantong1+kantong2 > 150 {
            fmt.Println("Proses selesai.")
            break
        }

        oleng := math.Abs(kantong1-kantong2) >= 9
        fmt.Printf("Sepeda motor Pak Andi akan oleng: %v\n", oleng)

        if kantong1 >= 9 || kantong2 >= 9 {
            fmt.Println("Proses selesai.")
            break
        }
    }
}