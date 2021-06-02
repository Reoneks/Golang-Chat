package room_handler

import (
	"net/http"
	"test/rooms"
	"test/user"

	"github.com/gin-gonic/gin"
)

type AddUsersRequest struct {
	RoomId     int64   `json:"room_id"`
	UsersToAdd []int64 `json:"other_users"`
}

func AddUsersHandler(roomsService room.RoomService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var addUsersRequest AddUsersRequest
		if err := ctx.BindJSON(&addUsersRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		thisUser, _ := ctx.Get("user")
		err := roomsService.AddUsers(
			addUsersRequest.RoomId,
			thisUser.(*user.User).Id,
			addUsersRequest.UsersToAdd,
		)
		if err != nil && err.Error() == "you are not allowed to do it" {
			ctx.JSON(http.StatusForbidden, gin.H{})
		} else if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
		} else {
			ctx.JSON(http.StatusCreated, gin.H{})
		}
	}
}
