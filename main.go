package main

import (
	"fmt"
	"os"
	"padi-back-go/route"
	"padi-back-go/setup"
)

func main() {
	r := setup.Setup()

	route.Initialize(r)

	port := os.Getenv("PORT")
	r.Run(fmt.Sprint(":", port))
}
