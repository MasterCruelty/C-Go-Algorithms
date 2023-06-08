package main

import (
	"fmt"
)



func main() {
	var n,m int
	fmt.Scan(&n,&m)
	r := make([]int,0)
	input := 0
	for i:=0;i<=n;i++{
		fmt.Scan(&input)
		r = append(r,input)
	}
	//calcolo il numero di incrementi su tutte le n rilevazioni
	first := r[0]
	c1 := 0
	for i:=1;i<len(r);i++{
		if r[i] > first{
			c1++
		}
		first = r[i]
	}
	//calcolo il numero di incrementi in una finestra di m rilevazioni
	p := m-1
	c2 := 0
	for k:=0;k<= n-m;k++{
		if r[p] > r[k]{
			c2++
		}
		p++
	}
	fmt.Print(c1,c2)
	fmt.Println()
}
