package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

/*
 *ctx.Errors = append(ctx.Errors, &gin.Error{
 *	Err:  errors.New("some error"),
 *	Type: http.StatusBadRequest,
 *	Meta: "Some message",
 *})
 */
func ErrorHandling(log *logrus.Entry) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		err := ctx.Errors.Last()
		if err == nil {
			return
		}
		errInfo := err.Meta.(string)

		log.WithFields(logrus.Fields{
			"time":    time.Now(),
			"message": errInfo,
		}).WithError(err.Err)

		ctx.AbortWithStatusJSON(int(err.Type), gin.H{
			"Message:": errInfo,
		})
	}
}
