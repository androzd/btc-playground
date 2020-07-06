package api

import (
	"btc-faucet.drozd.by/modules/generator"
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
)

func MiningStart(w http.ResponseWriter, r *http.Request) {
	generator.StartGeneratorRoutine()
	jsonResponse(w, map[string]string{
		"status": "ok",
	})
}

func MiningStop(w http.ResponseWriter, r *http.Request) {
	generator.StopGeneratorRoutine()
	jsonResponse(w, map[string]string{
		"status": "ok",
	})
}

func MiningStatus(w http.ResponseWriter, r *http.Request) {
	status := "false"
	if generator.StatusGeneratorRoutine() {
		status = "true"
	}
	jsonResponse(w, map[string]string{
		"status": status,
	})
}

func Router() http.Handler {
	r := chi.NewRouter()
	r.Get("/mining/start", MiningStart)
	r.Get("/mining/stop", MiningStop)
	r.Get("/mining/status", MiningStatus)
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