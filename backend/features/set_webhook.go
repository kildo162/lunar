package features

import (
	"backend/shared"
	"encoding/json"
	"net/http"
)

func HandleSetWebhook(w http.ResponseWriter, r *http.Request) {
	var request struct {
		URL string `json:"url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response := shared.NewError(shared.BAD_REQUEST.Code(), "Invalid request payload")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	err := shared.TelegramBot.SetWebhook(request.URL)
	if err != nil {
		response := shared.NewError(shared.SERVER_ERROR.Code(), "Failed to set webhook")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}
	response := shared.OK
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
