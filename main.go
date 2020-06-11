package main

import (
	"fmt"
	"os"
	"padi-back-go/config"
	"padi-back-go/setup"
)

func main() {
	cfg := config.NewConfig("config.yaml")

	r := setup.Setup()

	port := os.Getenv("PORT")
	if port == "" {
		port = fmt.Sprint(cfg.GetInt("app.port"))
	}
	r.Run(fmt.Sprint(":", port))
}
