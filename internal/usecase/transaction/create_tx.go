package transactionusecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	transactiondomain "github.com/lehoangvuvt/go-ent-boilerplate/internal/domain/transaction"
	repositoryports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/repository"
	transactionusecasedto "github.com/lehoangvuvt/go-ent-boilerplate/internal/usecase/transaction/dto"
)

type CreateTransactionUsecase struct {
	repo repositoryports.TransactionRepository
}

func NewCreateTransactionUsecase(repo repositoryports.TransactionRepository) *CreateTransactionUsecase {
	return &CreateTransactionUsecase{repo: repo}
}

func (uc *CreateTransactionUsecase) Execute(ctx context.Context, userID uuid.UUID, req transactionusecasedto.CreateTransactionRequest) (*transactiondomain.Transaction, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	tx := &transactiondomain.Transaction{
		ID:        uuid.New(),
		Amount:    req.Amount,
		Currency:  req.Currency,
		UserID:    userID,
		Method:    req.Method,
		Status:    transactiondomain.Pending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Visa:      req.Visa,
		Banking:   req.Banking,
		EWallet:   req.EWallet,
		QRPay:     req.QRPay,
	}

	return uc.repo.Create(ctx, tx)
}
