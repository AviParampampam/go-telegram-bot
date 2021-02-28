package handler

import (
	"log"
	"time"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

// Handler ..
type Handler struct {
	Bot *tgbotapi.BotAPI
}

// HandleUpdate ..
func (h *Handler) HandleUpdate(update tgbotapi.Update) {
	var (
		msgText = update.Message.Text
		msgFrom = update.Message.From
	)

	// Logging message
	log.Printf("[%s] %s", msgFrom.UserName, msgText)

	// Enable "typing" status on 1 sec.
	h.Bot.Send(tgbotapi.NewChatAction(update.Message.Chat.ID, "typing"))
	time.Sleep(time.Second)

	switch {
	case isHello(msgText):
		h.helloResponse(update)
	default:
		h.defaultResponse(update)
	}
}
