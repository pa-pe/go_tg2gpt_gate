package telegram

import (
	"fmt"
	"upserv/logger"
	"upserv/src/service/telegram/controller"
	"upserv/src/service/telegram/service"
)

type BotImpl struct {
	token string
}

func NewTelegramBotService(token string) *BotImpl {
	return &BotImpl{token: token}
}

func (bot *BotImpl) Start() {
	processor := usecases.NewMessageProcessor()

	tgController, err := controller.NewTelegramController(bot.token, processor)
	if err != nil {
		//		log.Fatalf("Error creating Telegram tgController: %v", err)
		errMsg := fmt.Sprintf("Error creating Telegram tgController: %v", err)
		logger.LaunchLog(errMsg)
	}

	// Start bot in goroutine
	go func() {
		if err := tgController.CheckConnection(); err != nil {
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
	}()
}
