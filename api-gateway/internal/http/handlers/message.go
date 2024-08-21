package handlers

import (
	"api-gateway/genproto/messagepb"
	"net/http"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *HandlerST) WsHandler(ctx *gin.Context) {
	user_id := ctx.Query("user_id")
	to_id := ctx.Query("to_id")

	for {
		conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			ctx.JSON(400, gin.H{
				"message": "WebSocket aloqasini ochishda xato",
				"error":   err.Error(),
			})
			time.Sleep(5 * time.Second)
			continue
		}

		for {
			userResponseFirst, err := h.service.GetMessagesByTo(ctx, &messagepb.GetMessagesByToRequest{
				UserId: user_id,
				To:     to_id,
			})
			if err != nil {
				ctx.JSON(400, gin.H{
					"message": "Xabarlarni olishda xato",
					"error":   err.Error(),
				})
				conn.Close()
				break
			}

			userResponseTwo, err := h.service.GetMessagesByTo(ctx, &messagepb.GetMessagesByToRequest{
				UserId: to_id,
				To:     user_id,
			})
			if err != nil {
				ctx.JSON(400, gin.H{
					"message": "Xabarlarni olishda xato",
					"error":   err.Error(),
				})
				conn.Close()
				break
			}

			a := append(userResponseFirst.Messages, userResponseTwo.Messages...)
			sort.Slice(a, func(i, j int) bool {
				return a[i].CreatedAt < a[j].CreatedAt
			})

			if err := conn.WriteJSON(a); err != nil {
				ctx.JSON(400, gin.H{
					"message": "conn ga writeJson qilishda xato",
					"error":   err.Error(),
				})
				conn.Close()
				break
			}

			select {
			case <-time.After(0):
				continue
			case <-ctx.Done():
				conn.Close()
				return
			}
		}
	}
}

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
