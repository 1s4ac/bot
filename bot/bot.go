package main

import (
	"botpack/handlers"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

func main() {
	//reads token from text file, stores in token variable
	token, err := ioutil.ReadFile("token/token.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Creates Discord session using bot token
	discord, err := discordgo.New("Bot " + string(token))

	if err != nil {
		fmt.Println("error occured")
		return
	}

	// messageCreate func registered as callback for messageCreate events
	discord.AddHandler(handlers.MessageCreate)
	// ready registered as a callback for the ready events
	discord.AddHandler(ready)

	// Receiving message events
	discord.Identify.Intents = discordgo.IntentsGuildMessages

	// Opens websocket connection to discord
	err = discord.Open()
	if err != nil {
		fmt.Println("error occurred opening connection", err)
		return
	}

	// Runs until ctrl-c signal is received
	fmt.Println("bot online")
	sig := make(chan os.Signal, 1)
	//signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sig

	// Close discord session
	discord.Close()
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	// Set the bot's status
	s.UpdateStreamingStatus(0, "Online!", "https://www.youtube.com/watch?v=AVblOqZBlJw") //sets bot to "streaming"
	s.ChannelMessageSend("731615059758940203", "bot online")                             //you will have to change the first argument

	ticker := time.NewTicker(5 * time.Second)
	for _ = range ticker.C {
		fact := handlers.RandFact()
		s.ChannelMessageSend("731615059758940203", fact) //you will have to change the first argument
	}

}
