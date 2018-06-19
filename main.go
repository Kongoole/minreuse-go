package main

import (
	"net/http"
	"github.com/Kongoole/minreuse/route"
)

func main() {
	route.Register()
	//http.Handle("/", http.FileServer(http.Dir("route")))
	http.ListenAndServe(":8080", nil)
}