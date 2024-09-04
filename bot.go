package tgbot

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/go-telegram/bot"
)

type Bot struct {
	b *bot.Bot

	ctx       context.Context
	ctxcancel context.CancelFunc

	OnMsg      OnTextMessage
	OnBtnClick OnButton

	onlineUsers   map[int64]*User //chatid => user
	onlineUsers_m sync.Mutex
}

type User struct {
	Id   int64
	Nick string
}

type OnTextMessage func(chatId int64, user *User, text string)
type OnButton func(chatId int64, user *User, name string)

func New(token string) (*Bot, error) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	thisBot := &Bot{
		ctx:         ctx,
		ctxcancel:   cancel,
		onlineUsers: make(map[int64]*User),
	}

	//on text message
	defaultOnTextMessage := func(chatId int64, user *User, text string) {
		thisBot.SendMessage(chatId, fmt.Sprintf("message from %s (%d) : %s", user.Nick, user.Id, text))
	}
	thisBot.OnMsg = defaultOnTextMessage

	//on button click
	defaultOnButton := func(chatId int64, user *User, name string) {
		thisBot.SendMessage(chatId, fmt.Sprintf("button click from %s (%d) : %s", user.Nick, user.Id, name))
	}
	thisBot.OnBtnClick = defaultOnButton

	opts := []bot.Option{
		bot.WithDefaultHandler(onTextHandler(thisBot)),
		bot.WithCallbackQueryDataHandler("button", bot.MatchTypePrefix, callbackHandler(thisBot)),
	}

	b, err := bot.New(token, opts...)
	if err != nil {
		return nil, err
	}

	thisBot.b = b

	go func() {
		b.Start(ctx)
	}()

	return thisBot, nil
}
func (b *Bot) Close() {
	b.ctxcancel()
}
