package main
 
import (
	"fmt"
	"os"
)
 
func main() {
	message := "This content has been added to this file and using the append function "
	filename := "empty1.txt"
 
	// f, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	f,err := os.OpenFile(filename,os.O_RDWR | os.O_APPEND | os.O_CREATE,0660)
 
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer f.Close()
 
	fmt.Fprintf(f, "%s\n", message)
}