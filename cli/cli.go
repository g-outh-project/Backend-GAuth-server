package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/Backend-GAuth-server/server"
)

func usage() {
	fmt.Printf("Welcome to GAUTH\n\n")
	fmt.Printf("Please see the following flags:\n\n")
	fmt.Printf("-port:			Set the PORT of the server\n")
	runtime.Goexit()
}

func Start() {

	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 8080, "Set port of ther server")

	flag.Parse()
	server.Start(*port)
}
