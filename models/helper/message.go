package message

import (
	"encoding/json"
	"log"
)

type Message struct{
  Value string `json:"value"`
  Status string `json:"status"`
  data map[string]interface{} `json:"data"`
}

func Response(res string,data []byte) Message{
	var response Message
	switch res{
		case "get":
			json.Unmarshal(data,&response.data)
			response = Message{
				Value : "ok",
				Status : "200",
			}
	}
	conf, _:= json.Marshal(response)
	log.Println(string(conf))
	return response
}