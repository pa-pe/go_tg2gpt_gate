package controller

import (
	"reflect"
	"upserv/src/controller/internal"
	"upserv/src/service"
)

var ControllerImp controllerImp

type controllerImp struct {
	HelloWorldController *internal.HelloWorldController
	ChatBotMsgController *internal.ChatBotMsgController
}

func InitControllers(services *service.Services) {
	ControllerImp.HelloWorldController = internal.NewHelloWorldController(services.HelloWorld)
	ControllerImp.ChatBotMsgController = internal.NewChatBotMsgController(services.ChatBotMsg)
}

func IsValid() bool {
	reflectC := reflect.ValueOf(ControllerImp)
	for i := 0; i < reflectC.Type().NumField(); i++ {
		if reflect.Indirect(reflectC).Field(i).IsNil() {
			return false
		}
	}
	return true
}
