package handlers

import (
	"api-gateway/genproto/messagepb"

	"github.com/gin-gonic/gin"
)

func (h *HandlerST) CreateMessage(ctx *gin.Context) {
	var message messagepb.CreateMessageRequest
	err := ctx.BindJSON(&message)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	resp, err := h.service.CreateMessage(ctx, &message)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}

func (h *HandlerST) UpdateMessage(ctx *gin.Context) {
	user_id := ctx.Param("id")
	var message messagepb.UpdateMessageRequest
	err := ctx.BindJSON(&message)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	message.UserId = user_id

	resp, err := h.service.UpdateMessage(ctx, &message)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}

func (h *HandlerST) DeleteMessage(ctx *gin.Context) {
	user_id := ctx.Query("user_id")
	message_id := ctx.Query("message_id")
	message := messagepb.DeleteMessageRequest{
		UserId:    user_id,
		MessageId: message_id,
	}

	resp, err := h.service.DeleteMessage(ctx, &message)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}

func (h *HandlerST) GetMessagesByTo(ctx *gin.Context) {
	user_id := ctx.Query("user_id")
	to_id := ctx.Query("to_id")
	message := messagepb.GetMessagesByToRequest{
		UserId: user_id,
		To:     to_id,
	}
	resp, err := h.service.GetMessagesByTo(ctx, &message)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}
