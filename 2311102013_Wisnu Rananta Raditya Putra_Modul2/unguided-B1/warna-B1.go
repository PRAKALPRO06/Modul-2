package main

import (
    "fmt"
)

func main() {
    var warna_2311102013 [5][4]string
    urutanBenar_2311102013 := [4]string{"merah", "kuning", "hijau", "ungu"}
    berhasil_2311102013 := true

    for i := 0; i < 5; i++ {
        fmt.Printf("Masukkan warna untuk percobaan %d (pisahkan dengan spasi): ", i+1)
        fmt.Scan(&warna_2311102013[i][0], &warna_2311102013[i][1], &warna_2311102013[i][2], &warna_2311102013[i][3])
    }

    for i := 0; i < 5; i++ {
        fmt.Printf("percobaan %d: %s %s %s %s\n", i+1, warna_2311102013[i][0], warna_2311102013[i][1], warna_2311102013[i][2], warna_2311102013[i][3])
        if warna_2311102013[i] != urutanBenar_2311102013 {
            berhasil_2311102013 = false
        }
    }

    fmt.Printf("berhasil : %t\n", berhasil_2311102013)
}