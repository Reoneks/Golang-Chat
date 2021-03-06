package middleware

import (
	"net/http"
	"strings"
	"test/user"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/jwtauth"
)

const BearerToken = "Bearer "

func getToken(r *http.Request, findTokenFns ...func(r *http.Request) string) string {
	tokenStr := ""
	for _, fn := range findTokenFns {
		tokenStr = fn(r)
		if tokenStr != "" {
			break
		}
	}
	return tokenStr
}

func tokenFromQuery(r *http.Request) string {
	return r.URL.Query().Get("jwt")
}

func tokenFromHeader(r *http.Request) string {
	bearer := r.Header.Get("Authorization")
	if strings.Contains(bearer, BearerToken) {
		bearer = bearer[len(BearerToken):]
	}
	return bearer
}

func AuthMiddleware(userService user.UserService, jwt *jwtauth.JWTAuth, app_env string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := getToken(ctx.Request, tokenFromQuery, tokenFromHeader)
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}

		payload, err := jwtauth.VerifyToken(jwt, token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}

		userId, ok := payload.Get("user_id")
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}

		obtainedUser, err := userService.GetUser(int64(userId.(float64)))
		if err != nil {
			if app_env == "development" {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": err,
				})
			} else {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{})
			}
			return
		}

		if obtainedUser == nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}

		if obtainedUser.Status != user.Active {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}

		ctx.Set("user", obtainedUser)
		ctx.Next()
	}
}
