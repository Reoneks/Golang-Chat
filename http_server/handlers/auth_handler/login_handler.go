package auth_handler

import (
	"net/http"
	"test/auth"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"min=4,max=32,alphanum"`
}

func LoginHandler(authService auth.AuthService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var loginRequest LoginRequest
		if err := ctx.BindJSON(&loginRequest); err != nil {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusBadRequest,
				Meta: "login_handler 19: Bind failed",
			})
			return
		}

		validate := validator.New()
		if err := validate.Struct(&loginRequest); err != nil {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusBadRequest,
				Meta: "login_handler 29: Validation failed",
			})
			return
		}

		login, err := authService.Login(loginRequest.Email, loginRequest.Password)
		if err != nil {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusInternalServerError,
				Meta: "login_handler 39: Login error",
			})
		} else if login == nil {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusNotFound,
				Meta: "login_handler 45: Login failed",
			})
		} else {
			ctx.JSON(http.StatusOK, login)
		}
	}
}
