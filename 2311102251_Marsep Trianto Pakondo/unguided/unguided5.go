package main

import (
	"fmt"
)

func main()  {
    var a, b, selisih float32
    var oleng, selesai bool
    var d float32 

    for !selesai {
        fmt.Printf("masukkan berat belanjaan di kedua kantong : ")
        fmt.Scanln(&a, &b)

        d = a + b

        if d <= 0 || d >= 150 {
            fmt.Println("Proses selesai.")
            selesai = true
        }else {
            selisih = a - b
            
            if selisih >= 9 || selisih <= -9 {
                oleng = true
            } else {
                oleng = false
            }
            fmt.Printf("Sepedah motor pak Andi akan oleng : %v\n", oleng)
        }
    }
}