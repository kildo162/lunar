package features

import (
	"backend/shared"
	"encoding/json"
	"net/http"
)

func HandleHealthz(w http.ResponseWriter, r *http.Request) {
	response := shared.OK
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
