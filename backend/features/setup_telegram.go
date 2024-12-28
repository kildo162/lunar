package features

import (
	"backend/shared"
	"encoding/json"
	"net/http"
)

func HandleSetupTelegram(w http.ResponseWriter, r *http.Request) {
	var request struct {
		ChatID string `json:"chat_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response := shared.NewError(shared.BAD_REQUEST.Code(), "Invalid request payload")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	shared.TelegramBot.AddChatID(request.ChatID)
	response := shared.OK
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
