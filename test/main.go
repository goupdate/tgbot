package main

import (
	"embed"
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/goupdate/tgbot"
)

//go:embed test.png
var images embed.FS

func main() {
	token := flag.String("token", "", "telegram bot token")
	flag.Parse()

	if *token == "" {
		flag.Usage()
		return
	}

	b, err := tgbot.New(*token)
	if err != nil {
		panic(err.Error())
		return
	}
	defer b.Close()

	b.OnMsg = func(chatId int64, user *tgbot.User, text string) {
		switch text {
		case "/buttons":
			kb := [][]tgbot.InlineKeyboardButton{
				{
					{Text: "Button A", CallbackData: "button_A"},
					{Text: "Button B", CallbackData: "button_B"},
				}, {
					{Text: "Button C", CallbackData: "button_C"},
				},
			}
			b.SendMessageWithButtons(chatId, "button message", kb)
		case "/pic_by_link":
			b.SendPictureByUrl(chatId, "weblink.png", "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT9sKiPVlcEz6FB2JPT36drHd1XKqiYmChoZw&s")
		case "/pic_by_body":
			testpng, _ := images.ReadFile("test.png")
			b.SendPicture(chatId, "localfile.png", testpng)
		case "/broadcast":
			b.Broadcast("this message is broadcasted")
		default:
			b.SendMessage(chatId, fmt.Sprintf("got message %s from %s(%d)", text, user.Nick, user.Id))
		}
	}

	b.OnBtnClick = func(chatId int64, user *tgbot.User, name string) {
		b.SendMessage(chatId, fmt.Sprintf("btn click %s from %s(%d)", name, user.Nick, user.Id))
	}

	b.SetMenu([]tgbot.BotCommand{
		{
			Command:     "buttons",
			Description: "Show buttons",
		},
		{
			Command:     "broadcast",
			Description: "Broadcast message",
		},
		{
			Command:     "some",
			Description: "Some command",
		},
		{
			Command:     "pic_by_link",
			Description: "Send picture by link",
		},
		{
			Command:     "pic_by_body",
			Description: "Send picture by body",
		},
	})

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	//wait ctrl+c
	fmt.Println("wait ctrl+c")
	<-c
}
