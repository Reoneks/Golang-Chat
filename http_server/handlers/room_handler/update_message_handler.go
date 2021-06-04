package room_handler

import (
	"net/http"
	. "test/room"
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

func UpdateMessageHandler(roomsService RoomService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var updateMessageRequest UpdateMessageRequest
		if err := ctx.BindJSON(&updateMessageRequest); err != nil {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusBadRequest,
				Meta: "update_message_handler 24: Bind failed",
			})
			return
		}

		message, err := roomsService.UpdateMessage(Message{
			Id:        updateMessageRequest.Id,
			Text:      updateMessageRequest.Text,
			Status:    NewStatusType(updateMessageRequest.Status),
			RoomID:    updateMessageRequest.RoomID,
			CreatedBy: updateMessageRequest.CreatedBy,
			CreatedAt: updateMessageRequest.CreatedAt,
			UpdatedAt: updateMessageRequest.UpdatedAt,
		})
		if err != nil {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusInternalServerError,
				Meta: "update_message_handler 34: Update message error",
			})
		} else if message == nil {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusBadRequest,
				Meta: "update_message_handler 40: Update message failed",
			})
		} else {
			ctx.JSON(http.StatusNoContent, gin.H{})
		}
	}
}
