package main

import (
	_ "github.com/Kongoole/minreuse-go/bootstrap"
	"log"
	"net/http"
)

func main() {
	log.Println("server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}