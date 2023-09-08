package user

import "context"

type Repository interface {
	GetAllUsers(ctx context.Context)
}

type UseCase interface {
}
