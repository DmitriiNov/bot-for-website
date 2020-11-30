package bot

import (
	"bot-for-website/database"
	"fmt"
)

func SendLast10() {

}

func SendSQLiteDB() {

}

func SendForm(form database.Form) {
	line := fmt.Sprintf("Имя: %s\nУслуга: %s\nОписание: %s\n\n", form.Name, form.Service, form.Info)
	for _, val := range form.Contacts {
		line = fmt.Sprintf("%sCпособ связи: %s\nКонтакт: %s\n\n", line, val.Kind, val.Info)
	}
	SendMessage(line)
}
