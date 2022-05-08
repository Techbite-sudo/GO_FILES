package main
import(
	"log"
	"os"
)
func init(){
	log.SetPrefix("LOG:")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("file creation started....")
}
func main(){
	emptyFile,err := os.Create("empty.txt")
	errorCheck(err)
	log.Println(emptyFile)
	defer emptyFile.Close()

}
func errorCheck(err interface{}){
	if err != nil{
		log.Fatal(err)
	}

}