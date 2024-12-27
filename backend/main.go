package main

import (
	"backend/shared"
	"os"
)

func main() {

	shared.TelegramBot = shared.InitTelegram(os.Getenv("TELEGRAM_BOT_TOKEN"))

}
