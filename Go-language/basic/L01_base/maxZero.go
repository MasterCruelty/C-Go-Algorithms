package main

import ( "fmt"
)

//data una sequenza di interi terminata da -1, restituire il numero che al suo interno contiene più zeri
func main(){
	var result, n, max, count, save int
	for {
		fmt.Scan(&n)
		if n == -1{
			break
		}
		save = n
		for n != 0{
			if n == 0{
				break
			}
			if n % 10 == 0 {
				count++
				n = n / 10
			}else{
				n = n / 10
			}
		}
		if count > max{
			max = count
			result = save
		}
		count = 0
	}
	fmt.Println(result)
}
