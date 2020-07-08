package api

import (
	btc_client "btc-faucet.drozd.by/modules/btc-client"
	"btc-faucet.drozd.by/modules/generator"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

type SendToAddressRequest struct {
	Address string `json:"address"`
	Amount float64 `json:"amount"`
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num * output)) / output
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

	txid, err := btc_client.SendToAddress(request.Address, toFixed(float64(request.Amount), 8))
	if err != nil {
		jsonErrorResponse(w, "Unable to send money. Error: " + err.Error())
		return
	}
	jsonResponse(w, map[string]string{
		"txid": txid,
		"address": request.Address,
		"amount": fmt.Sprintf("%.8f", request.Amount),
	})
}

func FaucetInfo(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, map[string]string{
		"address": generator.GetRefundAddress(),
	})
}
