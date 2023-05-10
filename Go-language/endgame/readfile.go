package main
import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	var fileName string
	fmt.Scan(&fileName)
	myFile,err := os.Open(fileName)
	if err != nil{
		fmt.Println("file not found")
		return
	}
	defer myFile.Close()

	myScanner := bufio.NewScanner(myFile)
	for myScanner.Scan() {
		line := myScanner.Text()
		fmt.Println(line)
	}
}
