package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/beta/imgur-bot/cmds"
	"gopkg.in/tucnak/telebot.v2"
)

func main() {
	// Load API tokens.
	if err := loadTokens(); err != nil {
		log.Fatalf("[main] %v", err)
		return
	}

	// Init Telegram bot.
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  tokens.Telegram,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatalf("[main] failed to initialize Telegram bot, error: %v", err)
		return
	}
	route(bot)

	// Start bot.
	bot.Start()
}

var tokens struct {
	Telegram string
	Imgur    string
}

const (
	envTelegramToken = "TELEGRAM_API_TOKEN"
	envImgurToken    = "IMGUR_API_TOKEN"
)

func loadTokens() error {
	// Internal func for reading 1 token from env var.
	var loadToken = func(envVar string) (string, error) {
		token, ok := os.LookupEnv(envVar)
		token = strings.TrimSpace(token)
		if !ok || len(token) <= 0 {
			return "", fmt.Errorf("environment variable %s is not properly set", envVar)
		}
		return token, nil
	}

	// Read Telegram API token.
	telegramToken, err := loadToken(envTelegramToken)
	if err != nil {
		return err
	}
	// Read Imgur API token.
	imgurToken, err := loadToken(envImgurToken)
	if err != nil {
		return err
	}

	// Save to tokens.
	tokens.Telegram = telegramToken
	tokens.Imgur = imgurToken
	return nil
}

func route(bot *telebot.Bot) {
	bot.Handle("/hello", cmds.Hello(bot))
}
