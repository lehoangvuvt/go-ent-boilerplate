package httpuser

import (
	"net/http"

	userusecase "github.com/lehoangvuvt/go-ent-boilerplate/internal/usecase/user"
	userusercasedto "github.com/lehoangvuvt/go-ent-boilerplate/internal/usecase/user/dto"
	"github.com/lehoangvuvt/go-ent-boilerplate/pkg/httpx"
)

type NewUserHandlerArgs struct {
	CreateUserUC *userusecase.CreateUserUsecase
}

type UserHandler struct {
	createUserUC *userusecase.CreateUserUsecase
}

func NewUserHandler(args NewUserHandlerArgs) *UserHandler {
	return &UserHandler{
		createUserUC: args.CreateUserUC,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	req := new(userusercasedto.CreateUserRequest)

	if err := httpx.FromJSON(r, req); err != nil {
		httpx.ToJSON(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		return
	}

	user, err := h.createUserUC.Execute(r.Context(), req)
	if err != nil {
		httpx.ToJSON(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		return
	}

	httpx.ToJSON(w, user, http.StatusCreated)
}
