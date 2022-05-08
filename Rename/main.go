package main

import (
	"log"
	"os"
)
func main(){
	oldName := "text.txt"
	newName := "textrenamed.txt"
	err := os.Rename(oldName,newName)
	errorCheck(err)
}
func errorCheck(err interface{}){
	if err != nil{
		log.Fatal(err)
	}

}