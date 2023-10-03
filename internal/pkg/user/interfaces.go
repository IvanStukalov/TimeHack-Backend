package user

import (
	"context"
	"github.com/IvanStukalov/TimeHack-Backend/internal/models"
)

type Repository interface {
	GetAllUsers(ctx context.Context) ([]models.User, error)
}

type UseCase interface {
	GetAllUsers(ctx context.Context) ([]models.User, error)
}
