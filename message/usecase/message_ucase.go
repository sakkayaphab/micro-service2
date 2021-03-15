package usecase

import (
	"context"
	"github.com/sakkayaphab/micro-service2/domain"
	"time"
)

type messageUsecase struct {
	messageRepo domain.MessageRepositorty
	contextTimeout time.Duration
}

func NewMessageUsecase(messageRepo domain.MessageRepositorty,timeout time.Duration) domain.MessageUsecase {
	return &messageUsecase{
		messageRepo: messageRepo,
		contextTimeout: timeout,
	}
}

func (u *messageUsecase) Store(c context.Context,register *domain.Message) (error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	return u.messageRepo.Store(ctx,register)
}