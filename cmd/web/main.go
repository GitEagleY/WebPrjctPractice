package main

import (
	"log"
	"net/http"

	render "github.com/GitEagleY/WebPrjctPractice/pkg/Render"
	"github.com/GitEagleY/WebPrjctPractice/pkg/config"

	handlers "github.com/GitEagleY/WebPrjctPractice/pkg/Handlers"
)

const portnum = ":8080"

var app config.AppConfig

func main() {
	app.UseCache = false //because in development mode
	///////////////////////////CACHING////////////////////////////
	tc, err := render.CacheTemplate() //creating template cache
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc //giving just created cache to app config

	//////////////////RENDER//////////////////////////////////////
	render.NewTemplates(&app) //giving app config to renders

	//////////////////WORK WITH HANDLERS AND REPO/////////////////
	repo := handlers.NewRepo(&app) //making new repo
	handlers.NewHandlers(repo)     //giving new just created repo to Handlers

	//http.HandleFunc("/", handlers.Repo.Home)
	//////////////////////////////////////////////////////////////
	srv := &http.Server{
		Addr:    portnum,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	//http.ListenAndServe(portnum, nil) //running server
}
