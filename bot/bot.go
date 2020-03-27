package bot

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/beta/telegram-imgur-bot/bot/apis"
	"github.com/beta/telegram-imgur-bot/bot/db"
	"github.com/beta/telegram-imgur-bot/bot/imgur"
	"github.com/beta/telegram-imgur-bot/bot/middlewares"

	"gopkg.in/tucnak/telebot.v2"
)

// Start creates and starts the Imgur Telegram bot.
func Start() error {
	// Load env vars.
	if err := loadEnvs(); err != nil {
		return err
	}

	// Init DB.
	if err := db.Init(configs.DSN); err != nil {
		return fmt.Errorf("failed to initialize DB connection, error: %v", err)
	}

	// Init Imgur client.
	imgur.Init(configs.ImgurClientID)

	// Init Telegram bot.
	bot, err := telebot.NewBot(telebot.Settings{
		Token: configs.TelegramBotToken,
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

var configs struct {
	DSN              string
	TelegramBotToken string
	ImgurClientID    string
}

const (
	envDatabaseURL      = "DATABASE_URL"
	envTelegramBotToken = "TELEGRAM_BOT_TOKEN"
	envImgurClientID    = "IMGUR_CLIENT_ID"
)

func loadEnvs() error {
	// Internal func for reading 1 env var.
	var loadEnv = func(envVar string) (string, error) {
		val, ok := os.LookupEnv(envVar)
		val = strings.TrimSpace(val)
		if !ok || len(val) <= 0 {
			return "", fmt.Errorf("environment variable %s is not properly set", envVar)
		}
		return val, nil
	}

	// Read DSN.
	dsn, err := loadEnv(envDatabaseURL)
	if err != nil {
		return err
	}
	// Read Telegram API token.
	telegramBotToken, err := loadEnv(envTelegramBotToken)
	if err != nil {
		return err
	}
	// Read Imgur API token.
	imgurClientID, err := loadEnv(envImgurClientID)
	if err != nil {
		return err
	}

	// Save to configs.
	configs.DSN = dsn
	configs.TelegramBotToken = telegramBotToken
	configs.ImgurClientID = imgurClientID
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
