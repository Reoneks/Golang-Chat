package main

import (
	"log"
	"net/http"
	"test/auth"
	. "test/connector"
	"test/http_server"
	. "test/room"
	. "test/user"

	"github.com/gorilla/websocket"

	c "test/config"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	readBufferSize  = 1024
	writeBufferSize = 1024
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	config := c.NewConfig()

	db := config.DBClient()
	jwt := config.JWT()
	log := config.Log()
	app_env := config.AppEnvironment()

	var upgrader = &websocket.Upgrader{
		ReadBufferSize:  readBufferSize,
		WriteBufferSize: writeBufferSize,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	userRepository := NewUserRepository(db)
	roomsRepository := NewRoomRepository(db)
	uPRepository := NewRoomUsersRepository(db)
	commentsRepository := NewMessagesRepository(db)

	userService := NewUserService(userRepository)
	roomsService := NewRoomService(roomsRepository, uPRepository, commentsRepository)
	authService := auth.NewAuthService(userService, jwt)

	connector := NewWSConnector(log, roomsService)

	httpServer := http_server.NewHTTPServer(
		config.ServerAddress(),
		authService,
		userService,
		roomsService,
		connector,
		upgrader,
		jwt,
		log,
		app_env,
	)

	log.Printf("HTTP Server listening at: %v", config.ServerAddress().String())

	if err := httpServer.Start(); err != nil {
		panic(err)
	}
}
