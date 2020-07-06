package config

import (
	"encoding/json"
	"log"
	"os"
)

type ConfigStruct struct {
	BitcoinRPC struct{
		URL string `json:"url"`
		User string `json:"user"`
		Password string `json:"password"`
	} `json:"bitcoin-rpc"`
}

var appConfig ConfigStruct
var isInit = false

func Config() ConfigStruct {
	if !isInit {
		isInit = true
		file, _ := os.Open("config.json")

		defer func() {
			if err := file.Close(); err != nil {
				log.Println(err)
			}
		}()

		decoder := json.NewDecoder(file)

		err := decoder.Decode(&appConfig)
		if err != nil {
			log.Fatal(err)
		}
	}

	return appConfig
}
