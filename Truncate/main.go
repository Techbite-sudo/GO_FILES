package main
 
import (
	"log"
	"os"
)
 
func main() {
	fileName := "empty1.txt"
	err :=os.Truncate(fileName,100)
	errorCheck(err)
}
func errorCheck(err interface{}){
	if err != nil{
		log.Fatal(err)
	}

}