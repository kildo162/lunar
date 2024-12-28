package features

import (
	"backend/shared"
	"encoding/json"
	"log"
	"net/http"
)

func HandleHealthz(w http.ResponseWriter, r *http.Request) {
	log.Println("Received API request: /api/healthz")
	response := shared.OK
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
