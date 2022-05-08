package main

import(
	"log"
	"os"
)
func main(){
	oldLocation := "/home/boniface/GO_FILESDIR/Files/empty.txt"
	newLocation := "/home/boniface/GO_FILESDIR/Move/empty.txt"
	err := os.Rename(oldLocation,newLocation)
	errorCheck(err)

}
func errorCheck(err interface{}){
	if err != nil{
		log.Fatal(err)
	}

}