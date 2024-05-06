package ws

type Chatrooms struct {
	chatrooms map[string]ConnectedClients
}

func NewChatrooms() Chatrooms {
	return Chatrooms{
		chatrooms: make(map[string]ConnectedClients),
	}
}
