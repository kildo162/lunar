package features

import (
	"backend/shared"
	"encoding/json"
	"log"
	"net/http"
)

func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	log.Println("Received API request: /api/webhook")
	var update shared.TelegramUpdate
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		response := shared.NewError(shared.BAD_REQUEST.Code(), "Invalid request payload")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	// Process the update from Telegram
	log.Printf("Received update: %+v", update)
	response := shared.OK
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
