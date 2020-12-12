package main

import (
	"github.com/api/router"
)

func main() {
	engine := router.StartRouter()
	engine.Run(":4000")
}
