package main

import "fmt"

func main() {

	berhenti := false
	var pita string
	pita_count := 0


	for berhenti == false {
		fmt.Printf("Bunga %d: ", pita_count + 1)
		var bunga string
		fmt.Scanln(&bunga)
		if bunga == "SELESAI" {
			break
		} else {
			pita_count++;
			pita += bunga + " - "
		}
	}
	fmt.Println("Pita: ", pita)
	fmt.Println("Bunga: ", pita_count)

}