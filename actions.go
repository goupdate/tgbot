package tgbot

import (
	"bytes"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// send to all clients
func (b *Bot) Broadcast(text string) {
	for cid, _ := range b.onlineUsers {
		b.SendMessage(cid, text)
	}
}

func (b *Bot) SendMessage(chatId int64, text string) {
	_, err := b.b.SendMessage(b.ctx, &bot.SendMessageParams{
		ChatID: chatId,
		Text:   text,
	})
	if err != nil {
		b.onlineUsers_m.Lock()
		_, ok := b.onlineUsers[chatId]
		if ok {
			delete(b.onlineUsers, chatId)
		}
		b.onlineUsers_m.Unlock()
	}
}

type InlineKeyboardButton = models.InlineKeyboardButton

func (b *Bot) SendMessageWithButtons(chatId int64, text string, kb [][]InlineKeyboardButton) {
	inlinekb := &models.InlineKeyboardMarkup{
		InlineKeyboard: kb,
	}

	_, err := b.b.SendMessage(b.ctx, &bot.SendMessageParams{
		ChatID:      chatId,
		Text:        text,
		ReplyMarkup: inlinekb,
	})
	if err != nil {
		b.onlineUsers_m.Lock()
		_, ok := b.onlineUsers[chatId]
		if ok {
			delete(b.onlineUsers, chatId)
		}
		b.onlineUsers_m.Unlock()
	}
}

func (b *Bot) SendPictureByUrl(chatId int64, name, url string) {
	media1 := &models.InputMediaPhoto{
		Media:   url,
		Caption: name,
	}

	params := &bot.SendMediaGroupParams{
		ChatID: chatId,
		Media: []models.InputMedia{
			media1,
		},
	}

	_, err := b.b.SendMediaGroup(b.ctx, params)
	if err != nil {
		b.onlineUsers_m.Lock()
		_, ok := b.onlineUsers[chatId]
		if ok {
			delete(b.onlineUsers, chatId)
		}
		b.onlineUsers_m.Unlock()
	}
}

func (b *Bot) SendPicture(chatId int64, name string, body []byte) {
	media1 := &models.InputMediaPhoto{
		Media:           "attach://some.png",
		Caption:         name,
		MediaAttachment: bytes.NewReader(body),
	}

	params := &bot.SendMediaGroupParams{
		ChatID: chatId,
		Media: []models.InputMedia{
			media1,
		},
	}

	_, err := b.b.SendMediaGroup(b.ctx, params)
	if err != nil {
		b.onlineUsers_m.Lock()
		_, ok := b.onlineUsers[chatId]
		if ok {
			delete(b.onlineUsers, chatId)
		}
		b.onlineUsers_m.Unlock()
	}
}
