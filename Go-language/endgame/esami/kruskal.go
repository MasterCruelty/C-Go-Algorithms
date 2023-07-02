//grafo rappresentato come lista di archi con pesi
type edges struct{
	vertex int
	edge int
}

type app struct{
	 vert int 
	 edge int 
	 w int
}

type graph struct{
	 g map[edges]int
	 vert int
}

//funzioni per operazioni UF
func find(i int, parent []int) int{
	for parent[i] != i{
		i = parent[i]
	}
	return i
}

func union(parent []int, i int, j int){
	a := find(i, parent)
	b := find(j, parent)
	parent[a] = b
}

func kruskalMTS(g *graph, parent []int){
	mincost := 0;

	//makeset per ogni vertice
	for i:=0; i < g.vert; i++{
	parent[i] = i 
	}

	weigths := make([]app, 0)
	for i,k := range g.g{
		var j app 
		j.vert = i.vertex 
		j.edge = i.edge 
		j.w = k
		weigths = append(weigths, j)
	}

	//ordino gli archi in modo non decrescente in base ai pesi
	sort.SliceStable(weigths, func (i, j int) bool {return weigths[i].w < weigths[j].w})


	edgeCount := 0
	min := -1
	for edgeCount < g.vert +1{
		a := find(weigths[edgeCount].vert, parent)
		b := find(weigths[edgeCount].edge, parent)
		if a != b {
			union(parent,a,b)
			min = weigths[edgeCount].w 
			fmt.Printf("Arco %d:(%d %d) costo: %d \n", edgeCount, weigths[edgeCount].vert, weigths[edgeCount].edge, min)
			mincost += min 
		}
		edgeCount++
	}
	fmt.Printf("\nCosto minimo: %d \n", mincost)
}
