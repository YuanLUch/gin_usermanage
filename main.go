package main

import (
	"LoveDiary/routers"
	"LoveDiary/services"
)

func main() {
	services.InitDb()
	routers.InitRouter()
}
