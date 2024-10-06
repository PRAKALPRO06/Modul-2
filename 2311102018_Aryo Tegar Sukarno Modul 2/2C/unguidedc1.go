package main

import "fmt"

func main() {
        var beratParsel int
        fmt.Print("Berat parsel (gram): ")
        fmt.Scan(&beratParsel)

        kg := beratParsel / 1000
        gram := beratParsel % 1000
        
        biayaDasar := kg * 10000
   
        var biayaTambahan int
        if gram >= 500 {
                biayaTambahan = gram * 5
        } else {
                biayaTambahan = gram * 15
        }

      
        if kg > 10 {
                biayaTambahan = 0
        }

       
        totalBiaya := biayaDasar + biayaTambahan
        
        fmt.Printf("Detail berat: %d kg + %d gr\n", kg, gram)
        fmt.Printf("Detail biaya: Rp %d + Rp %d\n", biayaDasar, biayaTambahan)
        fmt.Printf("Total biaya: Rp %d\n", totalBiaya)
}