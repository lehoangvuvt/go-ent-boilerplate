package httpauth

import (
	"net/http"

	authusecase "github.com/lehoangvuvt/go-ent-boilerplate/internal/usecase/auth"
	authusecasedto "github.com/lehoangvuvt/go-ent-boilerplate/internal/usecase/auth/dto"
	"github.com/lehoangvuvt/go-ent-boilerplate/pkg/httpx"
)

type NewAuthHandlerArgs struct {
	LoginUC *authusecase.LoginUsecase
}

type AuthHandler struct {
	loginUC *authusecase.LoginUsecase
}

func NewAuthHandler(args NewAuthHandlerArgs) *AuthHandler {
	return &AuthHandler{
		loginUC: args.LoginUC,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	req := new(authusecasedto.LoginRequest)

	if err := httpx.FromJSON(r, req); err != nil {
		httpx.ToJSON(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		return
	}

	resp, err := h.loginUC.Execute(r.Context(), req)
	if err != nil {
		httpx.ToJSON(w, map[string]string{"error": err.Error()}, http.StatusUnauthorized)
		return
	}

	httpx.ToJSON(w, resp, http.StatusOK)
}
