package main

import (
	"net/http"
	"github.com/Kongoole/minreuse-go/route"
	"github.com/Kongoole/minreuse-go/config"
	"os"
	"fmt"
	"strings"
)

func main() {
	bootstrap()
	http.ListenAndServe(":8080", nil)
}

func bootstrap() {
	config.Config()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair)
	}
	route.Register()
}