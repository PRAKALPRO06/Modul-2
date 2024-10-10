package main

import (
    "fmt"
    "strings"
)

func main() {
    var pita string
    var jumlahBunga int
    var batasBunga int

    // Meminta pengguna untuk memasukkan batas jumlah bunga yang ingin ditambahkan
    fmt.Print("Masukkan jumlah bunga yang ingin ditambahkan: ")
    fmt.Scan(&batasBunga)

    // Proses input nama bunga sampai mencapai batas yang ditentukan
    for jumlahBunga < batasBunga {
        var bunga string
        fmt.Printf("Bunga %d: ", jumlahBunga+1)
        fmt.Scan(&bunga)

        // Jika input adalah 'SELESAI', keluar dari loop
        if strings.ToUpper(bunga) == "SELESAI" {
            break
        }

        // Gabungkan nama bunga dengan string pita menggunakan operator +
        if jumlahBunga == 0 {
            pita = bunga
        } else {
            pita += " - " + bunga
        }
        
        jumlahBunga++
    }

    // Menampilkan pita bunga dan jumlah total bunga yang dimasukkan
    fmt.Println("Pita:", pita)
    fmt.Println("Bunga:", jumlahBunga)
}
