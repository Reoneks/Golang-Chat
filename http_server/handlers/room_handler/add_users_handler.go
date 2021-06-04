package room_handler

import (
	"net/http"
	"test/room"
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
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusBadRequest,
				Meta: "add_users_handler 19: Bind failed",
			})
			return
		}

		thisUser, _ := ctx.Get("user")
		err := roomsService.AddUsers(
			addUsersRequest.RoomId,
			thisUser.(*user.User).Id,
			addUsersRequest.UsersToAdd,
		)
		if err != nil && err.Error() == "you are not allowed to do it" {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusForbidden,
				Meta: "add_users_handler 34: This user is not allowed to do this",
			})
		} else if err != nil {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusInternalServerError,
				Meta: "add_users_handler 40: Add users error",
			})
		} else {
			ctx.JSON(http.StatusCreated, gin.H{})
		}
	}
}
