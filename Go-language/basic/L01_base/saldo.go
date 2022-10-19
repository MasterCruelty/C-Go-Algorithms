package main

import(
	"fmt"
	"os"
	"strconv"
	)

func main(){
	saldo, _ := strconv.Atoi(os.Args[1])
	var spesa int
	for saldo > 0 {
		fmt.Println("Saldo rimanente: ",saldo)
		fmt.Scan(&spesa)
		saldo -= spesa
	}
	fmt.Println("Saldo finale:", saldo)
}
