package main

import "fmt"

func main() {
    var r float64
    const pi = 3.1415926535

    fmt.Print("Masukkan jari-jari bola: ")
    fmt.Scanln(&r)

    
    r2 := r * r 
    r3 := r * r * r 

    volume := (4.0 / 3.0) * pi * r3

    luas := 4 * pi * r2

	fmt.Printf("Bola dengan jejari 5 memiliki volume %.2f dan luas kulit %.2f", volume, luas)
}
