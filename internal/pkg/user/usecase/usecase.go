package usecase

import (
	"context"
	"github.com/IvanStukalov/TimeHack-Backend/internal/pkg/user"
)

type UseCase struct {
	repo user.Repository
}

func NewUseCase(repo user.Repository) user.UseCase {
	return &UseCase{repo: repo}
}

func (u *UseCase) GetAllUsers(ctx context.Context) {
	u.repo.GetAllUsers(ctx)
}
