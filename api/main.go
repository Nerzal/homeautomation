package main

import (
	"log"
	"net/http"
	"strings"
)

const dir = "./html"

var fs http.Handler

func main() {
	fs = http.FileServer(http.Dir(dir))
	log.Print("Serving " + dir + " on http://localhost:8080")

	http.ListenAndServe(":8080", http.HandlerFunc(handleRequest))
}

func handleRequest(resp http.ResponseWriter, req *http.Request) {
	println("requesting file:", req.URL.String())

	resp.Header().Add("Cache-Control", "no-cache")

	if strings.HasSuffix(req.URL.Path, ".wasm") {
		println("providing wasm.wasm file")
		resp.Header().Set("content-type", "application/wasm")
	}

	if strings.HasSuffix(req.URL.Path, ".css") {
		println("providing css file")
		resp.Header().Set("content-type", "text/css")
	}

	fs.ServeHTTP(resp, req)
}
