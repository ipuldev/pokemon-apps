package orm

import(
	"net/http"
	// "log"
	// "io/ioutil"
	// "encoding/json"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type rester struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Power int `json:"power"`
	Unique string `json:"unique"`
}
// Database Connect
// func dbconn()(db *sql.DB){
// 	dbDriver := "mysql"
// 	dbUser := "sijon"
// 	dbPass := "sijon"
// 	dbName := "pokemon"
// 	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:3306)/"+dbName)
// 	if err != nil{
// 		log.Println(err.Error())
// 	}
// 	return db
// }




func GetData(w http.ResponseWriter, r *http.Request){
	db, err := sql.Open("mysql","sijon:sijon@tcp(127.0.0.1:3306)/pokemon")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	// db := dbconn();
	// result, err := db.Query("SELECT * FROM pokemons ORDER BY id ASC")
	// if  err != nil {
	// 	panic(err.Error())
	// }
	// store := DBConn{}
	// res := []DBConn{}
	// for result.Next(){
	// 	var id, power int
	// 	var name, unique string
	// 	err = result.Scan(&id, &name, &power, &unique)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	store.Id = id;
	// 	store.Name = name;
	// 	store.Power = power;
	// 	store.Unique = unique;
	// 	res = append(res, store)
	// }
	// log.Println(res)
	// defer db.Close();
}

// Orm not use package
func InsertData(w http.ResponseWriter, r *http.Request){

}
func UpdateData(w http.ResponseWriter, r *http.Request){

}
func DeleteData(w http.ResponseWriter, r *http.Request){

}
