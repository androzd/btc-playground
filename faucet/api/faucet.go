package api

import (
	btc_client "btc-faucet.drozd.by/modules/btc-client"
	"encoding/json"
	"fmt"
	"net/http"
)

type SendToAddressRequest struct {
	Address string `json:"address"`
	Amount float64 `json:"amount"`
}

func FaucetSendToAddress(w http.ResponseWriter, r *http.Request) {
	// Declare a new Person struct.
	var request SendToAddressRequest

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	txid, err := btc_client.SendToAddress(request.Address, request.Amount)
	if err != nil {
		http.Error(w, "Unable to send money. Error: " + err.Error(), http.StatusBadRequest)
		return
	}
	jsonResponse(w, map[string]string{
		"txid": txid,
		"address": request.Address,
		"amount": fmt.Sprintf("%.8f", request.Amount),
	})
}
