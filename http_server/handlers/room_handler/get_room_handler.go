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
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusBadRequest,
				Meta: "get_room_handler 14: Bad parameter roomId",
			})
		}

		room, messages, err := roomsService.GetRoom(int64(id))
		if err != nil {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusInternalServerError,
				Meta: "get_room_handler 23: Get room error",
			})
		} else if room == nil {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusNotFound,
				Meta: "get_room_handler 29: Get room failed",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"room":     room,
				"messages": messages,
			})
		}
	}
}
