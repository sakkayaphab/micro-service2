package http

import (
	"github.com/labstack/echo/v4"
	"github.com/sakkayaphab/micro-service2/domain"
	"net/http"
	"time"
)

type messageHandler struct {
	messageUsecase domain.MessageUsecase
}

func NewMessageHandler(e *echo.Echo, us domain.MessageUsecase) {
	handler := &messageHandler{
		messageUsecase: us,
	}

	v1Route := e.Group("/v1")
	v1Route.POST("/message", handler.store)
}

func (h *messageHandler) store(c echo.Context) (err error) {
	ctx := c.Request().Context()

	//claims := c.Get("claims").(*domain.AdminAuthClaims)
	reqBody := domain.Message{}
	if err := c.Bind(&reqBody); err != nil {
		return err
	}

	err = h.messageUsecase.Store(ctx,&reqBody)
	if err!=nil {
		return c.JSON(http.StatusInternalServerError,Response{
			Code: "ERROR",
			ReceivedTime: time.Now().Format("2006-01-02T15:04:05.000Z"),
		})
	}

	return c.JSON(http.StatusCreated,Response{
		Code: "OK",
		ReceivedTime: time.Now().Format("2006-01-02T15:04:05.000Z"),
	})
}

type Response struct {
	Code string `json:"Code"`
	ReceivedTime string `json:"Received_Time"`
}