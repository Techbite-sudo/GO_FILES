package main
 
import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	
)
 
func main() {
	fileName := "empty1.txt"
	fileBuffer,error := ioutil.ReadFile(fileName)
	if error != nil{
		fmt.Println("error")
		os.Exit(1)
	}
	inputData := string(fileBuffer)
	data := bufio.NewScanner(strings.NewReader(inputData))
	data.Split(bufio.ScanRunes)

	for data.Scan(){
		fmt.Print(data.Text())
	}

}
