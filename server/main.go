package main

import (
	sw "./go"
	"log"
	"net/http"
	//sw "github.com/AlexeyRyabichev/EnglishTester/tree/master/server/go"
	// sw "./go"
)

func main() {
	log.Printf("Server started")
	sw.InitDB()
	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":8000", router))
}