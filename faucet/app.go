package main

import (
	"net"
	"net/http"
	"time"
)

func main() {
	server := &http.Server{
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 16,
	}

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){http.ServeFile(w, r, "template/index.html")})

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	server.Serve(listener)
}