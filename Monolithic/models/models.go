package models

type UserMessage struct {
	UserId    string `json:"userId"`
	MessageId int    `json:"messageId"`
	Message   string `json:"message"`
}

type UserMessages struct {
	List []UserMessage `json:"userMessages"`
}
