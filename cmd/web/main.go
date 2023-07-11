package main

import (
	"net/http"

	handlers "github.com/GitEagleY/WebPrjctPractice/pkg/Handlers"
)

const portnum = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.ListenAndServe(portnum, nil)
}
