package main
 
import (
	"fmt"
	"log"
	"os"
)
 
func main() {
	fileStat, err := os.Stat("string.txt")
	errorCheck(err)

	fmt.Println("File Name: ",fileStat.Name()) 
	fmt.Println("Size: ",fileStat.Size()) 
	fmt.Println("Permissions: ",fileStat.Mode())
	fmt.Println("Last Modified: ",fileStat.ModTime())
	fmt.Println("Is Directory: ",fileStat.IsDir())  
	fmt.Println("System: ",fileStat.Sys())  
}
func errorCheck(err interface{}){
	if err != nil{
		log.Fatal(err)
	}

}