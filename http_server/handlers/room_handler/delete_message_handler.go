package room_handler

import (
	"net/http"
	"strconv"
	"test/rooms"

	"github.com/gin-gonic/gin"
)

func DeleteMessageHandler(roomsService room.RoomService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("messageId"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		err = roomsService.DeleteMessage(int64(id))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusNoContent, gin.H{})
		}
	}
}
