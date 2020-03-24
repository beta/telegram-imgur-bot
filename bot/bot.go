package bot

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/beta/imgur-bot/bot/apis"
	"github.com/beta/imgur-bot/bot/middlewares"

	"gopkg.in/tucnak/telebot.v2"
)

// Start creates and starts the Imgur Telegram bot.
func Start() error {
	// Load API tokens.
	if err := loadTokens(); err != nil {
		return err
	}

	// Init Telegram bot.
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  tokens.Telegram,
		Poller: telebot.NewMiddlewarePoller(&telebot.LongPoller{Timeout: 10 * time.Second}, middlewares.Logger),
	})
	if err != nil {
		return fmt.Errorf("failed to initialize Telegram bot, error: %v", err)
	}
	route(bot)

	// Start bot.
	bot.Start()
	return nil
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
	api := apis.WithBot(bot)
	bot.Handle("/start", api.Hello)

	bot.Handle(telebot.OnText, api.Unsupported)
}
