package api

import (
	"btc-faucet.drozd.by/modules/generator"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type MiningSetIntervalRequest struct {
	Counter time.Duration `json:"counter"`
	Measurement string `json:"measurements"`
}

func MiningStart(w http.ResponseWriter, r *http.Request) {
	generator.StartGeneratorRoutine()
	jsonResponse(w, map[string]string{
		"status": getMiningStatusAsString(),
	})
}

func MiningStop(w http.ResponseWriter, r *http.Request) {
	generator.StopGeneratorRoutine()
	jsonResponse(w, map[string]string{
		"status": "stopped",
	})
}

func MiningNewBlock(w http.ResponseWriter, r *http.Request) {
	generator.GenerateBlock()
	jsonResponse(w, map[string]string{
		"status": getMiningStatusAsString(),
	})
}

func MiningSetInterval(w http.ResponseWriter, r *http.Request) {
	// Declare a new Person struct.
	var request MiningSetIntervalRequest

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	d := request.Counter * time.Second
	generator.SetGeneratorInterval(d)
	jsonResponse(w, map[string]string{
		"status": getMiningStatusAsString(),
	})
}

func MiningStatus(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, map[string]string{
		"status": getMiningStatusAsString(),
		"interval": getMiningIntervalAsString(),
	})
}

func getMiningStatusAsString() string {
	status := "stopped"
	if generator.StatusGeneratorRoutine() {
		status = "working"
	}

	return status
}

func getMiningIntervalAsString() string {
	interval := generator.GetGeneratorInterval()
	return strconv.Itoa(int(interval.Seconds()))
}