package controller

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"upserv/src/service/telegram/models"
	usecases "upserv/src/service/telegram/service"
)

type TelegramController struct {
	botAPI    *tgbotapi.BotAPI
	processor usecases.IMessageProcessor
	BotInfo   *tgbotapi.User
}

func NewTelegramController(botToken string, processor usecases.IMessageProcessor) (*TelegramController, error) {
	newBotAPI, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return nil, err
	}

	return &TelegramController{
		botAPI:    newBotAPI,
		processor: processor,
	}, nil
}

func (c *TelegramController) CheckConnection() error {
	return c.FetchBotInfo()
}

func (c *TelegramController) FetchBotInfo() error {
	botUser, err := c.botAPI.GetMe()
	if err != nil {
		return err
	}

	c.BotInfo = &botUser
	//	log.Printf("Bot Info: ID=%d, Username=%s, FirstName=%s, LastName=%s, LanguageCode=%s", c.BotInfo.ID, c.BotInfo.UserName, c.BotInfo.FirstName, c.BotInfo.LastName, c.BotInfo.LanguageCode)
	return nil
}

func (c *TelegramController) ListenForMessages() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := c.botAPI.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			msg := models.Message{
				ID:     int64(update.Message.MessageID),
				Text:   update.Message.Text,
				ChatID: update.Message.Chat.ID,
			}

			responseMsg := c.processor.RouteMessage(msg)

			response := tgbotapi.NewMessage(responseMsg.ChatID, responseMsg.Text)
			_, err := c.botAPI.Send(response)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
