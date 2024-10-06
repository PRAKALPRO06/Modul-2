package main

import "fmt"

func main() {
    var colors [5][4]string
    var allTrue = true

    for i := 0; i < 5; i++ {
        fmt.Printf("Masukan Warnanya %d:\n", i+1)
        for j := 0; j < 4; j++ {
            fmt.Scan(&colors[i][j])
        }
    }

    for i := 0; i < 5; i++ {
        if colors[i][0] != "Merah" || colors[i][1] != "Kuning" || colors[i][2] != "Hijau" || colors[i][3] != "Ungu" {
            allTrue = false
            break
        }
    }

    if allTrue {
        fmt.Println("Berhasil: true")
    } else {
        fmt.Println("Berhasil: false")
    }
}
