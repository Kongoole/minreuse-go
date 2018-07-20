package main

import (
	"log"
	"net/http"

	_ "github.com/kongoole/minreuse-go/bootstrap"
)

func main() {
	log.Println("server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
