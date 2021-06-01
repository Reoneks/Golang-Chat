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
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusNoContent, gin.H{})
		}
	}
}
