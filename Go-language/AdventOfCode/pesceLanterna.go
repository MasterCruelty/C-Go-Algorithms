package main

import( "fmt"
	"strings"
	"io/ioutil"
	"strconv"
)


//La traccia per questo esercizio è reperibile da qui: https://www.adventofcode.com/2021/day/6
func main(){
	content, err := ioutil.ReadFile("inputPesci.txt")
	if err != nil{
		panic(err)
	}
	input := strings.Split(string(content),",")
	var timer = []int{}
	fmt.Print("[")
	for _, i := range input {
		if i == "\n" {
			break
		}
		n,err := strconv.Atoi(i)
		if err != nil{
			panic(err)
		}
		fmt.Print(" ",n)
		timer = append(timer,n)
	}
	fmt.Print("]\n")
	gg := 1
	for gg < 80{
		for i:= 0;i < len(timer);i++{
			if timer[i] == 0{
				timer = append(timer,8)
				timer[i] = 6
			} else{
				timer[i] -= 1
			}
		}
		gg += 1
		if gg == 18{
			fmt.Println("Dopo 18 giorni ci sono il seguente numero di pesci: ",len(timer))
		}

	}
	fmt.Println("Dopo 80 giorni ci sono il seguente numero di pesci: ",len(timer))
}
