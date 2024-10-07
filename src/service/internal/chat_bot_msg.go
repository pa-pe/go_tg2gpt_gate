package internal

type chatBotMsgImpl struct{}

func (c *chatBotMsgImpl) Handle(msg string) string {
	return c.echo(msg)
}

func (c *chatBotMsgImpl) echo(msg string) string {
	return "Echo: " + msg
}

func NewChatBotMsgService() *chatBotMsgImpl {
	return &chatBotMsgImpl{}
}
