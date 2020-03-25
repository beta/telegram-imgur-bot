# Telegram Imgur Bot

[![Telegram: @beta_imgur_bot](https://img.shields.io/badge/Telegram-@beta__imgur__bot-%232CA5E0?logo=telegram "Telegram: @beta_imgur_bot")](https://t.me/beta_imgur_bot) [![Build status](https://img.shields.io/github/workflow/status/beta/telegram-imgur-bot/Go "Build status")](https://github.com/beta/telegram-imgur-bot/actions?query=workflow%3AGo)

A Telegram bot which uploads images sent from you to Imgur and replies with direct links. Useful for sharing images, writing blogs, etc.

## Screenshot

<p align="center"><img src="screenshot.png?raw=true" alt="Screenshot" title="Screenshot" /></p>
<p align="center"><sup>Image by Negative Space via <a href="https://www.pexels.com/photo/coffee-notebook-writing-computer-34601/">Pexels</a> (CC0 license)</sup></p>

## Prerequisites

- Go >= 1.11
- A Telegram bot created from [BotFather](https://t.me/BotFather)
- Imgur API client created following the guide at [apidocs.imgur.com](https://apidocs.imgur.com/)

## Getting started

```bash
$ git clone https://github.com/beta/telegram-imgur-bot.git
$ go get ./...
$ TELEGRAM_BOT_TOKEN=[YOUR_BOT_TOKEN] IMGUR_CLIENT_ID=[YOUR_IMGUR_CLIENT_ID] go run cmd/bot/bot.go
```

## Deploying to Heroku

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

### Running with free dynos

> Heroku shuts down free dynos after there is no traffic in a period. A web server is added ([cmd/web/web.go](cmd/web/web.go)) which supports [wakemydyno.com](http://wakemydyno.com/). Register your Heroku app there if you want to prevent it from sleeping.

## To-dos

- [ ] Support logging in to Imgur to upload with users' own accounts
- [ ] Support specifying which album to upload to
- [ ] Inline keyboard for deleting images from Imgur

## Credits

- [Telebot](https://github.com/tucnak/telebot)

## License

MIT
