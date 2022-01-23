package models

type Message struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
}

type UserMessage struct {
	UserId    string `json:"userId"`
	MessageId int    `json:"messageId"`
	Message   string `json:"message"`
}

type Messages struct {
	List []Message `json:"messages"`
}

type UserMessages struct {
	List []UserMessage `json:"userMessages"`
}
