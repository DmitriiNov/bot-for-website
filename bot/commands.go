package bot

import (
	"bot-for-website/database"
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendLast10() {
	forms, ids, err := database.GetLast10()
	if err != nil {
		SendMessage(err.Error())
		return
	}
	line := ""
	for i, v := range forms {
		line += "----------\n"
		line += "ID: " + strconv.Itoa(ids[i]) + "\n"
		line += PrepareForm(v)
		line += "----------\n"
	}
	SendMessage(line)
}

func SendSQLiteDB() {

}

func SendForm(form database.Form) {
	line := PrepareForm(form)
	SendMessage(line)
	for _, v := range form.FileInfo {
		file := tgbotapi.FileBytes{Name: v.Name, Bytes: v.Data.Data}
		SendFile(file)
	}
}

func PrepareForm(form database.Form) (line string) {
	line = fmt.Sprintf("Имя: %s\nУслуга: %s\nОписание: %s\n\n", form.Name, form.Service, form.Info)
	for _, val := range form.Contacts {
		line = fmt.Sprintf("%sCпособ связи: %s\nКонтакт: %s\n\n", line, val.Kind, val.Info)
	}
	return
}
