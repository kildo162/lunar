package main

import (
	"backend/features"
	"backend/shared"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Starting server...")

	shared.TelegramBot = shared.InitTelegram(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if shared.TelegramBot == nil {
		log.Fatal("Failed to initialize Telegram bot")
	}

	http.HandleFunc("/api", features.HandleAPI)
	http.HandleFunc("/api/setup-telegram", features.HandleSetupTelegram)
	http.HandleFunc("/api/set-webhook", features.HandleSetWebhook)
	http.HandleFunc("/api/webhook", features.HandleWebhook)
	http.HandleFunc("/api/healthz", features.HandleHealthz)

	go func() {
		log.Println("Listening on port 8080...")
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()

	select {} // Block forever to keep the main function running
}
