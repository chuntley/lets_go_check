package main

import "net/http"
import "flag"

func main() {
	checkPath := flag.String("checkPath", "/var/www/letsencrypt", "The path to check")
	serverPort := flag.String("port", "9009", "The server port")

	panic(http.ListenAndServe(":" + *serverPort, http.FileServer(http.Dir(*checkPath))))
}