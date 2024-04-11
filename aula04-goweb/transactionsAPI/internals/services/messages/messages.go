package messages_service

type IMessagesService interface {
	BuildMessage(name string) Message
}

type Message struct {
	Message string
}

type MessagesService struct {
}

func NewMessagesService() IMessagesService {
	return &MessagesService{}
}

func (s *MessagesService) BuildMessage(name string) Message {
	return Message{
		Message: name,
	}
}
