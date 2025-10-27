package message

type MessageService struct {
	repo MessageRepository
}

func NewMessageService() MessageService {
	return MessageService {
		repo: NewMessageRepository(),
	}
}

func (s MessageService) GetService() []Message{
	return s.repo.Get()
}
