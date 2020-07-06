package api

import (
	"btc-faucet.drozd.by/modules/generator"
	"net/http"
)

func MiningStart(w http.ResponseWriter, r *http.Request) {
	generator.StartGeneratorRoutine()
	jsonResponse(w, map[string]string{
		"status": getMiningStatusAsString(),
	})
}

func MiningStop(w http.ResponseWriter, r *http.Request) {
	generator.StopGeneratorRoutine()
	jsonResponse(w, map[string]string{
		"status": getMiningStatusAsString(),
	})
}

func MiningStatus(w http.ResponseWriter, r *http.Request) {

	jsonResponse(w, map[string]string{
		"status": getMiningStatusAsString(),
	})
}

func getMiningStatusAsString() string {
	status := "stopped"
	if generator.StatusGeneratorRoutine() {
		status = "working"
	}

	return status
}