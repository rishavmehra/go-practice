package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	discordToken := os.Getenv("DISCORD_TOKEN")

	dg, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		fmt.Println("error creating Discord session: ", err)
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection: ", err)
	}

	fmt.Println("Bot is running now. Press ctrl + c to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content != "ping" {
		return
	}

	channel, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		fmt.Println("error creaeting channnel: ", err)
		s.ChannelMessageSend(
			m.ChannelID,
			"something went wrong while sending the DM",
		)
		return
	}
	_, err = s.ChannelMessageSend(channel.ID, "PONG!")

	if err != nil {
		fmt.Println("error sending DM message: ", err)
		s.ChannelMessageSend(
			m.ChannelID,
			"Failed to send you a DM."+
				"Did you disable the DM in your setting",
		)
	}

}
