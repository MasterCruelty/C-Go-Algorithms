package main

import (
	"fmt"
	"testing"
)

var prog = "./es3-interruttori"

func TestMainEsempioConFile(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(t,
		prog,
		"input-es.txt",
		"output-es.txt",
		"NIENTE")
}

func TestPremi(t *testing.T) {
	curr := []bool{false, true, true, true, true, true, true}

	output := premi(curr, 7)
	expected := []bool{true, true, true, true, true, false, false}

	for i := range output {
		if output[i] != expected[i] {
			fmt.Println("Output sbagliato!! Il bit in posizione", i, "non corrisponde")
			fmt.Println("ESECUZIONE:", output)
			fmt.Println("ATTESO:", expected)
			t.Fail()
			return
		}
	}
}

func TestSequenza(t *testing.T) {
	curr := []bool{false, true, true, true, true, true, true}

	output := sequenza(curr, []int{1, 2, 7, 4})
	expected := []bool{true, true, true, false, false, false, true}

	for i := range output {
		if output[i] != expected[i] {
			fmt.Println("Output sbagliato!! Il bit in posizione", i, "non corrisponde")

			fmt.Println("ESECUZIONE:", output)
			fmt.Println("ATTESO:", expected)
			t.Fail()
			return
		}
	}
}

func TestSpegniTutto(t *testing.T) {
	curr := []bool{true, false, false, false, false, false, false}

	output := spegniTutto(curr)

	if output != 5 {
		fmt.Println("Output sbagliato:")

		fmt.Println("ESECUZIONE:", output)
		fmt.Println("ATTESO:", 5)
		t.Fail()
	}

}
