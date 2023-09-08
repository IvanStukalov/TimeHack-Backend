package delivery

import (
	"context"
	"github.com/IvanStukalov/TimeHack-Backend/internal/pkg/user"
	"net/http"
)

type Handler struct {
	uc user.UseCase
}

func NewHandler(uc user.UseCase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	h.uc.GetAllUsers(context.Background())
}
