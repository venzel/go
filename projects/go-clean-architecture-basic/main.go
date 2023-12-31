package main

import (
	"net/http"
	"server/internal/domain/user"
	"server/internal/handlers"
	"server/internal/infra/database"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	route := chi.NewRouter()

	route.Use(middleware.Logger)

	userRepository := database.UserRepository{}

	userService := user.ServiceImpl{
		Repository: &userRepository,
	}

	userHandler := handlers.Handler{
		UserService: &userService,
	}

	route.Post("/users", handlers.HandlerError(userHandler.CreateUser))

	port := ":3000"

	http.ListenAndServe(port, route)
}
