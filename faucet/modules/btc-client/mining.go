package btc_client

import (
	"log"
)

func GenerateToAddress(nblocks int, address string) (err error) {
	resp, err := request("generatetoaddress", nblocks, address)
	if err != nil {
		log.Fatal(err)
		return
	}
	if resp.Result != nil {
		return
	}
	return
}