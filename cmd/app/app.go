package main

import (
	"fmt"
	"net/http"
)

func startServer() {
	http.HandleFunc("/get")
}

func main() {
	fmt.Println("Starting service")
	startServer()
}
