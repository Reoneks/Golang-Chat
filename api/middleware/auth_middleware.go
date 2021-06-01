package middleware

import (
	"net/http"
	"strings"
	"test/user"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/jwtauth"
)

func AuthMiddleware(userService user.UserService, jwt *jwtauth.JWTAuth) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if len(ctx.Request.Header["Authorization"]) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Not enough data (There is no token)",
			})
			return
		} else if len(strings.Split(ctx.Request.Header["Authorization"][0], " ")) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token issued wrong",
			})
			return
		}
		reqToken := strings.Split(ctx.Request.Header["Authorization"][0], " ")[1]
		token, err := jwtauth.VerifyToken(jwt, reqToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		userId, ok := token.Get("user_id")
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "can't get user_id",
			})
			return
		}
		obtainedUser, err := userService.GetUser(int64(userId.(float64)))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		} else if obtainedUser == nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "can't find user",
			})
			return
		}
		if obtainedUser.Status != user.Active {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "can't login",
			})
			return
		}
		ctx.Set("user", obtainedUser)
		ctx.Next()
	}
}
