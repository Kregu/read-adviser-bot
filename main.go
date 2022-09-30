package read_adviser_bot

import (
	"flag"
	"log"
	"read-adviser-bot/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {

	tgClient = telegram.New(mustToken())

	// fetcher = fetcher.New()

	// processor = processor.New()

	// consumer.Start(fetcher, processor)

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
