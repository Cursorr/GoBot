package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Cursorr/gobot/utils"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	token := os.Getenv("TOKEN")

	session, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatal(err)
		return
	}

	session.Identify.Intents = discordgo.IntentsAll

	utils.RegisterEvents(session)

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