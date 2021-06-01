package main

import (
	"log"
	"test/api"
	"test/auth"
	"test/rooms"
	"test/user"

	c "test/config"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	config := c.NewConfig()

	db := config.DBClient()
	jwt := config.JWT()
	log := config.Log()

	userRepository := user.NewUserRepository(db)
	roomsRepository := rooms.NewRoomRepository(db)
	uPRepository := rooms.NewRoomUsersRepository(db)
	commentsRepository := rooms.NewMessagesRepository(db)

	userService := user.NewUserService(userRepository)
	roomsService := rooms.NewRoomService(roomsRepository, uPRepository, commentsRepository)
	authService := auth.NewAuthService(userService, jwt)

	httpServer := api.NewHTTPServer(
		config.ServerAddress(),
		authService,
		userService,
		roomsService,
		jwt,
		log,
	)

	log.Printf("HTTP Server listening at: %v", config.ServerAddress().String())

	if err := httpServer.Start(); err != nil {
		panic(err)
	}

}
