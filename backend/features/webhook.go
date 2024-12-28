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
		return getTodayDateInfo()
	case "/year":
		return getYearInfo()
	case "/detail":
		return getDetailDayInfo()
	case "/nextday":
		return getNextDayInfo()
	default:
		return "Unknown command. Type /help for a list of available commands."
	}
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	response := shared.NewError(shared.BAD_REQUEST.Code(), message)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}

func getTodayDateInfo() string {
	now := time.Now()
	calendar := calendar.NewCalendar(calendar.CalendarDate{
		Day:   now.Day(),
		Month: int(now.Month()),
		Year:  now.Year(),
	})
	return fmt.Sprintf("Today's date:\nGregorian: %s\nLunar: %s",
		calendar.ToSolar().Format(),
		calendar.ToLunar().Format())
}

func getYearInfo() string {
	now := time.Now()
	solarDate := calendar.NewSolarDate(now.Year(), int(now.Month()), now.Day())
	return solarDate.YearInfo()
}

func getDetailDayInfo() string {
	now := time.Now()
	solarDate := calendar.NewSolarDate(now.Year(), int(now.Month()), now.Day())
	return solarDate.Detail()
}

func getNextDayInfo() string {
	now := time.Now().AddDate(0, 0, 1)
	solarDate := calendar.NewSolarDate(now.Year(), int(now.Month()), now.Day())
	return fmt.Sprintf("Next day's date:\nGregorian: %s\nLunar: %s",
		solarDate.Format(),
		calendar.NewLunarDate(now.Year(), int(now.Month()), now.Day()).Format())
}
