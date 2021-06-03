package room_handler

import (
	"net/http"
	"test/room"
	"time"

	"github.com/gin-gonic/gin"
)

type AddMessageRequest struct {
	Id        int64     `json:"id"`
	Text      string    `json:"text"`
	Status    int64     `json:"status"`
	RoomID    int64     `json:"room_id"`
	CreatedBy int64     `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func AddMessageHandler(roomsService room.RoomService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var addMessageRequest AddMessageRequest
		if err := ctx.BindJSON(&addMessageRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		Message, err := roomsService.AddMessage(room.Message(addMessageRequest))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
		} else if Message == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{})
		} else {
			ctx.JSON(http.StatusCreated, Message)
		}
	}
}
