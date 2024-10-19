package types

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type WebhookSink struct {
	KippySink
}

func (sink *WebhookSink) Send(messages []KippyMessage) error {
	jsonStr, err := json.Marshal(messages)
	if err != nil {
		log.Printf("Error marshalling messages: %v\n", err)
		return err
	}
	req, err := http.NewRequest("POST", sink.Config, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Printf("Error creating request: %v\n", err)
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending webhook message: %v\n", err)
	}
	defer resp.Body.Close()
	log.Println("Sending Webhook message")
	return nil
}
