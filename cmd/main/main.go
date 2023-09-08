package main

import (
	"context"
	userHandler "github.com/IvanStukalov/TimeHack-Backend/internal/pkg/user/delivery"
	userRepository "github.com/IvanStukalov/TimeHack-Backend/internal/pkg/user/repo"
	userUseCase "github.com/IvanStukalov/TimeHack-Backend/internal/pkg/user/usecase"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

func main() {
	muxRoute := mux.NewRouter()
	connString := "postgres://postgres:12345@localhost:5432/timehack"
	pool, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		log.Fatal("Connection to postgres failed", err)
	}
	defer pool.Close()

	userRepo := userRepository.NewRepo(pool)
	userUC := userUseCase.NewUseCase(userRepo)
	userH := userHandler.NewHandler(userUC)

	api := muxRoute.PathPrefix("/api").Subrouter()
	{
		user := api.PathPrefix("/user").Subrouter()
		{
			user.HandleFunc("/")
		}
	}

}
