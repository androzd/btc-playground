package main

import (
	"btc-faucet.drozd.by/api"
	"github.com/go-chi/chi"
	"net"
	"net/http"
	"path/filepath"
)

func main() {
	listen()
}

func listen() {
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		panic(err)
	}

	router := chi.NewRouter()

	staticPath, _ := filepath.Abs("./assets/")
	fs := http.StripPrefix("/assets/", http.FileServer(http.Dir(staticPath)))
	router.Handle("/assets/*", fs)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){http.ServeFile(w, r, "assets/index.html")})
	router.Mount("/api/", api.Router())

	http.Serve(listener, router)
}