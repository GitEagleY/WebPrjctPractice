package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	render "github.com/GitEagleY/WebPrjctPractice/pkg/Render"
	"github.com/GitEagleY/WebPrjctPractice/pkg/config"
	"github.com/alexedwards/scs/v2"

	handlers "github.com/GitEagleY/WebPrjctPractice/pkg/Handlers"
)

const portnum = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.UseCache = false //because in development mode
	app.Production = false
	// /////////////////////SESSION////////////////////
	session = scs.New()
	session.Lifetime = 2 * time.Hour
	session.Cookie.Persist = false
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.Production
	//session.Cookie.Name="lol"
	//session.Cookie.Path="/"
	fmt.Println(session.Cookie)
	app.Session = session

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
