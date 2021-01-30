package main

import (
	"net/http"

	"medium/medium_webserver_series/1/server"
)

func main() {
	mux := http.NewServeMux()
	server.New(mux)

	http.ListenAndServe(":8080", mux)
}
