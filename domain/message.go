package domain

import "context"

type Message struct {
	MsgId int64 `json:"Msg_id"`
	Sender string `json:"Sender"`
	Msg string `json:"Msg"`
}

type MessageUsecase interface {
	Store(c context.Context,m *Message) error
}

type MessageRepositorty interface {
	Store(c context.Context,m *Message) error
}