package user_handler

import (
	"net/http"
	"strconv"
	"strings"
	"test/user"

	"github.com/gin-gonic/gin"
)

type GetUsersRequest struct {
	FirstName *string `form:"first_name"`
	LastName  *string `form:"last_name"`
	Email     *string `form:"email"`
	Status    *string `form:"status"`
}

func GetUsersHandler(userService user.UserService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var getUsersRequest GetUsersRequest
		if err := ctx.Bind(&getUsersRequest); err != nil {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusBadRequest,
				Meta: "get_users_handler 22: Bind failed",
			})
			return
		}

		filter := &user.UserFilter{
			FirstName: getUsersRequest.FirstName,
			LastName:  getUsersRequest.LastName,
			Email:     getUsersRequest.Email,
		}
		if getUsersRequest.Status != nil && len(*getUsersRequest.Status) > 0 {
			statuses := strings.Split(*getUsersRequest.Status, ",")
			var statusesInt64 []int64
			for _, a := range statuses {
				i, err := strconv.Atoi(a)
				if err != nil {
					ctx.Errors = append(ctx.Errors, &gin.Error{
						Err:  err,
						Type: http.StatusBadRequest,
						Meta: "get_users_handler 41: Status convert failed",
					})
					return
				}
				statusesInt64 = append(statusesInt64, int64(i))
			}
			filter.Status = statusesInt64
		}
		user, err := userService.GetUsers(filter)
		if err != nil {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusInternalServerError,
				Meta: "get_users_handler 54: Get users error",
			})
		} else if user == nil {
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusNotFound,
				Meta: "get_users_handler 60: Get users failed",
			})
		} else {
			ctx.JSON(http.StatusOK, user)
		}
	}
}
