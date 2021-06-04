package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ErrorHandling(log *logrus.Entry, app_env string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		err := ctx.Errors.Last()
		if err == nil {
			return
		}
		errInfo := err.Meta.(string)

		log.Error(errInfo, " err:", err.Err)

		if app_env == "development" {
			ctx.AbortWithStatusJSON(int(err.Type), gin.H{
				"Error:": errInfo,
			})
		} else {
			ctx.AbortWithStatusJSON(int(err.Type), gin.H{})
		}
	}
}
