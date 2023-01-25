package main

import "fmt"

/*
 * Intervallo pesato definito come tripla di interi (i,f,v)
 * i = tempo di inizio
 * f = tempo di fine
 * v = valore/peso dell'intervallo.
 * Dato I insieme di intervalli, una soluzione è data da un sottoinsieme S contenuto in I di intervalli non sovrapposti tra loro.
 * Il valore di una soluzione S è dato dalla somma dei valori degli intervalli S.
 * 
 *
 *
 * Step 1: ordinare gli elementi di i in base al tempo di fine così da avere I= {(i1,f1,v1),(i2,f2,v2),...,(in,fn,vn)} con f1<=f2<=..<=fn.
 * Step 2: per ogni indice j tra 1 e n, definire p(j) come il più grande indice i < j  tale che l'intervallo di indice i non si sovrappone all'intervallo di indice j.
 * Step 3: Definire Opt(j) come il valore di una qualsiasi soluzione ottimale Sj costruita usando gli intervalli di indici{1,2,...,j}.
 * Quindi vale questa proprietà ricorsiva: Opt(j) = max{vj + Opt(p(j)),Opt(j-1)} (NB: se j appartiene a Sj vale il primo argomento, altrimenti il secondo.)
 *
*/


// Interval rappresenta un intervallo con un peso
type Interval struct {
    start int
    end int
    weight int
}

// maxWeightSequence calcola la sottosequenza di intervalli con la somma massima di pesi
// che non si sovrappongono
func maxWeightSequence(intervals []Interval) int {
    n := len(intervals)
    //dp è un array che viene utilizzato per memorizzare i pesi massimi delle sottosequenze di intervalli non sovrapposti
    //In pratica memorizza le soluzioni intermedie per non doverle ricalcolare ogni volta(memoization prog dinamica)
    dp := make([]int, n)
    //inizializzo la prima posizione di dp per avere un caso base per la programmazione dinamica
    //In pratica assumo che la sottosequenza di intervalli con peso massimo che non si sovrappone sia costituita solo dal primo intervallo
    dp[0] = intervals[0].weight

    for i := 1; i < n; i++ {
        dp[i] = intervals[i].weight
        for j := 0; j < i; j++ {
            if intervals[j].end <= intervals[i].start && dp[i] < dp[j]+intervals[i].weight {
                dp[i] = dp[j] + intervals[i].weight
            }
        }
    }
    maxWeight := 0
    for i := 0; i < n; i++ {
        if dp[i] > maxWeight {
            maxWeight = dp[i]
        }
    }
    return maxWeight
}

func main() {
    intervals := []Interval{
        {start: 1, end: 2, weight: 5},
        {start: 3, end: 4, weight: 1},
        {start: 0, end: 6, weight: 8},
        {start: 5, end: 7, weight: 4},
        {start: 3, end: 8, weight: 6},
    }
    fmt.Println(maxWeightSequence(intervals))
}

