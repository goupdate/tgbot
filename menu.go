package tgbot

import (
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type BotCommand = models.BotCommand

func (b *Bot) SetMenu(cmds []BotCommand) {
	b.b.SetMyCommands(b.ctx, &bot.SetMyCommandsParams{
		Commands: cmds,
		Scope:    &models.BotCommandScopeAllPrivateChats{},
	})
}
