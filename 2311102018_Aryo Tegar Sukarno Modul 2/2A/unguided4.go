package main 

import "fmt"

func main(){
	var celcius float64

	fmt.Print("temperatur Celcius: ")
	fmt.Scan(&celcius)

	fahrenheit :=(celcius * 9/5) + 32

	reamur := celcius * 4/5

	kelvin := celcius + 273.15

	fmt.Printf("Temperatur Celcius : %.0f\n",celcius)
	fmt.Printf("Temperatur Celcius : %.0f\n",reamur)
	fmt.Printf("Temperatur Celcius : %.0f\n",fahrenheit)
	fmt.Printf("Temperatur Celcius : %.0f\n",kelvin)
}