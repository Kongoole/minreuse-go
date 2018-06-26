package config

import (
	"os"
	"github.com/kongoole/minreuse-go/env"
)

var configuration = map[string]string{
	"view_folder": "view/",
	"controller_folder":"controller/",
}

func init() {
	// set app config
	for key, val := range configuration {
		os.Setenv(key, string(val))
	}

	// set env config
	env.ParseEnv()
}

