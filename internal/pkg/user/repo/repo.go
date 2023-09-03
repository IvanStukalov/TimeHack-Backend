package repo

import (
	"github.com/IvanStukalov/TimeHack-Backend/internal/pkg/user"
	"github.com/jackc/pgx/v4/pgxpool"
)

type repoPostgres struct {
	Conn *pgxpool.Pool
}

func NewRepoPostgres(conn *pgxpool.Pool) user.Repository {
	return &repoPostgres{Conn: conn}
}
