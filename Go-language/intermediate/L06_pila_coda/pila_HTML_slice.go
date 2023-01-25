package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)


//definizione struttura dati pila mediante slice
type stack struct {
	items []string
}

//inserisce un elemento in cima alla pila
func (s *stack) push(item string){
	s.items = append(s.items,item)
}

//preleva l'elemento in cima alla pila e lo elimina
func (s *stack) pop() (string,error) {
	if len(s.items) == 0 {
		return "",fmt.Errorf("Pila vuota")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

//controlla se la pila è vuota
func (s stack) isEmpty() bool {
	return len(s.items) == 0
}

/*
 * Prende una stringa in input che rappresenta un documento HTML scritto su una riga e utilizza una pila
 * per verificare se il documento è ben formattato. La funzione esegue la seguente procedura:
 * Crea una nuova istanza della struttura pila.
 * Utilizza strings.Split per dividere il documento in una serie di tag.
 * Con un ciclo for viene analizzato ogni tag della serie.
 * Se il tag non contiene "/" allora è un tag di apertura, viene aggiunto alla pila con push().
 * Se invee contiene "/" allora è un tag di chiusura. Viene rimosso l'ultimo tag di apertura dalla pila con pop()
 * Se il tag di apertura rimosso non corrisponde al tag di chiusura, c'è un errore e viene restituito false e la posizione.
 * Alla fine del ciclo se la pila non è vuota, allora ci sono tag di apertura non corrispondenti, restituisco false e una lista di non chiusi.
 * altrimenti restituisco true
*/
func checkTags(doc string) (bool,int,[]string) {
	var stack stack = stack{nil}
	tags := strings.Split(doc," ")
	for i,t := range tags {
		if !strings.Contains(t,"/"){
			stack.push(t)
		} else if strings.Contains(t,"/") {
			openingTag,err:= stack.pop()
			if err != nil {
				return false,i+1,nil
			}
			if openingTag[1:] != t[2:] {
				return false,i+1,nil
			}
		}
	}
	if stack.isEmpty() {
		return true,-1,nil
	}else{
		var tag_rimasti = make([]string,0)
		for !stack.isEmpty() {
			item,err:= stack.pop()
			if err != nil{
				tag_rimasti = append(tag_rimasti,item)
			}
		}
		return false,-1,tag_rimasti
	}
}

//main di test per analizzare documenti HTML su riga singola (usare i tag senza <> ad esempio: "a b /a /b")
func main() {
	fmt.Print("Inserisci documento html(una riga): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	doc := scanner.Text()
	benFormato,errPos,tag_rimasti := checkTags(doc)
	if benFormato {
		fmt.Println("Il documento è ben formato")
	} else{
		fmt.Println("Il documento NON è ben formato")
		if len(tag_rimasti) > 0 {
			fmt.Println("\ttag non chiuso:",tag_rimasti)
		} else{
			fmt.Println("\tPosizione errore:",errPos)
		}
	}
}
