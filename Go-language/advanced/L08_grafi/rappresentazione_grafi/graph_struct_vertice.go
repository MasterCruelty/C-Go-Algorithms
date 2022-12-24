package main

import "fmt"

type item int
type chiave string

//rappresentazione del grafo definendo una struct per i vertici che vontiene una slice di puntatori ai vertici vicini
type vertice struct {
	valore item
	k chiave
	adiacenti []*vertice // insieme dei vertici adiacenti
}

type grafo map[chiave]*vertice


func leggiGrafo() grafo {
	g := make(map[chiave]*vertice)
	i:=0
	var v1,v2 chiave
	_,err:=fmt.Scan(&v1,&v2)

	for err == nil {
		_,ok:=g[v1]
		if !ok {
			g[v1] = &vertice{item(i),v1,nil}
			i++
		}
		_,ok = g[v2]
		if !ok {
			g[v2] = &vertice{item(i),v2,nil}
			i++
		}
		g[v1].adiacenti = append(g[v1].adiacenti,g[v2])
		g[v2].adiacenti = append(g[v2].adiacenti,g[v1])
		_,err=fmt.Scan(&v1,&v2)
		if i > 5{
			break
		}
	}
	return g

}

func stampaVertice(v *vertice) {
	fmt.Print(v.k,", ",v.valore,", adiacenti: ")
	for _,el := range v.adiacenti {
		fmt.Print(el.k," ")
	}
	fmt.Println()
}

func stampaGrafo(g grafo) {
	fmt.Printf("Il grafo ha %d nodi\n",len(g))
	for _,vertice := range g {
		stampaVertice(vertice)
	}
	fmt.Println()
}

func main() {
	g := leggiGrafo()
	stampaGrafo(g)
}
