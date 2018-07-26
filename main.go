package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/kongoole/minreuse-go/bootstrap"
)

func main() {
	port := string(os.Getenv("SERVER_PORT"))
	log.Println("server listening on " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
