package main
import "fmt"
func main() {
	var numero int
	for{
		fmt.Scan(&numero)
		if numero == -1{
			break
		}
		if numero > 100 {
			fmt.Println(numero)
			return
		}
	}
	fmt.Println("Nessun numero maggiore di 100")
}
