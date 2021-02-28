package app

import (
	"github.com/AviParampampam/go-telegram-bot/internal/bot"
	"github.com/AviParampampam/go-telegram-bot/internal/config"
	"github.com/BurntSushi/toml"
)

// Start is a program entry point.
func Start(cfgpath string) {
	var cfg config.Config
	if _, err := toml.DecodeFile(cfgpath, &cfg); err != nil {
		panic(err)
	}

	bot, err := bot.New(cfg.TelegramBot.Token)
	if err != nil {
		panic(err)
	}

	if err := bot.Listen(); err != nil {
		panic(err)
	}
}
