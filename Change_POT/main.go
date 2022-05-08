package main

import (
	// "fmt"
	"log"
	"os"
	"time"
)
func main(){
		//test file existence 
		_,err := os.Stat("empty1.txt")
		if err != nil{
			if os.IsNotExist(err){
				log.Fatal("File does not exist")
			}
		}
		log.Println("File exists")

		//change permissions 
		err = os.Chmod("empty1.txt",0777)
		if err != nil{
			log.Println(err)
		}
		//change ownership
		err = os.Chown("empty1.txt",os.Getuid(),os.Getgid())
		if err != nil{
			log.Println(err)
		}
		//change timestampp
		addOneDayFromNow := time.Now().Add(24 * time.Hour)
		lastAccessTime := addOneDayFromNow
		lastModifyTime := addOneDayFromNow
		err = os.Chtimes("empty1.txt",lastAccessTime,lastModifyTime)
		if err != nil{
			log.Println(err)
		}
		
}