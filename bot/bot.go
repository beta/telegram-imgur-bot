package bot

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/beta/imgur-bot/bot/apis"
	"github.com/beta/imgur-bot/bot/imgur"
	"github.com/beta/imgur-bot/bot/middlewares"

	"gopkg.in/tucnak/telebot.v2"
)

// Start creates and starts the Imgur Telegram bot.
func Start() error {
	// Load API keys.
	if err := loadAPIKeys(); err != nil {
		return err
	}

	// Init Imgur client.
	imgur.Init(apiKeys.ImgurClientID)

	// Init Telegram bot.
	bot, err := telebot.NewBot(telebot.Settings{
		Token: apiKeys.TelegramBotToken,
	})
	bot.Poller = telebot.NewMiddlewarePoller(&telebot.LongPoller{Timeout: 10 * time.Second}, middlewares.Logger(bot))
	if err != nil {
		return fmt.Errorf("failed to initialize Telegram bot, error: %v", err)
	}
	route(bot)

	// Start bot.
	bot.Start()
	return nil
}

var apiKeys struct {
	TelegramBotToken string
	ImgurClientID    string
}

const (
	envTelegramBotToken = "TELEGRAM_BOT_TOKEN"
	envImgurClientID    = "IMGUR_CLIENT_ID"
)

func loadAPIKeys() error {
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
	telegramBotToken, err := loadToken(envTelegramBotToken)
	if err != nil {
		return err
	}
	// Read Imgur API token.
	imgurClientID, err := loadToken(envImgurClientID)
	if err != nil {
		return err
	}

	// Save to apiKeys.
	apiKeys.TelegramBotToken = telegramBotToken
	apiKeys.ImgurClientID = imgurClientID
	return nil
}

func route(bot *telebot.Bot) {
	api := apis.WithBot(bot)
	bot.Handle("/start", messageHandler(api.Hello))

	bot.Handle(telebot.OnPhoto, messageHandler(api.Photo))
	bot.Handle(telebot.OnDocument, messageHandler(api.File))

	bot.Handle(telebot.OnText, func(m *telebot.Message) {
		api.Unsupported(&apis.Message{Message: m})
	})
}

func messageHandler(handler func(*apis.Message)) func(*telebot.Message) {
	return func(m *telebot.Message) {
		handler(&apis.Message{Message: m})
	}
}

func callbackHandler(handler func(*apis.Callback)) func(*telebot.Callback) {
	return func(cb *telebot.Callback) {
		handler(&apis.Callback{Callback: cb})
	}
}
