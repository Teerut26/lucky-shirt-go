package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}
	client := &http.Client{}
	bot, err := messaging_api.NewMessagingApiAPI(
		os.Getenv("CHANNEL_ACCESS_TOKEN"),
		messaging_api.WithHTTPClient(client),
	)

	if err != nil {
		panic(err)
	}

	// broadcast message
	broadcastMessage := &messaging_api.BroadcastRequest{
		Messages: []messaging_api.MessageInterface{
			messaging_api.TextMessage{
				Text: "asdfasdf",
			},
		},
	}
	_, err = bot.Broadcast(broadcastMessage, "")
	if err != nil {
		panic(err)
	}
}
