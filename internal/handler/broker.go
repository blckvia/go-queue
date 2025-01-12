package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/rand"

	"gtihub.com/blckvia/go-queue/internal/broker"
)

// @Summary Send message to queue
// @Tags Queues
// @Description Send a message to a specific queue
// @ID send-message
// @Accept  json
// @Produce  json
// @Param queue_name path string true "Queue Name"
// @Param input body models.Message true "Message to send"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /v1/queues/{queue_name}/messages [post]
func (h *Handler) sendMessage(c *gin.Context) {
	queueName := c.Param("queue_name")
	var msg broker.Message
	if err := c.ShouldBindBodyWithJSON(&msg); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	rand.Seed(uint64(time.Now().UnixNano()))

	err := h.broker.SendMessage(queueName, &broker.Message{
		ID:      rand.Int(),
		Message: msg.Message,
	})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// @Summary Subscribe to queue
// @Tags Queues
// @Description Subscribe to a specific queue
// @ID subscribe-queue
// @Accept  json
// @Produce  json
// @Param queue_name path string true "Queue Name"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /v1/queues/{queue_name}/subscriptions [post]
func (h *Handler) subscribe(c *gin.Context) {
	queueName := c.Param("queue_name")

	subscriber := make(chan *broker.Message, 1)

	err := h.broker.AddSubscriber(queueName, subscriber)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"subscribed"})
}
