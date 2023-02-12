package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"
)

func LanciaGenerica(
	t *testing.T,
	progname string,
	strinput string,
	oracolo string,
	args ...string) {

	os := runtime.GOOS
	switch os {
	case "windows":
		progname += ".exe"
	}

	subproc := exec.Command(progname, args...)

	go func() {
		time.Sleep(time.Second * 10)
		if subproc.ProcessState.ExitCode() == -1 {
			fmt.Println("time limit!")
			subproc.Process.Kill()
		}
	}()

	subproc.Stdin = strings.NewReader(strinput)
	stdout, err := subproc.CombinedOutput()

	if err != nil {
		fmt.Printf("Attenzione! Uscito con codice: %s\n>>> ", err)
		t.Fail()
		return
	}

	if string(stdout) != oracolo {
		fmt.Println("Fail esecuzione main")
		fmt.Printf("/// Argomenti: %s\n", args)
		fmt.Printf("\n/// Input:\n%s\n", strinput)
		fmt.Printf("\n/// eseguo diff... \n")
		fmt.Println("\nESECUZIONE:\n<<<<<")
		fmt.Printf("%s", string(stdout))
		fmt.Println(">>>>")

		fmt.Println("\nATTESO:\n<<<<<")
		fmt.Printf("%s", oracolo)
		fmt.Println(">>>>")
		t.Fail()
	}

	subproc.Process.Kill()
	fmt.Println()
}

func LanciaGenericaConFileOutAtteso(
	t *testing.T,
	nomeProg string,
	strinput string,
	oracoloFilename string,
	args ...string) {

	content, err := ioutil.ReadFile(oracoloFilename)
	if err != nil {
		log.Fatal(err)
	}
	text := string(content)

	LanciaGenerica(t, nomeProg, strinput, text, args...)
}

func LanciaGenericaConFileInOutAtteso(
	t *testing.T,
	nomeProg string,
	inputFilename string,
	oracoloFilename string,
	args ...string) {

	input, err := ioutil.ReadFile(inputFilename)
	if err != nil {
		log.Fatal(err)
	}
	in := string(input)

	exout, err := ioutil.ReadFile(oracoloFilename)
	if err != nil {
		log.Fatal(err)
	}
	out := string(exout)

	LanciaGenerica(t, nomeProg, in, out, args...)
}

func intorno(a, b float64) bool {
	return math.Abs(a-b) < 10e-6
}

func CaptureOutput(fn interface{}, args ...interface{}) string {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f := reflect.ValueOf(fn)
	if f.Type().NumIn() != len(args) {
		panic("incorrect number of parameters!")
	}
	inputs := make([]reflect.Value, len(args))
	for k, in := range args {
		inputs[k] = reflect.ValueOf(in)
	}
	f.Call(inputs)

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	return string(out)
}

func checkOutput(output, expected string) bool {
	if output != expected {
		fmt.Printf("\n/// eseguo diff... \n")
		fmt.Println("\nESECUZIONE:\n<<<<<")
		fmt.Printf("%s", output)
		fmt.Println(">>>>")

		fmt.Println("\nATTESO:\n<<<<<")
		fmt.Printf("%s", expected)
		fmt.Println(">>>>")
		return false
	}
	return true

}
