package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sakkayaphab/micro-service2/domain"
	"github.com/segmentio/kafka-go"
)

type messageRepository struct {
	conn *kafka.Writer
}

func NewMessageRepository(conn *kafka.Writer) domain.MessageRepositorty {
	return &messageRepository{
		conn: conn,
	}
}

func (r *messageRepository) Store(c context.Context,message *domain.Message) error {

	jsonData, err := json.MarshalIndent(message, "", "    ")
	if err!=nil {
		return err
	}

	patition := r.conn.Balancer.Balance(kafka.Message{Value: jsonData},1,2,3)
	fmt.Println(patition)
	err = r.conn.WriteMessages(context.Background(),
		kafka.Message{Value: jsonData,Partition: patition},
	)

	return err
}