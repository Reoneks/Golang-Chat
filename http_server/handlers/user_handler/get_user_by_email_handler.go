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
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusBadRequest,
				Meta: "get_user_by_email_handler 18: Bind failed",
			})
			return
		}

		validate := validator.New()
		if err := validate.Struct(&getUserByEmailRequest); err != nil {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusBadRequest,
				Meta: "get_user_by_email_handler 28: Validation failed",
			})
			return
		}

		user, err := userService.GetUserByEmail(getUserByEmailRequest.Email)
		if err != nil {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusInternalServerError,
				Meta: "get_user_by_email_handler 38: Get user by email error",
			})
		} else if user == nil {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusNotFound,
				Meta: "get_user_by_email_handler 44: Get user by email failed",
			})
		} else {
			ctx.JSON(http.StatusOK, user)
		}
	}
}
