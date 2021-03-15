package repository

import (
	"context"
	"encoding/json"
	"github.com/sakkayaphab/micro-service2/domain"
	"github.com/segmentio/kafka-go"
)

type messageRepository struct {
	conn *kafka.Conn
}

func NewMessageRepository(conn *kafka.Conn) domain.MessageRepositorty {
	return &messageRepository{
		conn: conn,
	}
}

func (r *messageRepository) Store(c context.Context,message *domain.Message) error {

	jsonData, err := json.MarshalIndent(message, "", "    ")
	if err!=nil {
		return err
	}

	_, err = r.conn.WriteMessages(
		kafka.Message{Value: jsonData},
	)

	return err
}