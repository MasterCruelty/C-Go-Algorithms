package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
 * C'è una rete con numero arbitrario n di interruttori. Ogni interruttore è collegato con il successivo e il precedente.
 * L'ultimo è collegato con il primo e il penultimo, il primo è collegato con l'ultimo e il secondo.
 *  
*/





//slice di bool per rappresentare lo stato d'accensione
type conf []bool

//main di test per le funzioni
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

//preme l'interruttore numero x e restituisce lo stato d'accensione aggiornato
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

//Preme in sequenza gli interruttori della slice x, restituisce lo stato d'accensione aggiornato
func sequenza(curr conf, x []int) conf {

	for i := 0; i < len(x); i++ {
		curr = premi(curr, x[i])
	}

	return curr
}

//Restituisce la sequenza di bottoni da premere per spegnere tutte le luci
func sequenzaSpegniTutto(curr conf) []int{
	var coda []conf = []conf{curr}
	var stato string
	percorso := make(map[string][]int)
	for !is_spenta(curr) {
		for i := 1;i<=len(curr);i++ {
			coda = append(coda, premi(curr,i))
			stato = trasformaInStringa(premi(curr,i))
			percorso[stato] = append(percorso[stato], i)
		}
		curr = coda[1]
		coda = coda[1:]
	}
	return percorso[trasformaInStringa(curr)]
}

//restituisce tutti gli stati raggiungibili premendo un bottone a partire dallo stato curr
func raggiungibili(curr conf) []conf{
	raggiungibili := make([]conf,0)
	for i:=1;i<=len(curr)-1;i++{
		s := premi(curr,i)
		raggiungibili = append(raggiungibili,s)
	}
	return raggiungibili

}

//Conversione di slice di bool in stringa di zeri e uni
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

//controllo se ogni bit è a zero nello stato d'accensione
func is_spenta(curr conf) bool {

	for _, v := range curr {
		if v {
			return false
		}
	}
	return true
}

//Restituisce il numero di bottoni che è necessario premere per spegnere tutte le luci
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

