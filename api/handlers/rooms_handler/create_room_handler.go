package rooms_handler

import (
	"net/http"
	"test/rooms"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateRoomsRequest struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Status    int64     `json:"status"`
	CreatedBy int64     `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateRoomHandler(roomsService rooms.RoomService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var createRoomsRequest CreateRoomsRequest
		if err := ctx.BindJSON(&createRoomsRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		rooms, err := roomsService.CreateRoom(rooms.Rooms(createRoomsRequest))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else if rooms == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{})
		} else {
			ctx.JSON(http.StatusCreated, rooms)
		}
	}
}
