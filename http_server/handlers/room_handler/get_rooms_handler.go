package room_handler

import (
	"net/http"
	"test/room"

	"github.com/gin-gonic/gin"
)

type GetRoomsRequest struct {
	Name *string `form:"name"`
}

func GetRoomsHandler(roomsService room.RoomService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var getRoomsRequest GetRoomsRequest
		if err := ctx.Bind(&getRoomsRequest); err != nil {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusBadRequest,
				Meta: "get_rooms_handler 17: Bind failed",
			})
			return
		}

		filter := &room.RoomsFilter{
			Name: getRoomsRequest.Name,
		}
		rooms, err := roomsService.GetRooms(filter)
		if err != nil {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusInternalServerError,
				Meta: "get_rooms_handler 30: Get rooms error",
			})
		} else if rooms == nil {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusNotFound,
				Meta: "get_rooms_handler 36: Get rooms failed",
			})
		} else {
			ctx.JSON(http.StatusOK, rooms)
		}
	}
}
