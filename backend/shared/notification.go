package shared

import (
	"log"
	"os"
)

// SendDeploymentSuccessMessage sends a deployment success message and sets the webhook
func SendDeploymentSuccessMessage() {
	err := TelegramBot.SendMessage("🚀 Deployment Successful 🎉\nServer is running and ready to accept requests.")
	if err != nil {
		log.Printf("Failed to send deployment success message: %v", err)
	}

	// Re-set the webhook to ensure it's always up-to-date
	err = TelegramBot.SetWebhook(os.Getenv("WEBHOOK_URL"))
	if err != nil {
		log.Printf("Failed to set webhook: %v", err)
	}
}
