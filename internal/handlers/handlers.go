package handlers

import (
	"L0/internal"
	"L0/static"
	"context"
	"github.com/labstack/echo/v4"
	json "github.com/mailru/easyjson"
	"github.com/nats-io/stan.go"
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"
)

const (
	durableID = "order-service-durable"
)

type Handler struct {
	logger *zap.SugaredLogger
	u      internal.Usecase
}

func RegisterHandlers(SC stan.Conn, u internal.Usecase, echo *echo.Echo, logger *zap.SugaredLogger) (*Handler, error) {
	handler := Handler{logger: logger, u: u}

	// Subscribe with manual ack mode, and set AckWait to 60 seconds
	aw, _ := time.ParseDuration("60s")

	_, err := SC.Subscribe(os.Getenv("NATS_CHANNEL"), handler.ReceiveOrder, stan.DurableName(durableID),
		stan.MaxInflight(25),
		stan.SetManualAckMode(),
		stan.AckWait(aw),
	)
	if err != nil {
		return nil, err
	}

	echo.GET("/api/getOrder", handler.GetOrder())
	return &handler, nil
}

func (h *Handler) ReceiveOrder(msg *stan.Msg) {
	err := msg.Ack()
	if err != nil {
		h.logger.Error(err)
		return
	} // Manual ACK
	order := internal.Order{}
	// Unmarshal JSON that represents the Order data
	err = json.Unmarshal(msg.Data, &order)
	if err != nil {
		h.logger.Error(err)
		return
	}
	err = h.u.SaveData(context.Background(), &order)
	if err != nil {
		h.logger.Error(err)
	}
}

func (h *Handler) GetOrder() echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.QueryParam("uid")
		if uid != "" {
			order, err := h.u.GetOrder(context.Background(), uid)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err)
			}
			if order == nil {
				return c.HTML(http.StatusNotFound, static.GenerateNotFound())
			}
			return c.HTML(http.StatusOK, static.GeneratePage(order))
		}
		return c.NoContent(http.StatusBadRequest)
	}
}
