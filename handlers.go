package tgbot

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func callbackHandler(tgb *Bot) func(ctx context.Context, b *bot.Bot, update *models.Update) {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		user := &User{
			Id:   update.CallbackQuery.From.ID,
			Nick: update.CallbackQuery.From.Username,
		}

		tgb.onlineUsers_m.Lock()
		tgb.onlineUsers[update.CallbackQuery.Message.Message.Chat.ID] = user
		tgb.onlineUsers_m.Unlock()

		// answering callback query first to let Telegram know that we received the callback query,
		// and we're handling it. Otherwise, Telegram might retry sending the update repetitively
		// as it thinks the callback query doesn't reach to our application. learn more by
		// reading the footnote of the https://core.telegram.org/bots/api#callbackquery type.
		b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})
		//callback on btnclick
		tgb.OnBtnClick(update.CallbackQuery.Message.Message.Chat.ID,
			user,
			update.CallbackQuery.Data)
	}
}

func onTextHandler(tgb *Bot) func(ctx context.Context, b *bot.Bot, update *models.Update) {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		user := &User{
			Id:   update.Message.From.ID,
			Nick: update.Message.From.Username,
		}

		tgb.onlineUsers_m.Lock()
		tgb.onlineUsers[update.Message.Chat.ID] = user
		tgb.onlineUsers_m.Unlock()

		//callback on btnclick
		tgb.OnMsg(update.Message.Chat.ID,
			user,
			update.Message.Text)
	}
}
