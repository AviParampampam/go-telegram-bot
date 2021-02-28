package handler

import (
	"math/rand"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

var defaultMessages = []string{
	"Я не знаю такой команды :С", "Не понимаю о чем ты..", "Что?",
	"Что. Это. За. Слово.", "м?", "Кажется, ты ошибься командой", "Прочитай-ка получше инструкцию",
}

func (h *Handler) defaultResponse(update tgbotapi.Update) {
	resMsg := defaultMessages[rand.Intn(len(defaultMessages))]
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, resMsg)
	h.Bot.Send(msg)
}
