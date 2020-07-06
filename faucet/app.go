package main

import (
	"btc-faucet.drozd.by/modules/generator"
	"encoding/json"
	"net"
	"net/http"
	"time"
)

func main() {
	initHttpRoutes()

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

func initHttpRoutes() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){http.ServeFile(w, r, "template/index.html")})
	http.HandleFunc("/api/startmining", func(w http.ResponseWriter, r *http.Request){
		generator.StartGeneratorRoutine()
		response(w)
	})
	http.HandleFunc("/api/stopmining", func(w http.ResponseWriter, r *http.Request){
		generator.StopGeneratorRoutine()
		response(w)
	})
}

func response(w http.ResponseWriter) {
	response := map[string]string{"status":"ok"}

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}


func listen() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 16,
	}
	server.Serve(listener)
}