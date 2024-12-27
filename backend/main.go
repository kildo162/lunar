package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Bot token từ BotFather
	botToken := "7736088830:AAHNO8cfkPcGgJLowfnFfq6laoNIaBehBaY"
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

	// ID người nhận tin nhắn (chat_id)
	chatID := "1020039794" // Thay bằng chat ID của bạn

	// Nội dung tin nhắn
	messageText := "🚨 CẢNH BÁO: Hệ thống gặp sự cố! Xin vui lòng chờ đội ngũ kỹ thuật xử lý."

	// Tạo payload để gửi request
	payload := map[string]string{
		"chat_id": chatID,
		"text":    messageText,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error encoding payload: %v", err)
	}

	// Gửi HTTP POST request
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

	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		var update struct {
			Message struct {
				Text string `json:"text"`
				Chat struct {
					ID int64 `json:"id"`
				} `json:"chat"`
			} `json:"message"`
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatalf("Error reading request body: %v", err)
		}
		err = json.Unmarshal(body, &update)
		if err != nil {
			log.Fatalf("Error decoding JSON: %v", err)
		}

		if update.Message.Text == "/hello" {
			replyMessage(apiURL, update.Message.Chat.ID, "Xin chào, bạn cần giúp gì?")
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func replyMessage(apiURL string, chatID int64, messageText string) {
	payload := map[string]interface{}{
		"chat_id": chatID,
		"text":    messageText,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error encoding payload: %v", err)
	}

	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Lỗi: %s\n", resp.Status)
	}
}
