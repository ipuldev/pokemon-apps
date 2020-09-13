package main

import(
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/saiful344/pokemon_app/models/construct"
	"github.com/saiful344/pokemon_app/models/picking"
)


func HandleFunc(){
 	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"ok")
	})
	r.HandleFunc("/get",picking.GetDataPicking)
	r.HandleFunc("/pokemon/get/{limit}",construct.GetData)
	r.HandleFunc("/pokemon/{name}",construct.GetByName)
	log.Fatal(http.ListenAndServe(":9000",r))

}

func main(){
	HandleFunc();
	log.Println("Server Start ON port :9000")
}
