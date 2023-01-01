package models

import (
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

func NewIndex(index string) Index {
	newIndex := Index{}
	newIndex.Index = index
	newIndex.Records = &[]Mail{}
	return newIndex
}

func NewMail(message_id string, from string, to string, content string) Mail {
	newMail := Mail{}
	newMail.MessageId = message_id
	newMail.From = from
	newMail.To = to
	newMail.Content = content
	return newMail
}

