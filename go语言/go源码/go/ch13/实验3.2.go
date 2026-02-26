package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Server struct {
	ServerName string `json:"serverName"`
	ServerIP   string `json:"serverIP"`
}

type Serverslice struct {
	Servers []Server `json:"servers"`
}

func main() {
	jsonData := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	var serverData Serverslice
	err := json.Unmarshal([]byte(jsonData), &serverData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(serverData)
}
