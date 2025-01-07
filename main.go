package main

import (
	"os"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func main() {
	bot, err := messaging_api.NewMessagingApiAPI(
		os.Getenv("LINE_CHANNEL_TOKEN"),
	)

	if err != nil {
		panic(err)
	}

}
