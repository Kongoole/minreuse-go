package main

import (
	"github.com/kongoole/minreuse-go/utils/log"
	"net/http"
	"os"

	_ "github.com/kongoole/minreuse-go/bootstrap"
)

func main() {
	port := string(os.Getenv("SERVER_PORT"))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
