package main

import "fmt"

func main() {
	var x,y int

	fmt.Print("primo numero: ")
	fmt.Scan(&x)
	fmt.Print("secondo numero: ")
	fmt.Scan(&y)

	fmt.Println(multiply(x,y))
}

func multiply(x,y int) int{
	if y == 1 {
		return x
	} else {
		return x + multiply(x,y-1)
	}
}
