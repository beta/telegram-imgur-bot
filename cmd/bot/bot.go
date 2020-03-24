package main

import (
	"log"

	"github.com/beta/imgur-bot/bot"
)

func main() {
	if err := bot.Start(); err != nil {
		log.Fatalf("[bot] %v", err)
	}
}
