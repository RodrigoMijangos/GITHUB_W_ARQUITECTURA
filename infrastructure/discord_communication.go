package infrastructure

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func SendMessageToDiscord(msg string) int {

	discord_webhook_url := os.Getenv("DISCORD_WEBHOOK_URL")

	if discord_webhook_url == "" {
		log.Println("No hay url para discord")
		return 500
	}

	payload := map[string]string{
		"content": msg,
	}

	jsonPayload, err := json.Marshal(payload)

	if err != nil {
		log.Println("Error al serializar json")
		return 500
	}

	resp, err := http.Post(discord_webhook_url, "application/json", bytes.NewBuffer(jsonPayload))

	if err != nil {
		log.Println("Error al mandar el mensaje")
		return 500
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error al mandar mensaje, código: %d", resp.StatusCode)
		return 400
	}

	return 200

}
