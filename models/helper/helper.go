package helper

import (
	"log"
	"os"
	"fmt"
	"net/http"
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

func EnableCors(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
