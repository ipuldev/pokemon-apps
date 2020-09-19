package main

import(
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/saiful344/pokemon_app/models/construct"
	"github.com/saiful344/pokemon_app/models/picking"
)

func RedirectToHTTPSRouter(next http.Handler) http.Handler {
    return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
        proto := req.Header.Get("x-forwarded-proto")
        if proto == "http" || proto == "HTTP" {
            http.Redirect(res, req, fmt.Sprintf("https://%s%s", req.Host, req.URL), http.StatusPermanentRedirect)
            return
        }

        next.ServeHTTP(res, req)

    })
}


func HandleFunc(){
 	r := mux.NewRouter()
	httpsRouter := RedirectToHTTPSRouter(r)
	r.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"ok")
	})
	r.HandleFunc("/pokemon/get",picking.GetDataPicking).Methods("GET")
	r.HandleFunc("/pokemon/getone/{id}",picking.GetDataOnePicking).Methods("GET")
	r.HandleFunc("/pokemon/post",picking.PostDataPicking).Methods("POST")
	r.HandleFunc("/pokemon/put/{id}",picking.PutDataPicking).Methods("PUT")
	r.HandleFunc("/pokemon/delete/{id}",picking.DeleteDataPicking).Methods("DELETE")
	// From res out
	r.HandleFunc("/pokemon/get/{limit}",construct.GetData).Methods("GET","OPTIONS")
	r.HandleFunc("/pokemon/{name}",construct.GetByName).Methods("GET","OPTIONS")
	log.Fatal(http.ListenAndServe(":9000", httpsRouter))
}

func main(){
	HandleFunc();
	log.Println("Server Start ON port :9000")
}
