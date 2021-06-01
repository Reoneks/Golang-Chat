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
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		validate := validator.New()
		if err := validate.Struct(&loginRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		login, err := authService.Login(loginRequest.Email, loginRequest.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else if login == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "can't login",
			})
		} else {
			ctx.JSON(http.StatusOK, login)
		}
	}
}
