package main

import (
	"log"
	"net/http"

	"rest/helpers"
)

func main() {
	router := helpers.SetupRouter()

	log.Fatal(http.ListenAndServe(":8000", router))
}
