package main

import (
	"context"
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

	userRepo :=
}
