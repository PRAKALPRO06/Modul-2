package main
import "fmt"
func main(){
	var tahun int
	
	fmt.Print("masukan tahun:")
	fmt.Scanln(&tahun)

	if (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0){
		fmt.Println("tahum kabisat : True")
	}else{
		fmt.Println("tahum kabisat : False")
	}
} 