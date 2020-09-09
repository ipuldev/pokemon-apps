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
	Name string `json:"name"`
	Url string `json:"url"`
}

const(
	
)

var dataParsing DataResulting

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
	
	// fmt.Fprintf(w,string(readData))
	json.Unmarshal(readData,&dataParsing)

	json.NewEncoder(w).Encode(dataParsing)
	Logging("Get Data limit" + limit)
}

func GetByName(w http.ResponseWriter, r *http.Request){
	// vars := mux.Vars(r)
	// name := vars["name"]
	log.Println(dataParsing.Results)
	for _,results := range dataParsing.Results{
		log.Println(results)
	}
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