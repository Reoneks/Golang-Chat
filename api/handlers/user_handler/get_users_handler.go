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
			ctx.JSON(http.StatusBadRequest, err)
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
					ctx.JSON(http.StatusBadRequest, gin.H{})
					return
				}
				statusesInt64 = append(statusesInt64, int64(i))
			}
			filter.Status = statusesInt64
		}
		user, err := userService.GetUsers(filter)
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
