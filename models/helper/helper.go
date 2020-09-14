package helper

import (
	"log"
	"os"
	"fmt"
)

func ErrorCheck(err error){
	if err != nil{
		panic(err.Error())
		fmt.Println(err)
	}
}
func Logging_error(data string){
	f, err := os.OpenFile("Logging_error", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
	    log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(data)
}