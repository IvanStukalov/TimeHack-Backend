package repo

import (
	"context"
	"github.com/IvanStukalov/TimeHack-Backend/internal/models"
	"github.com/IvanStukalov/TimeHack-Backend/internal/pkg/user"
	"github.com/jackc/pgx/v4/pgxpool"
)

type repoPostgres struct {
	Conn *pgxpool.Pool
}

func NewRepo(conn *pgxpool.Pool) user.Repository {
	return &repoPostgres{Conn: conn}
}

func (r *repoPostgres) GetAllUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User

	sql := `SELECT * FROM "user"`
	rows, err := r.Conn.Query(ctx, sql)
	if err != nil {
		return users, err
	}

	for i := 0; rows.Next(); i++ {
		user := models.User{}
		err = rows.Scan(&user.ID, &user.Login, &user.Username, &user.PasswordHash)
		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}
