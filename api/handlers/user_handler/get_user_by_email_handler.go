package user_handler

import (
	"net/http"
	"test/user"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type GetUserByEmailRequest struct {
	Email string `form:"email" json:"email" validate:"email"`
}

func GetUserByEmailHandler(userService user.UserService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var getUserByEmailRequest GetUserByEmailRequest
		if err := ctx.Bind(&getUserByEmailRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		validate := validator.New()
		if err := validate.Struct(&getUserByEmailRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		user, err := userService.GetUserByEmail(getUserByEmailRequest.Email)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else if user == nil {
			ctx.JSON(http.StatusNotFound, gin.H{})
		} else {
			ctx.JSON(http.StatusOK, user)
		}
	}
}
