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
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		thisUser, _ := ctx.Get("user")
		err = roomsService.DeleteRoom(int64(id), thisUser.(*user.User).Id)
		if err != nil && err.Error() == "you are not allowed to do it" {
			ctx.JSON(http.StatusForbidden, gin.H{})
		} else if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusNoContent, gin.H{})
		}
	}
}
