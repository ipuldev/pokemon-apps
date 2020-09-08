package construct

import (
	// "fmt"
	"log"
	"net/http"
	"encoding/json"
	"os"
	"io/ioutil"
	"github.com/gorilla/mux"
)

type DataResulting struct{
	Total int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []Children `json:"results"`

}

type Children struct{
	Id int `json:"Id"`
	Name string `json:"name"`
	Url string `json:"url"`
}



func GetData(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	limit := vars["limit"]
	data,err := http.Get("https://pokeapi.co/api/v2/pokemon?limit="+limit)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(2)
	}

	readData, err := ioutil.ReadAll(data.Body)
	if err != nil{
		log.Fatal(err)
	}
	
	var dataParsing DataResulting
	// fmt.Fprintf(w,string(readData))
	json.Unmarshal(readData,&dataParsing)
	for i := 1; i < len(dataParsing.Results); i++{
		dataParsing.Results[i].Id = i
	}
	json.NewEncoder(w).Encode(dataParsing)
	Logging("Get Data limit" + limit)
}


func Logging(data string){
	f, err := os.OpenFile("Logging", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
	    log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(data)
}