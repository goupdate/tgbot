package tgbot

import (
	"bytes"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (b *Bot) SendMessage(chatId int64, text string) {
	b.b.SendMessage(b.ctx, &bot.SendMessageParams{
		ChatID: chatId,
		Text:   text,
	})
}

type InlineKeyboardButton = models.InlineKeyboardButton

func (b *Bot) SendMessageWithButtons(chatId int64, text string, kb [][]InlineKeyboardButton) {
	inlinekb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Button 1", CallbackData: "button_1"},
				{Text: "Button 2", CallbackData: "button_2"},
			}, {
				{Text: "Button 3", CallbackData: "button_3"},
			},
		},
	}

	b.b.SendMessage(b.ctx, &bot.SendMessageParams{
		ChatID:      chatId,
		Text:        text,
		ReplyMarkup: inlinekb,
	})
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
		fmt.Printf("err: %s\n", err.Error())
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
		fmt.Printf("err: %s\n", err.Error())
	}
}
