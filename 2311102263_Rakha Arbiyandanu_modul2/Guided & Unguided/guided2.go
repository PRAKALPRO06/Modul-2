package main
import "fmt"

func main(){

	var tahun int
	
	fmt.Print("Masukkan tahun : ")
	fmt.Scanln(&tahun)

	if tahun%400 == 0 {
		fmt.Println("Tahun ini kabisat")
	} else if tahun%100 == 0 {
		fmt.Println("Ini bukan tahun kabisat")
	} else if tahun%4 == 0 {
		fmt.Println("Tahun ini kabisat")
	} else {
		fmt.Print("Tidak terdifinisi")
	}
}
