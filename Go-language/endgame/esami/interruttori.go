package main

import "fmt"

type arcoUscente struct{
   node []bool
   bottone int
}

type grafo struct{
   n int  //numero nodi
   m int //numero archi
   adiacenti  [][]arcoUscente
}


func BFSWithGoal(g grafo, start int, goal int) []int {
  queue := []int{start}
  visited := make([]bool, g.n)
  visited[start] = true
  var result []int

  for len(queue) > 0 {
    curr := queue[0]
    queue = queue[1:]
    result = append(result, curr)

    if curr == goal {
      return result
    }

    for _, arco := range g.adiacenti[curr] {
      if !visited[arco.bottone] {
        visited[arco.bottone] = true
        queue = append(queue, arco.bottone)
      }
    }
  }

  return result
}

func creaGrafo(n int) grafo {
   g := grafo{n: n, m: 0, adiacenti: make([][]arcoUscente, n)}
   for i := 0; i < n; i++ {
      var adiacenti []arcoUscente
      if i == 0 {
         adiacenti = append(adiacenti, arcoUscente{node: []bool{true, false, false, false, false, false, false}, bottone: (i+1)%n})
         adiacenti = append(adiacenti, arcoUscente{node: []bool{true, true, false, false, false, false, false}, bottone: (n+i-1)%n})
      } else if i == n-1 {
         adiacenti = append(adiacenti, arcoUscente{node: []bool{false, false, false, false, false, false, true}, bottone: (i+1)%n})
         adiacenti = append(adiacenti, arcoUscente{node: []bool{false, false, false, false, false, true, true}, bottone: (n+i-1)%n})
      } else {
         adiacenti = append(adiacenti, arcoUscente{node: []bool{false, false, false, false, false, false, false}, bottone: (i + 1) % n})
         adiacenti = append(adiacenti, arcoUscente{node: []bool{false, false, false, false, false, false, false}, bottone: (n + i - 1) % n})
         //adiacenti[i][0].node[i] = true
         //adiacenti[i][1].node[n-i-1] = true
      }
      g.adiacenti[i] = adiacenti
   }
   return g
}

func main(){
	// Crea il grafo
	g := creaGrafo(7)
	// Usa la visita in ampiezza per trovare il cammino minimo
	result := BFSWithGoal(g, 0, 6)
	fmt.Println("Cammino minimo: ", result)
}
