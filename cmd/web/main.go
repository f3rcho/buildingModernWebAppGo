package main

import (
	"fmt"
	"net/http"

	"github.com/f3rcho/buildingModernWebAppGo/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	// this is not wrong, but its better practice use " Printf "
	// fmt.Println(fmt.Sprintf("Starting applicatieon on port %s", portNumber))
	fmt.Printf("Starting application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
