package user_handler

import (
	"net/http"
	"test/user"

	"github.com/gin-gonic/gin"
)

func DeleteUserHandler(userService user.UserService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		thisUser, _ := ctx.Get("user")
		err := userService.DeleteUser(thisUser.(*user.User).Id)
		if err != nil {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusInternalServerError,
				Meta: "delete_user_handler 14: Delete user error",
			})
		} else {
			ctx.JSON(http.StatusNoContent, gin.H{})
		}
	}
}
