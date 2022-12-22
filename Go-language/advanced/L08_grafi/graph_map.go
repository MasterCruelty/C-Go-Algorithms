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
	i := 0
	for err == nil {
		//leggo i vertici
		g[v1] = append(g[v1],v2)
		g[v2] = append(g[v2],v1)
		//stampaGrafo(g)
		_,err = fmt.Scan(&v1,&v2)
		i++
		if i > 6{
			break
		}
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

//Restituisce il grado del vertice v
//Si ricorda che il grado di un vertice è definito come il numero di vertici ad esso adiacenti.
func degree(g grafo,v string) {
	for vertice,lista : range g{
		if vertice == v{
			fmt.Printf("Il grado del vertice %s è %d",v,len(lista))
			break
		}
	}
}

//funzione ausiliaria per controllare se un vertice è nella lista di adiacenza
func verticeInLista(w string,list []string) bool{
	for _,b := range list{
		if b == w{
			return true
		}
	}
	return false
}

//Controlla se esiste un cammino semplice tra il vertice v e il vertice w
//Si ricorda che un cammino si dice semplice quando attraversa ogni vertice al più una volta.
func path(g grafo,v string,w string){
	for vertice,lista: range g{
		if vertice ==v && verticeInlista(w,lista){
			fmt.Printf("Esiste un cammino semplice tra %s e %s",v,w)
			break
	}
}

//visita in profondità
//la funzione fa uso di una struttura di supporto aux che serve a ricordare quali vertici
//sono già stati visitati. Usiamo una mappa da string a bool perché i vertici sono identificati da stringhe
//prima di invocare la funzione, aux dovrà essere allocata e inizializzata.
func dfs1(g grafo, v string,aux map[string]bool) {
	fmt.Println(v)
	aux[v]=true
	for _, v2 := range g[v]{
		if aux[v2] != true {
			dfs1(g,v2,aux)
		}
	}
}

//visita in ampiezza
//Usiamo una coda con una slice di stringhe. Inseriamo(enquee) in fondo con append e estraiamo(dequeue) dall'inizio con [1:]
//La visita avviene quando si estraggono gli elementi dalla coda.Quando un vertice viene inserito nella coda si marca come visitato
func bfs1(g grafo,v string,aux map[string]bool){
	coda := []string{v}
	aux[v] = true
	
	for len(coda) > 0 {
		v := coda[0]
		coda = coda[1:]
		fmt.Println("\t",v)

		for _,v2 := range g[v] {
			if !aux[v2] {
				coda = append(coda,v2)
				aux[v2] = true
			}
		}
	}
}



//main per testare grafo implementato come mappa di stringhe
func main() {
	g := leggiGrafo()
	stampaGrafo(g)
	//test visita in profondità
	fmt.Println("Visita in profondità")
	aux := make(map[string]bool)
	dfs1(g,"anna",aux)
	//test visita in ampiezza
	fmt.Println("Visita in ampiezza")
	aux = make(map[string]bool)
	bfs1(g,"anna",aux)
}
