package api

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
)

func Router() http.Handler {
	r := chi.NewRouter()
	r.Post("/mining/start", MiningStart)
	r.Post("/mining/stop", MiningStop)
	r.Post("/mining/new-block", MiningNewBlock)
	r.Post("/mining/set-interval", MiningSetInterval)
	r.Get("/mining/status", MiningStatus)
	r.Post("/faucet/send", FaucetSendToAddress)
	r.Get("/faucet/info", FaucetInfo)
	return r
}

func jsonResponse(w http.ResponseWriter, content map[string]string) {
	js, err := json.Marshal(content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func jsonErrorResponse(w http.ResponseWriter, errorMessage string) {
	content := map[string]string{"error": errorMessage}
	js, err := json.Marshal(content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}