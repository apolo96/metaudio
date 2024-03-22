package main

import (
	"flag"
	"fmt"

	"github.com/apolo96/metaudio/handlers"
)

func main() {
	var port int
	flag.IntVar(&port, "p", 8000, "Port for metadata service")
	flag.Parse()
	fmt.Printf("Starting API at http://localhost:%d\n", port)
	handlers.Listen(port)
}
