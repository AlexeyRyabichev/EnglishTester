package main

import (
	sw "./go"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")
	sw.InitDB()
	err := sw.CreateSchemaTeachers()
	if err != nil {
		log.Print(err)
	}
	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
