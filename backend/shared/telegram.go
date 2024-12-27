package shared

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var TelegramBot *Telegram

type Telegram struct {
	BotToken string
	ChatIDs  []string
}

func InitTelegram(botToken string) *Telegram {
	return &Telegram{
		BotToken: botToken,
	}
}

func (t *Telegram) AddChatID(chatID string) {
	t.ChatIDs = append(t.ChatIDs, chatID)
}

func (t *Telegram) SendMessage(message string) {
	// Tạo payload để gửi request
	payload := map[string]string{
		"chat_id": t.ChatIDs[0],
		"text":    message,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error encoding payload: %v", err)
	}

	// Gửi HTTP POST request
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", t.BotToken)
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Kiểm tra phản hồi từ Telegram
	if resp.StatusCode == http.StatusOK {
		log.Println("Tin nhắn đã được gửi thành công!")
	} else {
		log.Printf("Lỗi: %s\n", resp.Status)
	}
}

func (t *Telegram) InitAllChatIDs() {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates", t.BotToken)
	resp, err := http.Get(apiURL)
	if err != nil {
		log.Fatalf("Error fetching updates: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: %s", resp.Status)
	}

	var result struct {
		Ok     bool `json:"ok"`
		Result []struct {
			Message struct {
				Chat struct {
					ID string `json:"id"`
				} `json:"chat"`
			} `json:"message"`
		} `json:"result"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatalf("Error decoding response: %v", err)
	}

	if !result.Ok {
		log.Fatalf("Error: response not ok")
	}

	for _, update := range result.Result {
		chatID := update.Message.Chat.ID
		t.AddChatID(chatID)
	}
}
