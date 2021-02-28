package bot

import (
	"fmt"
	"log"

	"github.com/AviParampampam/go-telegram-bot/internal/handler"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

// Bot ..
type Bot struct {
	bot     *tgbotapi.BotAPI
	handler *handler.Handler
}

// New ..
func New(token string) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	// bot.Debug = true
	log.Printf("authorized on account %s", bot.Self.UserName)

	return &Bot{
		bot:     bot,
		handler: &handler.Handler{Bot: bot},
	}, err
}

// Listen ..
func (b *Bot) Listen() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	// Начинаем прослушивание канала
	for update := range updates {
		fmt.Println(len(updates))
		if update.Message == nil { // игнорировать если нет обновлений
			continue
		}
		b.handler.HandleUpdate(update)
	}

	return nil
}
