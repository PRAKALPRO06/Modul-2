package main

import "fmt"

func main() {
        var celsius float64

        fmt.Print("Masukkan suhu dalam Celcius: ")
        fmt.Scanln(&celsius)

        // Konversi ke Fahrenheit
        fahrenheit := (celsius * 9/5) + 32

        // Konversi ke Reamur
        reamur := celsius * 4 / 5

        // Konversi ke Kelvin
        kelvin := celsius + 273.15

        fmt.Printf("Suhu dalam Fahrenheit: %.2f\n", fahrenheit)
        fmt.Printf("Suhu dalam Reamur: %.2f\n", reamur)
        fmt.Printf("Suhu dalam Kelvin: %.2f\n", kelvin)
}