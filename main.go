package main

import (
	"padi-back-go/config"
	"padi-back-go/route"
	"padi-back-go/setup"
)

func main() {
	config := config.NewConfig("config.yaml")

	r := setup.Setup()

	route.Initialize(r)

	host := config.GetString("app.host")
	r.Run(host)
}
