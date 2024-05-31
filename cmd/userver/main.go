package main

import (
	"log"

	"github.com/zviadadze/userver/api/server"
)

const configPath = "./configs/server-config.json"

func main() {
	if err := server.StartServer(configPath); err != nil {
		log.Fatal(err.Error())
	}
}
