package bootstrap

import (
	_ "github.com/Kongoole/minreuse-go/config"
	"github.com/Kongoole/minreuse-go/route"
)

func init() {
	route.Register()
}
