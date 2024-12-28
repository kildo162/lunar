package shared

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
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

func (t *Telegram) SendMessage(message string) error {
	if len(t.ChatIDs) == 0 {
		return fmt.Errorf("no chat IDs available")
	}

	payload := map[string]string{
		"chat_id": t.ChatIDs[0],
		"text":    message,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error encoding payload: %v", err)
	}

	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", t.BotToken)
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error: %s", resp.Status)
	}

	log.Println("Message sent successfully!")
	return nil
}

func (t *Telegram) SendMessageToChatID(chatID, message string) error {
	payload := map[string]string{
		"chat_id": chatID,
		"text":    message,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error encoding payload: %v", err)
	}

	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", t.BotToken)
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error: %s", resp.Status)
	}

	log.Println("Message sent successfully!")
	return nil
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

func (t *Telegram) SetWebhook(url string) error {
	payload := map[string]string{
		"url": url,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error encoding payload: %v", err)
	}

	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/setWebhook", t.BotToken)
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error: %s", resp.Status)
	}

	return nil
}

func (t *Telegram) InitChatIDsFromEnv() {
	chatIDs := os.Getenv("TELEGRAM_CHAT_IDS")
	if chatIDs == "" {
		log.Println("No chat IDs found in environment variables")
		return
	}
	t.ChatIDs = strings.Split(chatIDs, ",")
	log.Printf("Initialized chat IDs from environment variables: %v", t.ChatIDs)
}

type TelegramUpdate struct {
	UpdateID int `json:"update_id"`
	Message  struct {
		MessageID int `json:"message_id"`
		From      struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Username  string `json:"username"`
		} `json:"from"`
		Chat struct {
			ID   int    `json:"id"`
			Type string `json:"type"`
		} `json:"chat"`
		Date     int    `json:"date"`
		Text     string `json:"text"`
		Entities []struct {
			Type   string `json:"type"`
			Offset int    `json:"offset"`
			Length int    `json://length"`
		} `json:"entities"`
	} `json:"message"`
}
