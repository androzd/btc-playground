package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"pg-notificator/lib"
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
		_, err = w.Write([]byte("Not OK"))
		if err != nil {
			log.Println(err)
		}
	}
	log.Println(producer.Publish(os.Getenv("NSQ_TOPIC"), message))
	_, err = w.Write([]byte("OK"))
	if err != nil {
		log.Println(err)
	}
}

func scheduleBlock(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	producer := lib.GetNsqProducer()
	m := lib.MessageStruct{
		Event: "block",
		Hash:  ps.ByName("block"),
	}
	message, err := json.Marshal(m)
	if err != nil {
		_, err = w.Write([]byte("Not OK"))
		if err != nil {
			log.Println(err)
		}
	}
	log.Println(producer.Publish(os.Getenv("NSQ_TOPIC"), message))
	_, err = w.Write([]byte("OK"))
	if err != nil {
		log.Println(err)
	}
}