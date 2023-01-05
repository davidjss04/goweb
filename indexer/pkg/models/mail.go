package models

import (
	"fmt"
	"os"
	"sort"

	"github.com/jhillyerd/enmime"
)

type Mail struct {
	MessageId string `json:"message_id"`
	From      string `json:"from"`
	To        string `json:"to"`
	Content   string `json:"content"`
}

type Index struct {
	Index   string  `json:"index"`
	Records *[]Mail `json:"records"`
}

func NewMail(message_id string, from string, to string, content string) Mail {
	newMail := Mail{}
	newMail.MessageId = message_id
	newMail.From = from
	newMail.To = to
	newMail.Content = content
	return newMail
}

func NewIndex(index string) Index {
	newIndex := Index{}
	newIndex.Index = index
	newIndex.Records = &[]Mail{}
	return newIndex
}


func ReadMail(path string, option string) string {

	r, err := os.Open(path)
	if err != nil {
		fmt.Print(err)
		return ""
	}

	// Parse message with enmime.
	env, err := enmime.ReadEnvelope(r)
	if err != nil {
		fmt.Print(err)
		return ""
	}

	// A list of headers is retrieved via Envelope.GetHeaderKeys().
	headers := env.GetHeaderKeys()
	sort.Strings(headers)

	r.Close()

	switch option {
	case "message-id":
		return env.GetHeader("Message-ID")
	case "from":
		return env.GetHeader("From")
	case "to":
		if env.GetHeader("To") == "" {
			return "undiagnosed recipient"
		}

		return env.GetHeader("To")
	case "content":
		return env.Text
	default:
		return ""
	}
}


