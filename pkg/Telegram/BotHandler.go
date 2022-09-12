package telegram

import (
	"fmt"
	typeconv "github.com/adamhoof/GolangTypeConvertorWrapper/pkg"
	tb "gopkg.in/telebot.v3"
	"time"
)

type BotHandler struct {
	Bot   *tb.Bot
	Owner User
}

func (handler *BotHandler) SetToken(token string) {
	var err error
	handler.Bot, err = tb.NewBot(tb.Settings{
		Token: token,
		Poller: &tb.LongPoller{
			Timeout: 10 * time.Second,
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Telegram bot token is valid...")
}

func (handler *BotHandler) StartBot() {
	handler.Bot.Start()
}

func (handler *BotHandler) OwnerVerify(id int64) bool {
	if id != typeconv.StringToInt64(handler.Owner.Id) {
		return false
	}
	return true
}

func (handler *BotHandler) SendText(text string) {
	_, err := handler.Bot.Send(&handler.Owner, text)
	if err != nil {
		fmt.Println(err)
	}
}
