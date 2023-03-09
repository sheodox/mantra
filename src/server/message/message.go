package message

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Messager struct {
	discordWebhookUrl string
}

func NewMessager(webhookUrl string) Messager {
	return Messager{
		discordWebhookUrl: webhookUrl,
	}
}

type messagePayload struct {
	Content string `json:"content"`
}

func (m Messager) SendMessage(msg string) {
	go func() {
		if msg == "" {
			// don't send empty messages
			log.Print("Tried to send an empty message, ignoring.")
			return
		}

		p := messagePayload{msg}
		payload, err := json.Marshal(p)

		if err != nil {
			log.Printf("Error marshalling json for Discord message payload. msg: %q\n%v\n", msg, err)
		}

		resp, err := http.Post(m.discordWebhookUrl, "application/json", bytes.NewBuffer(payload))

		if err != nil {
			log.Println("Error encountered trying to send message")
		} else if resp.StatusCode > 299 {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Println("Error reading body: ", err)
			}
			resp.Body.Close()

			log.Printf("Error sending discord message: %v\n%v\n", resp.Status, string(body))
		} else {
			log.Println("Successfully sent message")
		}
	}()
}
