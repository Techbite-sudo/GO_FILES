package main

import(
	"log"
	"os"
	"io"
)
func main(){
	sourceFile,err := os.Open("/home/boniface/GO_FILESDIR/Move/empty1.txt")
	errorCheck(err)
	defer sourceFile.Close()

	newFile,err := os.Create("/home/boniface/GO_FILESDIR/Copy/empty1.txt")
	errorCheck(err)
	log.Println(newFile)
	defer newFile.Close()

	bytesCopied, err := io.Copy(newFile, sourceFile)
	errorCheck(err)
	log.Printf("Copied %d bytes.", bytesCopied)
}
func errorCheck(err interface{}){
	if err != nil{
		log.Fatal(err)
	}
}