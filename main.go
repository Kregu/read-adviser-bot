package main

import (
	"flag"
	"log"
	tgClient "read-adviser-bot/clients/telegram"
	"read-adviser-bot/consumer/event_consumer"
	"read-adviser-bot/events/telegram"
	"read-adviser-bot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	baseSize    = 100
)

func main() {

	storage := files.New(storagePath)

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		storage,
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, baseSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}

}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token os not specified")
	}

	return *token
}
