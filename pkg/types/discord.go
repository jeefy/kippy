package types

import (
	"fmt"
	"log"

	"github.com/gtuk/discordwebhook"
)

type DiscordSink struct {
	KippySink
}

func (sink *DiscordSink) Send(messages []KippyMessage) error {
	message := "The following events have occurred:\n\n"

	for _, m := range messages {
		message += fmt.Sprintf("%s %s/%s\n`%s`\n\n", m.Kind, m.Namespace, m.Name, m.Message)
	}

	log.Println("Sending Discord message")
	var username = "kippy"

	channelMessage := discordwebhook.Message{
		Username: &username,
		Content:  &message,
	}

	err := discordwebhook.SendMessage(sink.Config, channelMessage)
	if err != nil {
		log.Printf("Error sending Discord message: %v\n", err)
	}
	return nil
}
