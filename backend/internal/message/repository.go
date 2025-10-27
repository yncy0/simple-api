package message

type MessageRepository struct {
	message []Message
}

func NewMessageRepository() MessageRepository {
	return MessageRepository{
		message: []Message{
			{ID: "0", Message: "Hello"},
			{ID: "1", Message: "World!"},
		},
	}
}

func (r *MessageRepository) Get() []Message {
	return r.message
}
