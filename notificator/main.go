package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"pg-notificator/lib"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/tx/:txid", scheduleTx)
	router.GET("/block/:block", scheduleBlock)

	log.Fatal(http.ListenAndServe(":80", router))
}

func scheduleTx(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	producer := lib.GetNsqProducer()
	m := lib.MessageStruct{
		Event: "tx",
		Hash:  ps.ByName("txid"),
	}
	message, err := json.Marshal(m)
	if err != nil {
		w.Write([]byte("Fuck!!!"))
	}
	log.Println(producer.Publish(os.Getenv("NSQ_TOPIC"), message))
	w.Write([]byte("OK"))
}

func scheduleBlock(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	producer := lib.GetNsqProducer()
	m := lib.MessageStruct{
		Event: "block",
		Hash:  ps.ByName("block"),
	}
	message, err := json.Marshal(m)
	if err != nil {
		w.Write([]byte("Fuck!!!"))
	}
	log.Println(producer.Publish(os.Getenv("NSQ_TOPIC"), message))
	w.Write([]byte("OK"))

}