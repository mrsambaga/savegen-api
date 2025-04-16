package main

import (
	"log"
	"savegen-api/db"
	"savegen-api/server"
)

func main() {
	err := db.Connect()
	if err != nil {
		log.Println("Failed to connect DB", err)
		return
	}

	server.Init()
} 