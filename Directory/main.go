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
	_,err := os.Stat("test1")
	if os.IsNotExist(err) {
		errDir := os.MkdirAll("test1",0755)
		if errDir != nil{
			log.Fatal(err)
		}
	}

}
// func errorCheck(err interface{}){
// 	if err != nil{
// 		log.Fatal(err)
// 	}

// }
