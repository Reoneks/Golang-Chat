package user_handler

import (
	"net/http"
	"test/user"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RegistrationRequest struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" validate:"email"`
	Password  string `json:"password" validate:"min=8,max=32,alphanum"`
	Status    int64  `json:"status"`
}

func RegistrationHandler(userService user.UserService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var registrationRequest RegistrationRequest
		if err := ctx.BindJSON(&registrationRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		validate := validator.New()
		if err := validate.Struct(&registrationRequest); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})
			return
		}

		userModel := user.User{
			Id:        registrationRequest.Id,
			FirstName: registrationRequest.FirstName,
			LastName:  registrationRequest.LastName,
			Email:     registrationRequest.Email,
			Password:  registrationRequest.Password,
			Status:    user.NewStatusCode(registrationRequest.Status),
		}
		user, err := userService.Registration(userModel)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else if user == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "can't register the user",
			})
		} else {
			ctx.JSON(http.StatusCreated, user)
		}
	}
}
