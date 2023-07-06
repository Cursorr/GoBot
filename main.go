package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Cursorr/gobot/events"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

type Bot struct {
	Session *discordgo.Session
}

func (bot *Bot) RegisterHandlers() {
	var eventHandlers = []interface{}{
		events.OnReady,
		events.OnGuildJoin,
		events.OnGuildRemove,
		events.OnInviteCreate,
		events.OnInviteDelete,
		events.OnMemberJoin,
		events.OnMemberRemove,
	}

	for _, event := range eventHandlers {
		bot.Session.AddHandler(event)
	}

}

func (bot *Bot) Start() {
	godotenv.Load()
	token := os.Getenv("TOKEN")

	session, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatal(err)
		return
	}

	session.Identify.Intents = discordgo.IntentsAll

	bot.Session = session

	bot.RegisterHandlers()

	err = session.Open()

	if err != nil {
		log.Fatal(err)
		return
	}
	defer session.Close()

	fmt.Println("Bot en ligne.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func main() {
	bot := Bot{}
	bot.Start()
}