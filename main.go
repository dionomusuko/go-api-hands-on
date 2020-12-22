package main

import (
	"github.com/dionomusuko/go-api-hands-on/router"
)

func main() {
	e := router.Router()
	e.Logger.Fatal(e.Start(":8080"))
}
