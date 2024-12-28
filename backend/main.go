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
	} else {
		log.Println("Telegram bot initialized successfully")
	}

	// Send deployment success message
	go func() {
		err := shared.TelegramBot.SendMessage("🚀 Deploy Successfully 🎉\nServer is up and running on port 8080")
		if err != nil {
			log.Printf("Failed to send deployment success message: %v", err)
		}

		// Re-set the webhook to ensure it's always up-to-date
		err = shared.TelegramBot.SetWebhook(os.Getenv("WEBHOOK_URL"))
		if err != nil {
			log.Printf("Failed to set webhook: %v", err)
		}
	}()

	http.HandleFunc("/api", features.HandleAPI)
	http.HandleFunc("/api/healthz", features.HandleHealthz)
	http.HandleFunc("/api/telegram/setup", features.HandleSetupTelegram)
	http.HandleFunc("/api/telegram/set-webhook", features.HandleSetWebhook)
	http.HandleFunc("/api/telegram/webhook", features.HandleWebhook)

	go func() {
		log.Println("Listening on port 8080...")
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()

	select {} // Block forever to keep the main function running
}
