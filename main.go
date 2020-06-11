package main

import (
	"fmt"
	"os"
	"padi-back-go/config"
	"padi-back-go/setup"
)

func main() {
	r := setup.Setup()

	port := os.Getenv("PORT")
	if port == "" {
		port = fmt.Sprint(config.GetInt("app.port"))
	}
	r.Run(fmt.Sprint(":", port))
}
