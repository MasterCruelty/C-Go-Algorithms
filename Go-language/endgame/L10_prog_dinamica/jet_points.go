package main

import "fmt"

/*
 *Una compagnia aerea offre ai suoi viaggiatori dei jet-points: per ogni tratta è definito un valore in jet-points; 
 *i viaggiatori possono accumulare jet-points nelle varie tratte e usarli, al posto dei dollari, per acquistare delle tratte; 
 *I jet-points di un viaggio si accumulano additivamente, cioè i punti acquisiti con una tratta si dommano ai punti precedentemente acquisiti.
 *Il programma dei jet-points è scaduto ma è possibile convertirli in buoni euro chiamati swindles. 
 *Di seguito la tabella che riassume il costo in jet-points degli swindles:
 *
 *Swindles :jet-points
 *100:20000
 *107:22000
 *124:24000
 *133:26000
 *139:28000
 *155:30000
 *172:32000
 *178:34000
 *184:36000
 *190:38000
 *195:40000
 *
 *Ad esempio se ho 60000 jet-points posso avere swindles pari a 311 euro(uno da 133 e uno da 178).
 *Come è possibile calcolare il massimo valore in swindles acquistabili con j jet-pints, avendo in input una tabella tipo quella mostrata?
*/


// jetPointsToSwindles mappa i jet-points ai relativi swindles
var jetPointsToSwindles = map[int]int{
	100: 20000,
	107: 22000,
	124: 24000,
	133: 26000,
	139: 28000,
	155: 30000,
	172: 32000,
	178: 34000,
	184: 36000,
	190: 38000,
	195: 40000,
}

// maxSwindles calcola il massimo valore in swindles acquistabili con j jet-points
func maxSwindles(j int) int {
	// dp è la tabella che memorizza il massimo valore in swindles per ogni numero di jet-points
	dp := make([]int, j+1)

	// scorriamo attraverso tutti i jet-points
	for i := 1; i <= j; i++ {
		// consideriamo il massimo valore in swindles acquistabile con i jet-points
		maxSwindles := 0
		for swindles, jetPoints := range jetPointsToSwindles {
			// se i jet-points sono sufficienti per acquistare i swindles
			if jetPoints <= i {
				// verifichiamo se l'acquisto di questi swindles porta ad un valore maggiore rispetto all'acquisto precedente
				if swindles+dp[i-jetPoints] > maxSwindles {
					maxSwindles = swindles + dp[i-jetPoints]
				}
			}
		}
		// memorizziamo il massimo valore in swindles per i jet-points correnti
		dp[i] = maxSwindles
	}
	return dp[j]
}

func main() {
	jetPoints := 60000
	fmt.Println("Max swindles:", maxSwindles(jetPoints))
	jetPoints = 100000
	fmt.Println("Max swindles:", maxSwindles(jetPoints))
	jetPoints = 140000
	fmt.Println("Max swindles:", maxSwindles(jetPoints))
}

