package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ak1729/go-crud-mongodb/router"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

	port := os.Getenv("PORT")

	fmt.Println("Mongo Api")
	r := router.Router()
	fmt.Println("Server is getting started...")

	log.Fatal(http.ListenAndServe(":"+port, r))

	fmt.Println("listening at port 8000")
}
