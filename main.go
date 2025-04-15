package main

import (
	"log"
	"savegen-api/db"

	"github.com/gin-gonic/gin"
)

func main() {
	err := db.Connect()
	if err != nil {
		log.Println("Failed to connect DB", err)
		return
	}
	
	r := gin.Default()

	r.Run(":8080")
} 