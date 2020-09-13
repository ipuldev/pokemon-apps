package picking

import(
  "net/http"
  // "log"
  "fmt"
  "io/ioutil"
  "encoding/json"

  // "github.com/saiful344/pokemon_app/models/message"
  "github.com/saiful344/pokemon_app/models/helper"
  "github.com/jinzhu/gorm"
    // "github.com/gorilla/mux"
     _ "github.com/jinzhu/gorm/dialects/mysql"
)

type User_pick struct{
  gorm.Model
  Id int `json:"id"`
  Name string `json:"name"`
  Power int `json:"power"`
  Unique string `json:"unique"`
}
func (User_pick) TableName() string{
  return "user_pick"
}

func dbconn()(db *gorm.DB){
  db, err := gorm.Open("mysql", "root:root@/pokemon?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
     panic(err.Error())
     fmt.Println("can't connect to db")
  }
  return db
}
func GetDataPicking(w http.ResponseWriter,r *http.Request){
  w.Header().Set("Content-Type", "text/html; charset=ascii")
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Headers","Content-Type,access-control-allow-origin, access-control-allow-headers")
  db := dbconn()
  defer db.Close()
  db.AutoMigrate(User_pick{})
  var data []User_pick
  db.Find(&data)
  w.WriteHeader(200)
  json.NewEncoder(w).Encode(data)
  fmt.Println("success get data")
}

func PostDataPicking(w http.ResponseWriter,r *http.Request){
  res, err := ioutil.ReadAll(r.Body())
  helper.ErrorCheck(err)
  
}



