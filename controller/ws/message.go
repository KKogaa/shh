package ws

type Message struct {
	Msg      string `json:"msg"`
	User     string `json:"user"`
	Chatroom string `json:"chatroom"`
}
