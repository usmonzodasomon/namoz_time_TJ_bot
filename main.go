package main

import (
	"context"
	"echobot/app"
	"echobot/handler"
	"echobot/parser"
	"echobot/pkg/database"
	"echobot/scheduler"
	"echobot/storage/postgres"
	"echobot/telegram"
	"fmt"
	"github.com/go-co-op/gocron/v2"
	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
	"io"
	"log"
	"os"
	"os/signal"
)

func main() {
	file, err := os.OpenFile("data/logs/logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.Println("bot started")
	// Настроим логирование на вывод в файл и консоль одновременно
	log.SetOutput(io.MultiWriter(os.Stdout, file))

	if err := godotenv.Load(); err != nil {
		log.Printf("error loading env variables: %s", err.Error())
		return
	}

	dbConn, err := database.GetDBConnection(database.Config{
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
	})
	if err != nil {
		log.Println(fmt.Errorf("error connecting to db: %w", err))
		return
	}

	storage := postgres.NewPostgresStorage(dbConn)
	handler := handler.NewHandler(storage)
	parser := parser.NewParser()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	opts := []bot.Option{
		bot.WithDefaultHandler(handler.DefaultHandler),
		//bot.WithCallbackQueryDataHandler("region", bot.MatchTypePrefix, handler.CallbackHandler),
	}

	b, err := bot.New(os.Getenv("BOT_TOKEN"), opts...)
	if err != nil {
		panic(err)
	}
	tl := telegram.NewBot(b, handler)
	sh, err := gocron.NewScheduler()
	if err != nil {
		log.Println(err)
	}
	scheduler := scheduler.NewScheduler(parser, sh, storage, tl)
	app := app.NewApp(tl, scheduler)
	go app.Start(ctx)
	<-ctx.Done()
	log.Println("bot stopped")
}
