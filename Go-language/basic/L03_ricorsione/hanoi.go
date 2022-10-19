package main

import "fmt"

func main() {
	var n int
	var mosse uint64
	var risp string
	fmt.Print("Inserire il numero di dischi: ")
	fmt.Scan(&n)
	fmt.Print("Vuoi solo il numero di mosse? s/n")
	fmt.Scan(&risp)
	if risp == "n"{
		hanoi(n,string('0'),string('2'),string('1'))
	}
	mosse = contaMosse(n)
	fmt.Println("Numero mosse: ",mosse)
}

func hanoi(n int,from,to,temp string) {
	if n == 1{
		fmt.Printf("Sposto 1 da paletto %s a paletto %s\n",from,to)
		return
	}
	hanoi(n-1,from,temp,to)
	fmt.Printf("Sposto %d da paletto %s a paletto %s\n",n,from,to)
	hanoi(n-1,temp,to,from)
}

func contaMosse(n int) uint64 {
	if n == 0{
		return 0
	}
	return 2 * contaMosse(n-1) +1
}
