module github.com/beta/telegram-imgur-bot

// +heroku goVersion 1.14
// +heroku install ./cmd/...

go 1.14

require (
	github.com/lib/pq v1.3.0
	github.com/pkg/errors v0.9.1 // indirect
	gopkg.in/tucnak/telebot.v2 v2.0.0-20200301001213-9852df39ae6c
)
