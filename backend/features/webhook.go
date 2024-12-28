package features

import (
	"backend/shared"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	var update shared.TelegramUpdate
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		response := shared.NewError(shared.BAD_REQUEST.Code(), "Invalid request payload")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	// Process the update from Telegram
	log.Printf("Received update: %+v", update)

	// Handle commands
	switch update.Message.Text {
	case "/hello":
		responseMessage := "Hello! How can I assist you today?"
		err := shared.TelegramBot.SendMessageToChatID(fmt.Sprintf("%d", update.Message.Chat.ID), responseMessage)
		if err != nil {
			log.Printf("Error sending message: %v", err)
		} else {
			log.Printf("Sent message: %s", responseMessage)
		}
	case "/info":
		responseMessage := "This is a bot that helps you with various tasks."
		err := shared.TelegramBot.SendMessageToChatID(fmt.Sprintf("%d", update.Message.Chat.ID), responseMessage)
		if err != nil {
			log.Printf("Error sending message: %v", err)
		} else {
			log.Printf("Sent message: %s", responseMessage)
		}
	case "/help":
		responseMessage := "Available commands: /hello, /info, /help, /today"
		err := shared.TelegramBot.SendMessageToChatID(fmt.Sprintf("%d", update.Message.Chat.ID), responseMessage)
		if err != nil {
			log.Printf("Error sending message: %v", err)
		} else {
			log.Printf("Sent message: %s", responseMessage)
		}
	case "/today":
		gregorianDate := time.Now().Format("2006-01-02")
		responseMessage := fmt.Sprintf("Today's date:\nGregorian: %s\nLunar: %s", gregorianDate, "N/A")
		err := shared.TelegramBot.SendMessageToChatID(fmt.Sprintf("%d", update.Message.Chat.ID), responseMessage)
		if err != nil {
			log.Printf("Error sending message: %v", err)
		} else {
			log.Printf("Sent message: %s", responseMessage)
		}
	default:
		responseMessage := "Unknown command. Type /help for a list of available commands."
		err := shared.TelegramBot.SendMessageToChatID(fmt.Sprintf("%d", update.Message.Chat.ID), responseMessage)
		if err != nil {
			log.Printf("Error sending message: %v", err)
		} else {
			log.Printf("Sent message: %s", responseMessage)
		}
	}

	response := shared.OK
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
