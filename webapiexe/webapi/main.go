package main

import (
	"fmt"
	"net/http"
	"os"
)

var APPID string

func init() {
	APPID = os.Getenv("APPID")
	if APPID == "" {
		APPID = "APPID"
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>"+APPID+": Hello World</h1>")
}

func check(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>"+APPID+": Health check</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/health_check", check)
	fmt.Println("Server starting...")
	http.ListenAndServe(":3000", nil)
}
