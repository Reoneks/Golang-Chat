package room_handler

import (
	"net/http"
	"test/room"
	"time"

	"github.com/gin-gonic/gin"
)

type UpdateMessageRequest struct {
	Id        int64     `json:"id"`
	Text      string    `json:"text"`
	Status    int64     `json:"status"`
	RoomID    int64     `json:"room_id"`
	CreatedBy int64     `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func UpdateMessageHandler(roomsService room.RoomService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var updateMessageRequest UpdateMessageRequest
		if err := ctx.BindJSON(&updateMessageRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		message, err := roomsService.UpdateMessage(room.Message(updateMessageRequest))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
		} else if message == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{})
		} else {
			ctx.JSON(http.StatusNoContent, gin.H{})
		}
	}
}
