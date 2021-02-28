package handler

import (
	"math/rand"
	"strings"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

var helloMessages = []string{
	"hello", "hi", "привет", "hola", "hey", "здарова", "здравствуй", "здравствуйте", "алоха",
	"здравия желанию", "добрый день", "добрый вечер", "доброе утро", "ку", "прив", "хеллоу", "дароу",
	"bonjour", "банжур", "бонжур", "приветствую", "re",
}

func isHello(msgText string) bool {
	for _, h := range helloMessages {
		if h == strings.ToLower(msgText) {
			return true
		}
	}
	return false
}

func (h *Handler) helloResponse(update tgbotapi.Update) {
	resMsg := helloMessages[rand.Intn(len(helloMessages))]
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, strings.Title(resMsg))

	if len(*h.UpdatesChannel) > 0 {
		msg.ReplyToMessageID = update.Message.MessageID
	}

	h.Bot.Send(msg)
}
