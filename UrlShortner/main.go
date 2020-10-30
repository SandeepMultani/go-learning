package main

import (
	urlshort "UrlShortner/urlshort"
	"fmt"
	"net/http"
)

func main() {
	mux := defaultMux()
	fmt.Println(urlshort.HelloFromPackage("sandeep"))
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8000", mux)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	return mux
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}
