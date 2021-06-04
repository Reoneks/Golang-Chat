package room_handler

import (
	"net/http"
	"test/room"
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

func CreateRoomHandler(roomsService room.RoomService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var createRoomsRequest CreateRoomsRequest
		if err := ctx.BindJSON(&createRoomsRequest); err != nil {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusBadRequest,
				Meta: "create_room_handler 22: Bind failed",
			})
			return
		}

		rooms, err := roomsService.CreateRoom(room.Rooms(createRoomsRequest))
		if err != nil {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusInternalServerError,
				Meta: "create_room_handler 32: Create room error",
			})
		} else if rooms == nil {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusBadRequest,
				Meta: "create_room_handler 38: Create room failed",
			})
		} else {
			ctx.JSON(http.StatusCreated, rooms)
		}
	}
}
