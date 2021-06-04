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
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusBadRequest,
				Meta: "get_user_handler 14: Bad parameter userId",
			})
			return
		}

		obtainedUser, err := userService.GetUser(int64(id))
		if err != nil {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusInternalServerError,
				Meta: "get_user_handler 24: Get user error",
			})
		} else if obtainedUser == nil {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusNotFound,
				Meta: "get_user_handler 30: Get user failed",
			})
		} else {
			ctx.JSON(http.StatusOK, obtainedUser)
		}
	}
}
