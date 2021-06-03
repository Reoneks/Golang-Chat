package room_handler

import (
	"net/http"
	"strconv"
	"test/room"

	"github.com/gin-gonic/gin"
)

func GetRoomHandler(roomsService room.RoomService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("roomId"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		room, messages, err := roomsService.GetRoom(int64(id))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else if room == nil {
			ctx.JSON(http.StatusNotFound, gin.H{})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"room":     room,
				"messages": messages,
			})
		}
	}
}
