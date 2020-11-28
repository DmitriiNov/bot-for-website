package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

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

func SendMessage(bot *tgbotapi.BotAPI, line string) {
	msg := tgbotapi.NewMessage(chatID, line)
	bot.Send(msg)
}

func Run(token string) error {

	fmt.Println("Hello")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}

	bot.Debug = true

	fmt.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		if !checkForIDs(update.Message.From.ID) {
			continue
		}

		SendMessage(bot, "Аркаша пидрила")
		// log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		// msg.ReplyToMessageID = update.Message.MessageID

		// bot.Send(msg)
	}

	return nil
}
