package room_handler

import (
	"net/http"
	. "test/room"
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

func AddMessageHandler(roomsService RoomService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var addMessageRequest AddMessageRequest
		if err := ctx.BindJSON(&addMessageRequest); err != nil {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusBadRequest,
				Meta: "add_message_handler 24: Bind failed",
			})
			return
		}

		Message, err := roomsService.AddMessage(Message{
			Id:        addMessageRequest.Id,
			Text:      addMessageRequest.Text,
			Status:    NewStatusType(addMessageRequest.Status),
			RoomID:    addMessageRequest.RoomID,
			CreatedBy: addMessageRequest.CreatedBy,
			CreatedAt: addMessageRequest.CreatedAt,
			UpdatedAt: addMessageRequest.UpdatedAt,
		})
		if err != nil {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusInternalServerError,
				Meta: "add_message_handler 34: Add message error",
			})
		} else if Message == nil {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusBadRequest,
				Meta: "add_message_handler 40: Add message failed",
			})
		} else {
			ctx.JSON(http.StatusCreated, Message)
		}
	}
}
