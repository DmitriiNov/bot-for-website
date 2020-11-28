package main

import (
	"bot-for-website/bot"
	"log"
)

func main() {
	log.Println("Starting server...")
	err := StartServer()
	log.Println("Server started!")
	if err != nil {
		panic(err)
	}
	err = bot.Run(BOT_TOKEN)
	if err != nil {
		panic(err)
	}
}
