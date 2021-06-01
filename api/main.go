package api

import (
	"net/http"
	"net/url"

	"test/api/handlers/auth_handler"
	"test/api/handlers/rooms_handler"
	"test/api/handlers/user_handler"
	"test/api/middleware"
	"test/auth"
	"test/rooms"
	"test/user"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/jwtauth"
	"github.com/sirupsen/logrus"
)

type HTTPServer interface {
	Start() error
}

type httpServer struct {
	url          *url.URL
	authService  auth.AuthService
	userService  user.UserService
	roomsService rooms.RoomService
	jwt          *jwtauth.JWTAuth
	log          *logrus.Entry
}

func NewHTTPServer(
	url *url.URL,
	authService auth.AuthService,
	userService user.UserService,
	roomsService rooms.RoomService,
	jwt *jwtauth.JWTAuth,
	log *logrus.Entry,
) HTTPServer {
	return &httpServer{
		url,
		authService,
		userService,
		roomsService,
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

	private := router.Group("/")

	private.Use(middleware.AuthMiddleware(s.userService, s.jwt))
	{
		//^ User Handlers
		private.GET("/user/:userId", user_handler.GetUserHandler(s.userService))
		private.GET("/user/email", user_handler.GetUserByEmailHandler(s.userService))
		private.GET("/users", user_handler.GetUsersHandler(s.userService))
		private.DELETE("/user", user_handler.DeleteUserHandler(s.userService))

		//^ Rooms Handlers
		private.GET("/room/:roomId", rooms_handler.GetRoomHandler(s.roomsService))
		private.GET("/rooms", rooms_handler.GetRoomsHandler(s.roomsService))
		private.POST("/room", rooms_handler.CreateRoomHandler(s.roomsService))
		private.DELETE("/room/:roomId", rooms_handler.DeleteRoomHandler(s.roomsService))
		private.PUT("/room", rooms_handler.UpdateRoomHandler(s.roomsService))
		private.POST("/users", rooms_handler.AddUsersHandler(s.roomsService))
		private.POST("/message", rooms_handler.AddMessageHandler(s.roomsService))
		private.PUT("/message", rooms_handler.UpdateMessageHandler(s.roomsService))
		private.DELETE("/message/:messageId", rooms_handler.DeleteMessageHandler(s.roomsService))
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
