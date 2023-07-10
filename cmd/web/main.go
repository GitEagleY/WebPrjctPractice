package main

import "net/http"

const portnum = ":8080"

func homeContent(w http.ResponseWriter, r *http.Request) {

}
func main() {
	http.HandleFunc("/", homeContent)
	http.ListenAndServe(portnum, nil)
}
