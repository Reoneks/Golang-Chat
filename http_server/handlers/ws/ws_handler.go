package ws

import (
	"net/http"
	. "test/connector"
	. "test/user"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func WSHandler(connector Connector, upgrader *websocket.Upgrader) func(c *gin.Context) {
	return func(ctx *gin.Context) {
		conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		userCtx, _ := ctx.Get("user")
		connection := NewWSConnection(ctx.Request, conn, userCtx.(*User))
		connector.AddConnection(connection)
	}
}
