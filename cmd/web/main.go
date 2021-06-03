package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/f3rcho/buildingModernWebAppGo/pkg/config"
	"github.com/f3rcho/buildingModernWebAppGo/pkg/handlers"
	"github.com/f3rcho/buildingModernWebAppGo/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	// this is not wrong, but its better practice use " Printf "
	// fmt.Println(fmt.Sprintf("Starting applicatieon on port %s", portNumber))
	fmt.Printf("Starting application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
