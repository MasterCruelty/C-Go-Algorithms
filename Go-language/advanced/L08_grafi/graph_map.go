package main

import "fmt"


//mappa che associa ad ogni vertice(identificato da una stringa)
//l'insieme dei vertici(stringhe) a lui adiacenti
type grafo map[string][]string

func nuovoGrafo(n int) grafo {
	g := make(map[string][]string)
	return g
}

//lettura del grafo da standard input, si considera grafo non orientato
func leggiGrafo() grafo {
	g := make(map[string][]string)
	fmt.Println("Inserisci gli archi, uno per riga, nel formato 'v1 v2'")
	var v1,v2 string
	_,err := fmt.Scan(&v2,&v2)

	for err == nil {
		//leggo i vertici
		g[v1] = append(g[v1],v2)
		g[v2] = append(g[v2],v1)
		stampaGrafo(g)
		_,err = fmt.Scan(&v1,&v2)
	}
	return g
}


//stampa del grafo
func stampaGrafo(g grafo) {
	fmt.Printf("Il grafo ha %d nodi\n",len(g))
	for vertice,lista := range g{
		fmt.Println(vertice, ":",lista)
	}
}

//main per testare grafo implementato come mappa di stringhe
func main() {
	g := leggiGrafo()
	stampaGrafo(g)
}
