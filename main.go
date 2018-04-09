package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	address = flag.String("address", "localhost:8080", "Address used for the webserver [host]:port")
	help    = flag.Bool("help", false, "Display this message")
)

func main() {
	flag.Parse()
	if *help {
		flag.Usage()
		return
	}

	fmt.Printf("Starting webserver on http://%s/\n", *address)
	err := StartServer(*address)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
