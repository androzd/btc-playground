package btc_client

import (
	"fmt"
	"log"
)

func SendToAddress(address string, amount float64) (txid string, err error) {
	response, err := request("sendtoaddress", address, amount)
	if err != nil {
		log.Fatal(err)
		return
	}
	if response.Result != nil {
		txid = fmt.Sprintf("%v", response.Result)
		return
	}
	return
}