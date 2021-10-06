package main

import (
	"fmt"
	"log"

	"github.com/Backend-GAuth-server/server"
)

func main() {
	app := server.Start()
	// next update
	// cli.Start()
	log.Fatal(app.Listen(":" + fmt.Sprint(8000)))
}
