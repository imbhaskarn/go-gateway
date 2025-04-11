package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting API Gateway on :8080")
	http.ListenAndServe(":8080", nil)
}
