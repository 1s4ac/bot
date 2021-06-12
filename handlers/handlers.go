package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

// messageCreate function called from handler, everytime a messsage is sent to a channel
// the bot can access

var factState bool

//MessageCreate is
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignores if message is sent by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!fact on" {
		factState = true
		s.ChannelMessageSend(m.ChannelID, "facts turned on")
	}

	if m.Content == "!fact off" {
		factState = false
		s.ChannelMessageSend(m.ChannelID, "facts turned off")
	}

	switch m.Content {
	case "!help", "!commands", "!HELP", "!COMMANDS":
		s.ChannelMessageSend(m.ChannelID, "> use `!fact on/off` to toggle random facts")
	}

}

//RandFact is
func RandFact() string {
	if factState == true {

		response, err := http.Get("https://uselessfacts.jsph.pl/random.json?language=en")

		if err != nil {
			fmt.Printf("error occured, %s", err)
		}

		content, err := ioutil.ReadAll(response.Body)

		if err != nil {
			fmt.Println(err)
		}

		fact := make(map[string]interface{})
		err = json.Unmarshal([]byte(content), &fact)
		if err != nil {
			fmt.Println("error", err.Error())
		}
		randFact := fact["text"]
		return randFact.(string)
	}
	return ""
}
