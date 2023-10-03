package delivery

import (
	"github.com/IvanStukalov/TimeHack-Backend/internal/pkg/user"
	"github.com/IvanStukalov/TimeHack-Backend/internal/utils"
	"net/http"
)

type Handler struct {
	uc user.UseCase
}

func NewUserHandler(uc user.UseCase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.uc.GetAllUsers(r.Context())
	if err != nil {
		utils.Response(w, http.StatusNotFound, nil)
	}

	utils.Response(w, http.StatusOK, users)
}
