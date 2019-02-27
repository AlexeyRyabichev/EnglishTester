package main

import (
	sw "./go"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")
	sw.InitDB()
	sw.CreateSchemaTeachers()
	sw.CreateSchemaStudents()
	sw.CreateSchemaTest()
	sw.CreateSchemaAudio()
	router := sw.NewRouter()
	//sw.InsertTests()

	log.Fatal(http.ListenAndServe(":8080", router))
}
