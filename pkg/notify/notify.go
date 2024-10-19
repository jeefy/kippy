package notify

import (
	"log"
	"sync"

	"github.com/jeefy/kippy/pkg/config"
	"github.com/jeefy/kippy/pkg/types"
	"github.com/spf13/viper"
	v1 "k8s.io/api/core/v1"
)

func SendNotification(HeartBeat *types.KippyHeartbeat, HeartBeatLock *sync.Mutex) error {
	var filteredMessages []types.KippyMessage
	HeartBeatLock.Lock()
	for _, event := range HeartBeat.Events {
		// Filter out messages that are duplicates
		if event.Type != "Normal" {
			filteredMessages = append(filteredMessages, types.KippyMessage{
				Kind:      event.Kind,
				Name:      event.Name,
				Namespace: event.Namespace,
				Message:   event.Message,
				Timestamp: event.FirstTimestamp.Time,
			})
			log.Printf("%v", event)
		}
	}

	// Clear the events after we've filtered them
	HeartBeat.Events = []*v1.Event{}

	HeartBeatLock.Unlock()

	// Let's not send messages unless there are actually messages
	if len(filteredMessages) == 0 {
		return nil
	}

	log.Printf("Sending %d messages\n", len(filteredMessages))

	for _, sink := range config.Sinks {
		if sink.Config == "" {
			if viper.GetBool("debug") {
				log.Printf("Sink %s has no configuration\n", sink.Type)
			}
			continue
		}
		switch sink.Type {
		case "email":
			emailSink := types.EmailSink{KippySink: sink}
			emailSink.Send(filteredMessages)
		case "slack":
			slackSink := types.SlackSink{KippySink: sink}
			slackSink.Send(filteredMessages)
		case "webhook":
			webhookSink := types.WebhookSink{KippySink: sink}
			webhookSink.Send(filteredMessages)
		case "discord":
			discordSink := types.DiscordSink{KippySink: sink}
			discordSink.Send(filteredMessages)
		}
	}

	return nil
}
