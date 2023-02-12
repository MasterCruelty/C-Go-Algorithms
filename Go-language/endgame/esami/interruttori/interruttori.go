package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type conf []bool

func main() {

	var rete, interuttori []int
	var st_rete conf
	var n int
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		riga := scanner.Text()
		valori := strings.Split(riga, " ")
		switch valori[0] {
		case "f":
			os.Exit(0)
		case "+":
			n, _ = strconv.Atoi(valori[1])
			for i := 0; i < n; i++ {
				rete = append(rete, 0)
				st_rete = append(st_rete, false)
			}
		case "o":
			for j := 0; j < n; j++ {
				rete[j], _ = strconv.Atoi(string(valori[1][j]))
				if rete[j] == 0 {
					st_rete[j] = false
				} else {
					st_rete[j] = true
				}
			}
		case "p":
			fmt.Println(rete)
		case "x":
			fmt.Println(spegniTutto(st_rete))
		case "s":
			v, _ := strconv.Atoi(valori[1])
			st_rete = premi(st_rete, v)
			rete = passaggio(st_rete)
		case "S":
			for k, v := range valori {
				if k != 0 {
					val, _ := strconv.Atoi(v)
					interuttori = append(interuttori, val)
				}
			}
			st_rete = sequenza(st_rete, interuttori)
			rete = passaggio(st_rete)
		}
	}
}

func premi(curr conf, x int) conf {
	temp_conf := make(conf,len(curr))
	copy(temp_conf,curr)
	if x == 1 {
		temp_conf[x-1] = !temp_conf[x-1]
		temp_conf[x] = !temp_conf[x]
		temp_conf[len(temp_conf)-1] = !temp_conf[len(temp_conf)-1]
	} else if x == len(temp_conf) {
		temp_conf[x-1] = !temp_conf[x-1]
		temp_conf[0] = !temp_conf[0]
		temp_conf[x-2] = !temp_conf[x-2]
	} else {
		temp_conf[x-1] = !temp_conf[x-1]
		temp_conf[x] = !temp_conf[x]
		temp_conf[x-2] = !temp_conf[x-2]
	}

	return temp_conf
}

func sequenza(curr conf, x []int) conf {

	for i := 0; i < len(x); i++ {
		curr = premi(curr, x[i])
	}

	return curr
}

func sequenzaSpegniTutto(curr conf) []int {

	var coda []conf = []conf{curr}
	var stato string
	percorso := make(map[string][]int)
	for !is_spenta(curr) {
		for i := 1; i <= len(curr); i++ {
			coda = append(coda, premi(curr, i))
			stato = trasformaInStringa(premi(curr,i))
			percorso[stato] = append(percorso[stato], i)
		}
		curr = coda[1]
		coda = coda[1:]
	}

	return percorso[trasformaInStringa(curr)]
}

func trasformaInStringa(curr conf) string {

	var s string
	for _, v := range curr {
		if !v {
			s += "0"
		} else {
			s += "1"
		}
	}
	return s
}

func is_spenta(curr conf) bool {

	for _, v := range curr {
		if v {
			return false
		}
	}
	return true
}

func spegniTutto(curr conf) int {
	return len(sequenzaSpegniTutto(curr))
}

func passaggio(curr conf) []int {

	var stato []int
	for k, v := range curr {
		if v {
			stato[k] = 1
		} else {
			stato[k] = 0
		}
	}
	return stato
}

