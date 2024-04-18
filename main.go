package main

import (
	"flag"
	"log"
)

func main() {
	t := mustToken()

	//token = flags.Get(token)

	//tgClient = telegram.New(token)

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