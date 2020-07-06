package btc_client

import (
	"errors"
	"fmt"
	"log"
)

func GetAddressesByLabel(label string) (response map[string]interface{}, err error) {
	resp, err := request("getaddressesbylabel", label)
	if err != nil {
		log.Fatal(err)
		return
	}
	if resp.Result != nil {
		response = resp.Result.(map[string]interface{})
		return
	}
	return
}

func GetAddressByLabelOrNew(label string) (address string, err error) {
	addresses, err := GetAddressesByLabel(label)
	if err != nil {
		log.Fatal(err)
		return
	}
	for k := range addresses {
		return k, nil
	}
	response, err := request("getnewaddress", label)
	if response.Result != nil {
		address = fmt.Sprintf("%v", response.Result)
		return
	}
	err = errors.New("unable generate address")
	return
}
