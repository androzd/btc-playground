package btc_client

import (
	"errors"
	"fmt"
	"log"
)

func SendToAddress(address string, amount float64) (txid string, err error) {
	response, err := request("sendtoaddress", address, amount)
	log.Println(response, response.Result, err)
	if err != nil {
		log.Fatal(err)
		return
	}
	if response.Result != nil {
		txid = fmt.Sprintf("%v", response.Result)
		return
	}
	err = errors.New(response.Error.Message)
	return
}