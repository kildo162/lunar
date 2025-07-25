package main

import (
	"backend/features/api"
	"backend/shared"
	"log"
	"net/http"
	"os"

	customMiddleware "backend/features/middleware"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func main() {
	log.Println("Starting server...")

	// Print environment variables
	log.Printf("Environment Variables:\nTELEGRAM_BOT_TOKEN: %s\nWEBHOOK_URL: %s", os.Getenv("TELEGRAM_BOT_TOKEN"), os.Getenv("WEBHOOK_URL"))

	shared.TelegramBot = shared.InitTelegram(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if shared.TelegramBot == nil {
		log.Fatal("Failed to initialize Telegram bot")
	} else {
		log.Println("Telegram bot initialized successfully")
		shared.TelegramBot.InitChatIDsFromEnv()
	}

	r := chi.NewRouter()
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)
	r.Use(chiMiddleware.Timeout(60 * 1e9)) // 60s timeout
	r.Use(customMiddleware.CORSMiddleware)
	r.Get("/api", http.HandlerFunc(api.HandleAPI))
	r.Get("/api/healthz", http.HandlerFunc(api.HandleHealthz))
	r.Get("/api/calendar/today", http.HandlerFunc(api.HandleGetToday))
	r.Method("GET", "/api/calendar/solar-to-lunar", http.HandlerFunc(api.HandleConvertToLunar))
	r.Method("POST", "/api/calendar/solar-to-lunar", http.HandlerFunc(api.HandleConvertToLunar))
	r.Method("GET", "/api/calendar/lunar-to-solar", http.HandlerFunc(api.HandleConvertToSolar))
	r.Method("POST", "/api/calendar/lunar-to-solar", http.HandlerFunc(api.HandleConvertToSolar))
	r.Method("GET", "/api/calendar/month", http.HandlerFunc(api.HandleGetMonthCalendar))
	r.Method("POST", "/api/calendar/month", http.HandlerFunc(api.HandleGetMonthCalendar))
	r.Method("GET", "/api/calendar/good-days", http.HandlerFunc(api.HandleSearchGoodDays))
	r.Method("POST", "/api/calendar/good-days", http.HandlerFunc(api.HandleSearchGoodDays))
	// r.Method("GET", "/api/telegram/setup", features.HandleSetupTelegram)
	// r.Method("POST", "/api/telegram/setup", features.HandleSetupTelegram)
	// r.Method("GET", "/api/telegram/set-webhook", features.HandleSetWebhook)
	// r.Method("POST", "/api/telegram/set-webhook", features.HandleSetWebhook)
	// r.Method("POST", "/api/telegram/webhook", features.HandleWebhook)

	log.Println("Listening on port 8080...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
