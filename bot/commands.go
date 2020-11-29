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
	line := fmt.Sprintf("Имя: %s\nУслуга: %s\nОписание: %s\n", form.Name, form.Service, form.Info)
	SendMessage(line)
}
