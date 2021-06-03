package http_server

import (
	"github.com/gorilla/websocket"
	"net/http"
	"net/url"
	. "test/connector"
	"test/http_server/handlers/ws"

	. "test/auth"
	"test/http_server/handlers/auth_handler"
	"test/http_server/handlers/room_handler"
	"test/http_server/handlers/user_handler"
	"test/http_server/middleware"
	. "test/room"
	. "test/user"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/jwtauth"
	"github.com/sirupsen/logrus"
)

type HTTPServer interface {
	Start() error
}

type httpServer struct {
	url         *url.URL
	authService AuthService
	userService UserService
	roomService RoomService
	connector   Connector
	upgrader    *websocket.Upgrader
	jwt         *jwtauth.JWTAuth
	log         *logrus.Entry
}

func NewHTTPServer(
	url *url.URL,
	authService AuthService,
	userService UserService,
	roomsService RoomService,
	connector Connector,
	upgrader *websocket.Upgrader,
	jwt *jwtauth.JWTAuth,
	log *logrus.Entry,
) HTTPServer {
	return &httpServer{
		url,
		authService,
		userService,
		roomsService,
		connector,
		upgrader,
		jwt,
		log,
	}
}

func (s *httpServer) Start() error {
	router := gin.Default()
	router.Use(middleware.CorsMiddleware())

	//^ Auth Handlers
	router.POST("/login", auth_handler.LoginHandler(s.authService))
	router.POST("/registration", user_handler.RegistrationHandler(s.userService))

	router.LoadHTMLGlob("site/index.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	private := router.Group("/")

	private.Use(middleware.AuthMiddleware(s.userService, s.jwt))
	{
		//^ WS Handlers
		private.Any("/ws", ws.WSHandler(s.connector, s.upgrader))

		//^ User Handlers
		private.GET("/user/:userId", user_handler.GetUserHandler(s.userService))
		private.GET("/user/email", user_handler.GetUserByEmailHandler(s.userService))
		private.GET("/users", user_handler.GetUsersHandler(s.userService))
		private.DELETE("/user", user_handler.DeleteUserHandler(s.userService))

		//^ Rooms Handlers
		private.GET("/room/:roomId", room_handler.GetRoomHandler(s.roomService))
		private.GET("/rooms", room_handler.GetRoomsHandler(s.roomService))
		private.POST("/room", room_handler.CreateRoomHandler(s.roomService))
		private.DELETE("/room/:roomId", room_handler.DeleteRoomHandler(s.roomService))
		private.PUT("/room", room_handler.UpdateRoomHandler(s.roomService))
		private.POST("/users", room_handler.AddUsersHandler(s.roomService))
		private.POST("/message", room_handler.AddMessageHandler(s.roomService))
		private.PUT("/message", room_handler.UpdateMessageHandler(s.roomService))
		private.DELETE("/message/:messageId", room_handler.DeleteMessageHandler(s.roomService))
	}

	server := &http.Server{
		Addr:           s.url.Host,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return server.ListenAndServe()
}
