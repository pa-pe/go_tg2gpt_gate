package usecases

import (
	"upserv/src/service/telegram/models"
)

type IMessageProcessor interface {
	RouteMessage(msg models.Message) models.Message
}

type MessageProcessor struct{}

func NewMessageProcessor() *MessageProcessor {
	return &MessageProcessor{}
}

func (e *MessageProcessor) RouteMessage(msg models.Message) models.Message {
	// temporary return echo
	return models.Message{
		ID:     msg.ID,
		Text:   "Echo: " + msg.Text,
		ChatID: msg.ChatID,
	}
}
