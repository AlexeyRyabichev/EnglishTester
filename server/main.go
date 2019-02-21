package main

import (
	sw "./go"
	"log"
	"net/http"
	//sw "github.com/AlexeyRyabichev/EnglishTester"
	sw "./go"
)

func main() {
	log.Printf("Server started")
	sw.InitDB()
	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":8000", router))
}