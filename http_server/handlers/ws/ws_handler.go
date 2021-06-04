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
			ctx.Errors = append(ctx.Errors, &gin.Error{
				Err:  err,
				Type: http.StatusInternalServerError,
				Meta: "ws_handler 15: Upgrade error",
			})
			return
		}
		userCtx, _ := ctx.Get("user")
		connection := NewWSConnection(ctx.Request, conn, userCtx.(*User))
		connector.AddConnection(connection)
	}
}
