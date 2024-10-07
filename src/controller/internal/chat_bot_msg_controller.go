package internal

import "upserv/src/service"

type ChatBotMsgController struct {
	chatBotMsgService service.IChatBotMsg
}

func (e *ChatBotMsgController) Handle(msg string) string {
	return e.chatBotMsgService.Handle(msg)
}

func NewChatBotMsgController(chatBotMsgService service.IChatBotMsg) *ChatBotMsgController {
	return &ChatBotMsgController{
		chatBotMsgService: chatBotMsgService,
	}
}
