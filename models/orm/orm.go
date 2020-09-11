package orm

import(
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"database/mysql"

	_ "github.com/go-sql-driver/mysql"
)

type DBConn struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Power int `json:"power"`
	Unique string `json:"unique"`
}
// Database Connect
func dbconn()(db *sql.DB){
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "pokemon"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+dbName)
	err != nil{
		log.Println(err.Error())
	}
	return db
}


func getData(w htttp.ResponseWriter, r *http.Request){
	
}

// Orm not use package
func InsertData(w http.ResponseWriter, r *http.Request){

}
func UpdateData(w http.ResponseWriter, r *http.Request){

}
func DeleteData(w http.ResponseWriter, r *http.Request){

}