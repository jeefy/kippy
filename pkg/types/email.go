package types

import (
	"log"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/spf13/viper"
)

type EmailSink struct {
	KippySink
}

func (sink *EmailSink) Send(messages []KippyMessage) error {
	from := mail.NewEmail("Jeefy", "me@jeefy.dev") // Change to your verified sender
	subject := "kippy Notification: New Events Detected!"
	to := mail.NewEmail(sink.Config, sink.Config) // Change to your recipient

	toSend := "The following events have occurred:\n\n"

	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	table.SetHeader([]string{"Kind", "Namespace", "Name", "Message"})
	for _, m := range messages {
		table.Append([]string{m.Kind, m.Namespace, m.Name, m.Message})
	}

	table.Render()
	toSend += tableString.String()

	message := mail.NewSingleEmail(from, subject, to, toSend, toSend)
	client := sendgrid.NewSendClient(viper.GetString("sendgridApi"))
	_, err := client.Send(message)
	if err != nil {
		log.Printf("Error sending email: %v\n", err)
	}
	log.Println("Sending email to", sink.Config)
	return nil
}
