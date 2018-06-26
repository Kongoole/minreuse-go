package bootstrap

import (
	_ "github.com/kongoole/minreuse-go/config"
	"github.com/kongoole/minreuse-go/route"
)

func init() {
	route.Register()
}
