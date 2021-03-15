package main

import (
	"context"
	"github.com/labstack/echo/v4"
	_messageHttpDelivery "github.com/sakkayaphab/micro-service2/message/delivery/http"
	_messageRepo "github.com/sakkayaphab/micro-service2/message/repository"
	_messageUsecase "github.com/sakkayaphab/micro-service2/message/usecase"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {

	topic := "message"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	err = conn.SetWriteDeadline(time.Time{})
	if err!=nil {
		log.Fatal(err)
	}

	// init echo
	e := echo.New()

	// repository
	messageRepo := _messageRepo.NewMessageRepository(conn)

	// usecase
	timeoutContext := 20 * time.Second
	messageUsecase := _messageUsecase.NewMessageUsecase(messageRepo, timeoutContext)

	// delivery
	_messageHttpDelivery.NewMessageHandler(e,messageUsecase)

	// run server
	go func() {
		if err := e.Start(":8080"); err != nil {
			log.Fatal("Server start failed :" + err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Fatal("Shutting down the server error : " + err.Error())
	} else {
		if err := conn.Close(); err != nil {
			log.Fatal("failed to close writer:", err)
		}
		log.Fatal("Shutting down the server")
	}

}
