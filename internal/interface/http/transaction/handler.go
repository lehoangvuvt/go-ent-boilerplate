package httptransaction

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	httpmiddleware "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/http/middleware"
	transactionusecase "github.com/lehoangvuvt/go-ent-boilerplate/internal/usecase/transaction"
	transactionusecasedto "github.com/lehoangvuvt/go-ent-boilerplate/internal/usecase/transaction/dto"
	"github.com/lehoangvuvt/go-ent-boilerplate/pkg/httpx"
)

type NewTransactionHandlerArgs struct {
	CreateUC  *transactionusecase.CreateTransactionUsecase
	ConfirmUC *transactionusecase.ConfirmTransactionUsecase
	CancelUC  *transactionusecase.CancelTransactionUsecase
	FailUC    *transactionusecase.FailTransactionUsecase
	FindUC    *transactionusecase.FindTransactionByIDUsecase
	ListUC    *transactionusecase.ListTransactionsUsecase
}

type TransactionHandler struct {
	createUC  *transactionusecase.CreateTransactionUsecase
	confirmUC *transactionusecase.ConfirmTransactionUsecase
	cancelUC  *transactionusecase.CancelTransactionUsecase
	failUC    *transactionusecase.FailTransactionUsecase
	findUC    *transactionusecase.FindTransactionByIDUsecase
	listUC    *transactionusecase.ListTransactionsUsecase
}

func NewTransactionHandler(args NewTransactionHandlerArgs) *TransactionHandler {
	return &TransactionHandler{
		createUC:  args.CreateUC,
		confirmUC: args.ConfirmUC,
		cancelUC:  args.CancelUC,
		failUC:    args.FailUC,
		findUC:    args.FindUC,
		listUC:    args.ListUC,
	}
}

func (h *TransactionHandler) Create(w http.ResponseWriter, r *http.Request) {
	req := new(transactionusecasedto.CreateTransactionRequest)

	if err := httpx.FromJSON(r, req); err != nil {
		httpx.ToJSON(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		httpx.ToJSON(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		return
	}
	authClaims, ok := httpmiddleware.FromContext(r.Context())
	if !ok {
		httpx.ToJSON(w, map[string]string{"error": "Unauthorized"}, http.StatusUnauthorized)
		return
	}
	userID, err := uuid.Parse(authClaims.Subject)
	if err != nil {
		httpx.ToJSON(w, map[string]string{"error": "Unauthorized"}, http.StatusUnauthorized)
		return
	}
	tx, err := h.createUC.Execute(r.Context(), userID, *req)
	if err != nil {
		httpx.ToJSON(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	httpx.ToJSON(w, tx, http.StatusCreated)
}

func (h *TransactionHandler) Confirm(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	txID, err := uuid.Parse(id)
	if err != nil {
		httpx.ToJSON(w, map[string]string{"error": "invalid transaction ID"}, http.StatusBadRequest)
		return
	}

	tx, err := h.confirmUC.Execute(r.Context(), txID)
	if err != nil {
		httpx.ToJSON(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		return
	}

	httpx.ToJSON(w, tx, http.StatusOK)
}

func (h *TransactionHandler) Cancel(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	txID, err := uuid.Parse(id)
	if err != nil {
		httpx.ToJSON(w, map[string]string{"error": "invalid transaction ID"}, http.StatusBadRequest)
		return
	}

	tx, err := h.cancelUC.Execute(r.Context(), txID)
	if err != nil {
		httpx.ToJSON(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		return
	}

	httpx.ToJSON(w, tx, http.StatusOK)
}

func (h *TransactionHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	txID, err := uuid.Parse(id)
	if err != nil {
		httpx.ToJSON(w, map[string]string{"error": "invalid transaction ID"}, http.StatusBadRequest)
		return
	}

	tx, err := h.findUC.Execute(r.Context(), txID)
	if err != nil {
		httpx.ToJSON(w, map[string]string{"error": err.Error()}, http.StatusNotFound)
		return
	}

	httpx.ToJSON(w, tx, http.StatusOK)
}

func (h *TransactionHandler) List(w http.ResponseWriter, r *http.Request) {
	filter, err := parseTransactionFilter(r)
	if err != nil {
		httpx.ToJSON(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		return
	}

	items, err := h.listUC.Execute(r.Context(), filter)
	if err != nil {
		httpx.ToJSON(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	httpx.ToJSON(w, items, http.StatusOK)
}
