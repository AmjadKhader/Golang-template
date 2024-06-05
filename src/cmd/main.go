package main

import (
	config "main/pkg/config"
	routes "main/pkg/route"
)

func main() {

	config.Setup()

	router := routes.Setup()

	router.Run(":8080")
}
