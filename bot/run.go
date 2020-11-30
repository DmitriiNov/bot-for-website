package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var Bot *tgbotapi.BotAPI

var ids = []int{157111809, 155890995}

const chatID = -429497294

func checkForIDs(id int) bool {
	for _, v := range ids {
		if v == id {
			return true
		}
	}
	return false
}

func SendMessage(line string) {
	msg := tgbotapi.NewMessage(chatID, line)
	Bot.Send(msg)
}

func Run(token string) error {

	Bot, _ = tgbotapi.NewBotAPI(token)

	Bot.Debug = false

	fmt.Printf("Authorized on account %s\n", Bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := Bot.GetUpdatesChan(u)

	for update := range updates {
		fmt.Println()
		if update.Message == nil {
			continue
		}

		if !checkForIDs(update.Message.From.ID) {
			continue
		}
	}

	return nil
}
