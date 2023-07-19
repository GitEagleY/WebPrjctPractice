package main

import (
	"net/http"

	handlers "github.com/GitEagleY/WebPrjctPractice/pkg/Handlers"
	"github.com/GitEagleY/WebPrjctPractice/pkg/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mx := chi.NewRouter()
	mx.Use(middleware.Recoverer)
	mx.Use(WriteToConsole)
	mx.Use(NoSurf)
	mx.Get("/", handlers.Repo.Home)
	return mx
}
