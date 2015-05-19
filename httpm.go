package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	delay int
	statusCode int
)

func handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(delay) * time.Second)

	w.WriteHeader(statusCode)
	io.WriteString(w, "Hello world!")
}

func main() {
	port := flag.Int("port", 8000, "The server port.")

	flag.IntVar(&delay, "delay", 0, "Response delay.")
	flag.IntVar(&statusCode, "status", 200, "Response status code.")

	flag.Parse()

	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(": %v", *port), nil)
}


