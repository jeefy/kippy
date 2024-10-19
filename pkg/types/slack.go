package types

import (
	"fmt"
	"log"
	"strings"

	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/olekukonko/tablewriter"
)

type SlackSink struct {
	KippySink
}

func (sink *SlackSink) Send(messages []KippyMessage) error {
	message := "The following events have occurred:\n\n"

	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	table.SetHeader([]string{"Kind", "Namespace", "Name", "Message"})
	for _, m := range messages {
		table.Append([]string{m.Kind, m.Namespace, m.Name, m.Message})
	}

	table.Render()
	message += tableString.String()

	attachments := slack.Attachment{}
	/*
		attachment1.AddField(slack.Field{Title: "Author", Value: "Ashwanth Kumar"}).AddField(slack.Field{Title: "Status", Value: "Completed"})
		attachment1.AddAction(slack.Action{Type: "button", Text: "Book flights 🛫", Url: "https://flights.example.com/book/r123456", Style: "primary"})
		attachment1.AddAction(slack.Action{Type: "button", Text: "Cancel", Url: "https://flights.example.com/abandon/r123456", Style: "danger"})
	*/
	payload := slack.Payload{
		Text:        message,
		Username:    "kippy",
		IconEmoji:   ":fire:",
		Attachments: []slack.Attachment{attachments},
	}
	err := slack.Send(sink.Config, "", payload)
	if len(err) > 0 {
		fmt.Printf("error: %s\n", err)
	}
	log.Println("Sending Slack message")
	return nil
}
