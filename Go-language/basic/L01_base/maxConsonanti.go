package main

import( "fmt"
	"strings"
	"regexp"
)

//dato un vettore di stringhe, restituire il numero massimo di consonanti contenute in una stringa
func main() {
	const N = 5
	parole := make([]string,N)

	for i:= 0;i < N;i++ {
		fmt.Scan(&parole[i])
	}

	fmt.Println(maxConsonanti(parole))
}

func maxConsonanti(parole []string) int {
	max := 0
	for i := range parole{
		parola := strings.ToLower(parole[i])
		var regex = regexp.MustCompile("[aeiou]+")
		parola = regex.ReplaceAllString(parola,"")
		if len(parola) > max {
			max = len(parola)
		}
	}
	return max
}
