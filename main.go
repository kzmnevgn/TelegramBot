package main

import (
	"flag"
	"log"

	"project/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	t := mustToken()

	tgClient = telegram.New(mustToken())

	//fetcher = fetcher.New()

	//processor = processor.New()

	//consumer.Start(fetcher, processor)
}

func mustToken() string {
	// bot -t-bot-token 'my token'
	token := flag.String(
		"token-bot-token", //name
		"", //value
		"token for access to telegram bot" //usage
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}