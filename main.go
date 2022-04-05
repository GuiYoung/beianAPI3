package main

import (
	"beianAPI/model"
	"beianAPI/routes"
)

func main() {
	if err := model.InitDb(); err != nil {
		panic(err)
	}

	routes.InitRouter()
}
