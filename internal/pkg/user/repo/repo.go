package repo

import (
	"context"
	"fmt"
	"github.com/IvanStukalov/TimeHack-Backend/internal/pkg/user"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type repoPostgres struct {
	Conn *pgxpool.Pool
}

func NewRepo(conn *pgxpool.Pool) user.Repository {
	return &repoPostgres{Conn: conn}
}

func (r *repoPostgres) GetAllUsers(ctx context.Context) {
	sql := `SELECT * FROM "user"`
	rows, err := r.Conn.Query(ctx, sql)
	if err != nil {
		log.Fatal(err)
	}

	var id int
	var login string
	var name string
	var pass string
	for rows.Next() {
		err = rows.Scan(&id, &login, &name, &pass)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, login, name, pass)
	}
}
