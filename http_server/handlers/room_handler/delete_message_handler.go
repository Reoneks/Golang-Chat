package room_handler

import (
	"net/http"
	"strconv"
	"test/room"

	"github.com/gin-gonic/gin"
)

func DeleteMessageHandler(roomsService room.RoomService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("messageId"))
		if err != nil {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusBadRequest,
				Meta: "delete_message_handler 14: Bad parameter messageId",
			})
		}

		err = roomsService.DeleteMessage(int64(id))
		if err != nil {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusInternalServerError,
				Meta: "delete_message_handler 23: Delete message error",
			})
		} else {
			ctx.JSON(http.StatusNoContent, gin.H{})
		}
	}
}
