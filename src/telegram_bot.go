package src

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"upserv/logger"
	"upserv/src/handler"
)

type BotImpl struct {
	token string
}

type TGmsg struct {
	ID     int64
	Text   string
	ChatID int64
}

func NewTelegramBot(token string) *BotImpl {
	return &BotImpl{token: token}
}

func (bot *BotImpl) ListenAndServ() {
	tgController, err := NewTelegramController(bot.token)
	if err != nil {
		//		log.Fatalf("Error creating Telegram tgController: %v", err)
		errMsg := fmt.Sprintf("Error creating Telegram tgController: %v", err)
		logger.LaunchLog(errMsg)
	}

	//	go func() {
	if err := tgController.checkConnection(); err != nil {
		//			log.Fatalf("Failed to connect to Telegram: %v", err)
		errMsg := fmt.Sprintf("Failed to connect to Telegram: %v", err)
		logger.LaunchLog(errMsg)
		return
	}
	logger.LaunchLog("Telegram bot started with UserName=" + tgController.BotInfo.UserName)

	if err := tgController.ListenForMessages(); err != nil {
		//			log.Fatalf("Error listening for messages: %v", err)
		errMsg := fmt.Sprintf("Error listening for messages: %v", err)
		logger.LaunchLog(errMsg)
	}
	//	}()
}

type TelegramController struct {
	botAPI  *tgbotapi.BotAPI
	BotInfo *tgbotapi.User
}

func NewTelegramController(botToken string) (*TelegramController, error) {
	newBotAPI, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return nil, err
	}

	return &TelegramController{
		botAPI: newBotAPI,
	}, nil
}

func (c *TelegramController) checkConnection() error {
	return c.fetchBotInfo()
}

func (c *TelegramController) fetchBotInfo() error {
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
			msg := TGmsg{
				ID:     int64(update.Message.MessageID),
				Text:   update.Message.Text,
				ChatID: update.Message.Chat.ID,
			}

			//responseMsg := c.processor.RouteMessage(msg)
			responseMsg := handler.ChatBotMsgProcess(msg.Text)

			response := tgbotapi.NewMessage(msg.ChatID, responseMsg)
			_, err := c.botAPI.Send(response)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
