package main

import (
	"backend/features"
	"backend/shared"
	"log"
	"net/http"
	"os"
)

func logRequest(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received API request: %s %s", r.Method, r.URL.Path)
		handler(w, r)
	}
}

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

	http.HandleFunc("/api", logRequest(features.HandleAPI))
	http.HandleFunc("/api/healthz", logRequest(features.HandleHealthz))
	http.HandleFunc("/api/telegram/setup", logRequest(features.HandleSetupTelegram))
	http.HandleFunc("/api/telegram/set-webhook", logRequest(features.HandleSetWebhook))
	http.HandleFunc("/api/telegram/webhook", logRequest(features.HandleWebhook))

	go func() {
		log.Println("Listening on port 8080...")
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()

	select {} // Block forever to keep the main function running
}
