package main

import (
	"go-test-api/routers"
)

func main() {
	app := routers.Router()
	app.Run(":4002")
}
