package picking

import(
  "net/http"
  "fmt"
  "strconv"
  "encoding/json"

  // "github.com/saiful344/pokemon_app/models/helper"
  "github.com/jinzhu/gorm"
  "github.com/gorilla/mux"
     _ "github.com/jinzhu/gorm/dialects/mysql"
)

type User_pick struct{
  gorm.Model
  Id int `json:"id"`
  Name string `json:"name"`
  Power int64 `json:"power"`
  Unique string `json:"unique"`
}
func (User_pick) TableName() string{
  return "user_pick"
}

type Message struct{
  Value string `json:"value"`
  Status string `json:"status"`
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
  db := dbconn()
  defer db.Close()
  db.AutoMigrate(User_pick{})
  var data []User_pick
  db.Find(&data)
  w.WriteHeader(200)
  json.NewEncoder(w).Encode(data)
}

func GetDataOnePicking(w http.ResponseWriter,r *http.Request){
  db := dbconn()
  defer db.Close()
  db.AutoMigrate(User_pick{})
  vars := mux.Vars(r)
  id := vars["id"]
  var data []User_pick
  db.Model(&data).Where("id = ?",id).Find(&data)
  w.WriteHeader(200)
  json.NewEncoder(w).Encode(data)
}


func PostDataPicking(w http.ResponseWriter,r *http.Request){
  db := dbconn()
  defer db.Close()
  r.ParseForm()
  power ,err := strconv.ParseInt(r.FormValue("power"),10,32)
   if err != nil {
    fmt.Println(err)
  }
  var data User_pick
  data.Name  = r.FormValue("name")
  data.Power = power
  data.Unique= r.FormValue("unique")
  db.NewRecord(data)
  db.Create(&data)
  w.WriteHeader(200)
  output, err := json.Marshal(&Message{Value:"Success Added Data",Status:"200"})
  if err != nil {
      http.Error(w, err.Error(), 500)
      return
  }
  w.Header().Set("content-type", "application/json")
  w.Write(output)
}

func PutDataPicking(w http.ResponseWriter, r *http.Request){
  db := dbconn()
  defer db.Close()
  vars := mux.Vars(r)
  id := vars["id"]
  r.ParseForm()
  power ,err := strconv.ParseInt(r.FormValue("power"),10,32)
   if err != nil {
    fmt.Println(err)
  }
  var data User_pick
  data.Name  = r.FormValue("name")
  data.Power = power
  data.Unique= r.FormValue("unique")
  db.Model(&data).Where("id = ?",id).Update(&data)
  w.WriteHeader(200)
  output, err := json.Marshal(&Message{Value:"Success Update Data",Status:"200"})
  if err != nil {
      http.Error(w, err.Error(), 500)
      return
  }
  w.Header().Set("content-type", "application/json")
  w.Write(output)
}
func DeleteDataPicking(w http.ResponseWriter, r *http.Request){
  db := dbconn()
  defer db.Close()
  vars := mux.Vars(r)
  id := vars["id"]
  var data User_pick
  db.Model(&data).Where("id = ?", id).Delete(&data)
  w.WriteHeader(200)
  output, err := json.Marshal(&Message{Value:"Success Delete Data",Status:"200"})
  if err != nil {
      http.Error(w, err.Error(), 500)
      return
  }
  w.Header().Set("content-type", "application/json")
  w.Write(output)
}




