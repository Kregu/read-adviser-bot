package read_adviser_bot

import (
	"flag"
	"log"
	tgClient "read-adviser-bot/clients/telegram"
	"read-adviser-bot/consumer/event_consumer"
	"read-adviser-bot/storage/files"
	"read-adviser-bot/events/telegram"


)

const (
	tgBotHost = "api.telegram.org"
	storagePath = "storage"
	baseSize = 100
)

func main() {


	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer:=event_consumer.New(eventsProcessor, eventsProcessor, baseSize)
	if err:= consumer.Start(); err!=nil {
		log.Fatal("service is stopped", err)
	}

}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token os not specified")
	}

	return *token
}
