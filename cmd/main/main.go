package main

import (
	"context"
	userRepository "github.com/IvanStukalov/TimeHack-Backend/internal/pkg/user/repo"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

func main() {
	connString := "postgres://postgres:12345@localhost:5432/timehack"
	pool, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		log.Fatal("Connection to postgres failed", err)
	}
	defer pool.Close()

	userRepo := userRepository.NewRepo(pool)
	userRepo.GetUsers(context.Background())
}
