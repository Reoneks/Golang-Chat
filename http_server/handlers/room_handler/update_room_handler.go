package room_handler

import (
	"net/http"
	"test/rooms"
	"test/user"
	"time"

	"github.com/gin-gonic/gin"
)

type UpdateRoomRequest struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Status    int64     `json:"status"`
	CreatedBy int64     `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}

func UpdateRoomHandler(roomsService room.RoomService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var updateRoomRequest UpdateRoomRequest
		if err := ctx.BindJSON(&updateRoomRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		thisUser, _ := ctx.Get("user")
		room, err := roomsService.UpdateRoom(
			room.Rooms(updateRoomRequest),
			thisUser.(*user.User).Id,
		)
		if err != nil && err.Error() == "you are not allowed to do it" {
			ctx.JSON(http.StatusForbidden, gin.H{})
		} else if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else if room == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{})
		} else {
			ctx.JSON(http.StatusNoContent, gin.H{})
		}
	}
}
