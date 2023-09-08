package delivery

import (
	"github.com/IvanStukalov/TimeHack-Backend/internal/pkg/user"
)

type Handler struct {
	useCase user.UseCase
}

func NewHandler(uc user.UseCase) *Handler {
	return &Handler{useCase: uc}
}
