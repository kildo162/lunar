package features

import (
	"backend/features/calendar"
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
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	log.Printf("Received update: %+v", update)

	responseMessage := handleCommand(update.Message.Text)
	err := shared.TelegramBot.SendMessageToChatID(fmt.Sprintf("%d", update.Message.Chat.ID), responseMessage)
	if err != nil {
		log.Printf("Error sending message: %v", err)
	} else {
		log.Printf("Sent message: %s", responseMessage)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(shared.OK)
}

func handleCommand(command string) string {
	switch command {
	case "/hello":
		return "Hello! How can I assist you today?"
	case "/info":
		return "This is a bot that helps you with various tasks."
	case "/help":
		return "Available commands: /hello, /info, /help, /today, /year, /detail, /nextday"
	case "/today":
		now := time.Now()
		calendar := calendar.NewCalendar(calendar.CalendarDate{
			Day:      now.Day(),
			Month:    int(now.Month()),
			Year:     now.Year(),
			TimeZone: +7,
			Hour:     now.Hour(),
			Min:      now.Minute(),
			Second:   now.Second(),
		})
		solarDate := calendar.ToSolar()
		lunarDate := calendar.ToLunar()

		return fmt.Sprintf("%s /n %s", solarDate.Detail(), lunarDate.Detail())
	default:
		return "Unknown command. Type /help for a list of available commands."
	}
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	response := shared.NewError(shared.BAD_REQUEST.Code(), message)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}
