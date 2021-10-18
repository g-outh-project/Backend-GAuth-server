package main

import (
	"fmt"
	"log"

	"github.com/Backend-GAuth-server/db"
	"github.com/Backend-GAuth-server/server"
)

func main() {
	app, file := server.Start()
	// next update
	// cli.Start()
	// Release resource
	defer db.CloseDB()
	defer file.Close()

	log.Fatal(app.Listen(":" + fmt.Sprint(8000)))
}
