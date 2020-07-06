package main

import (
	"btc-faucet.drozd.by/api"
	"github.com/go-chi/chi"
	"net"
	"net/http"
	"path/filepath"
)

func main() {
	//fmt.Println("Start working. 1 second interval")
	//generator.StartGeneratorRoutine()
	//time.Sleep(time.Second * 5)
	//fmt.Println("Edit timer. 2 second interval")
	//generator.SetGeneratorInterval(2 * time.Second)
	//time.Sleep(time.Second * 5)
	//fmt.Println("Send stop signal")
	//generator.StopGeneratorRoutine()
	//fmt.Println("Stop signal sent")
	//return

	listen()
}

func listen() {
	listener, err := net.Listen("tcp", ":8080")
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