package room_handler

import (
	"net/http"
	"strconv"
	"test/room"
	"test/user"

	"github.com/gin-gonic/gin"
)

func DeleteRoomHandler(roomsService room.RoomService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("roomId"))
		if err != nil {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusBadRequest,
				Meta: "delete_room_handler 15: Bad parameter roomId",
			})
		}

		thisUser, _ := ctx.Get("user")
		err = roomsService.DeleteRoom(int64(id), thisUser.(*user.User).Id)
		if err != nil && err.Error() == "you are not allowed to do it" {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusForbidden,
				Meta: "delete_room_handler 25: This user is not allowed to do this",
			})
		} else if err != nil {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusInternalServerError,
				Meta: "delete_room_handler 31: Delete room error",
			})
		} else {
			ctx.JSON(http.StatusNoContent, gin.H{})
		}
	}
}
