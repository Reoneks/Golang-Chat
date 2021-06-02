package room_handler

import (
	"net/http"
	"test/rooms"

	"github.com/gin-gonic/gin"
)

type GetRoomsRequest struct {
	Name *string `form:"name"`
}

func GetRoomsHandler(roomsService room.RoomService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var getRoomsRequest GetRoomsRequest
		if err := ctx.Bind(&getRoomsRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		filter := &room.RoomsFilter{
			Name: getRoomsRequest.Name,
		}
		rooms, err := roomsService.GetRooms(filter)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else if rooms == nil {
			ctx.JSON(http.StatusNotFound, gin.H{})
		} else {
			ctx.JSON(http.StatusOK, rooms)
		}
	}
}
