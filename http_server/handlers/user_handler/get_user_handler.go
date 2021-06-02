package user_handler

import (
	"net/http"
	"strconv"
	"test/user"

	"github.com/gin-gonic/gin"
)

func GetUserHandler(userService user.UserService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("userId"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		obtainedUser, err := userService.GetUser(int64(id))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else if obtainedUser == nil {
			ctx.JSON(http.StatusNotFound, gin.H{})
		} else {
			ctx.JSON(http.StatusOK, obtainedUser)
		}
	}
}
