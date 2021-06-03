package ws

import (
	"net/http"
	"strconv"
	. "test/connector"
	. "test/user"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var id = int64(0)

func WSHandler(connector Connector, upgrader *websocket.Upgrader) func(c *gin.Context) {
	return func(ctx *gin.Context) {
		conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		//userCtx, _ := ctx.Get("user")
		id++
		user := &User{
			Id:        id,
			FirstName: "SomeName" + strconv.FormatInt(id, 10),
			LastName:  "SomeName",
			Email:     "some" + strconv.FormatInt(id, 10) + "@example.com",
			Password:  "wefwef",
			Status:    1,
		}
		connection := NewWSConnection(ctx.Request, conn /*userCtx.(*User)*/, user)
		connector.AddConnection(connection)
	}
}
