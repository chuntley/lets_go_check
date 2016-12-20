package main

import (
	"net/http"
	"flag"
	"fmt"
)

func main() {
	checkPath := flag.String("checkPath", "/var/www/letsencrypt", "The path to check")
	serverPort := flag.String("port", "9009", "The server port")
	flag.Parse()

	fmt.Printf("Running server on port %s, checking the path %s", *serverPort, *checkPath)
	panic(http.ListenAndServe(":" + *serverPort, http.FileServer(http.Dir(*checkPath))))
}