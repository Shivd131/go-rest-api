// main.go
package main

import (
	"fmt"
	"log"

	"github.com/Shivd131/api/db"
	"github.com/Shivd131/api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Connect to db
	err := db.InitDB()
	if err != nil {
		log.Fatal("Error initializing database:", err)
	}

	//  API routes
	router.GET("/items", handlers.GetItems)
	router.GET("/items/:id", handlers.GetItem)
	router.POST("/items", handlers.CreateItem)
	router.PUT("/items/:id", handlers.UpdateItem)
	router.DELETE("/items/:id", handlers.DeleteItem)

	port := ":8080"
	fmt.Printf("Server is running on port %s\n", port)
	log.Fatal(router.Run(port))
}
