package main

import (
	"log"
	"net/http"

	render "github.com/GitEagleY/WebPrjctPractice/pkg/Render"
	"github.com/GitEagleY/WebPrjctPractice/pkg/config"

	handlers "github.com/GitEagleY/WebPrjctPractice/pkg/Handlers"
)

const portnum = ":8080"

func main() {

	var app config.AppConfig
	tc, err := render.CacheTemplate()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.ListenAndServe(portnum, nil)
}
