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

type Single struct{
	Abilities []Abilities `json:"abilities"`
	Forms []Forms `json:"forms"`
	Id int `json:"id"`
}
type Abilities struct{
	Ability map[string]interface{} `json:"ability"`
	Is_hiddden bool `json:"is_hidden"`
	Slot int `json:"slot"`
}

type Forms struct{
	Name string `json:"name"`
	Url string `json:"url"`
}

type Img struct{
	Sprites Sprites `json:"sprites"`
}
type Sprites struct{
	Back_default string `json:"back_default"`
	Back_shiny string `json:"back_shiny"`
	Front_default string `json:"front_default"`
	Front_shiny string `json:"front_shiny"`
}

const(
	urls = "https://pokeapi.co/api/v2/pokemon/"
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
	vars := mux.Vars(r)
	name := vars["name"]
	result, err := http.Get(urls + name)
	if err != nil{
		log.Fatal(err)
	}
	parseData, err := ioutil.ReadAll(result.Body)
	if err != nil{
		log.Fatal(err)
	}
	var Data Single
	json.Unmarshal(parseData,&Data)

	// Getting image
	result_img, err := http.Get(Data.Forms[0].Url)
	if err != nil{
		log.Fatal(err)
	}
	parseData_img, err := ioutil.ReadAll(result_img.Body)
	if err != nil{
		log.Fatal(err)
	}
	var img Img
	json.Unmarshal(parseData_img, &img)

	StoreData := make(map[string]interface{})
	StoreData["Data"] = Data
	StoreData["Data_Img"] = img

	json.NewEncoder(w).Encode(StoreData)
	Logging("Get Single Data" + name)
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
