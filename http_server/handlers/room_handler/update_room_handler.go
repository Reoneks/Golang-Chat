package room_handler

import (
	"net/http"
	"test/room"
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
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusBadRequest,
				Meta: "update_room_handler 23: Bind failed",
			})
			return
		}

		thisUser, _ := ctx.Get("user")
		room, err := roomsService.UpdateRoom(
			room.Rooms(updateRoomRequest),
			thisUser.(*user.User).Id,
		)
		if err != nil && err.Error() == "you are not allowed to do it" {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusForbidden,
				Meta: "update_room_handler 37: This user is not allowed to do this",
			})
		} else if err != nil {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusInternalServerError,
				Meta: "update_room_handler 43: Update room error",
			})
		} else if room == nil {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusBadRequest,
				Meta: "update_room_handler 49: Update room failed",
			})
		} else {
			ctx.JSON(http.StatusNoContent, gin.H{})
		}
	}
}
