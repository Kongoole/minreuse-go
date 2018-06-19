package main

import (
	_ "github.com/Kongoole/minreuse-go/bootstrap"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", nil)
}