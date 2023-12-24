package main

import (
	"echobot/parser"
	"echobot/repository"
	"echobot/telegram"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func main() {
	file, err := os.OpenFile("data/logs/logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.Println("bot started")
	// Настроим логирование на вывод в файл
	log.SetOutput(file)

	if err := godotenv.Load(); err != nil {
		log.Printf("error loading env variables: %s", err.Error())
		return
	}
	// Получите токен вашего бота из окружения или укажите его здесь
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Panic(err.Error())
	}

	// bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	parser := parser.NewParser()
	db, err := repository.NewSqlite()
	if err != nil {
		log.Panic(err.Error())
	}
	if err := db.Init(); err != nil {
		log.Panic(err.Error())
	}

	telegramBot := telegram.NewBot(bot, parser, db)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
