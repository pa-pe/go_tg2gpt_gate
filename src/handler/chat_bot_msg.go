package handler

import (
	"upserv/src/controller"
)

func ChatBotMsgProcess(msg string) string {
	// codereview: controller.ControllerImp.MsgProcessor.Handle(msg) <- записать в дб и вернуть id записи
	return controller.ControllerImp.ChatBotMsgController.Handle(msg)
}
