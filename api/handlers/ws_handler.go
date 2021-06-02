package handlers

import (
	"net/http"
	"test/user"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]*user.User)
var broadcast = make(chan string)

func WSHandler(upgrader *websocket.Upgrader) func(c *gin.Context) {
	return func(ctx *gin.Context) {
		ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		defer ws.Close()
		thisUser, _ := ctx.Get("user")
		clients[ws] = thisUser.(*user.User)
		for {
			var msg string
			err := ws.ReadJSON(&msg)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": err,
				})
				delete(clients, ws)
				break
			}
			broadcast <- msg
		}
		/*connection := apiws.NewConnection(ctx.Request, conn, thisUser.(*user.User))
		connector.AddConnection(connection)*/
	}
}
