package transactionusecase

import (
	"context"

	"github.com/google/uuid"
	transactiondomain "github.com/lehoangvuvt/go-ent-boilerplate/internal/domain/transaction"
	repositoryports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/repository"
)

type ConfirmTransactionUsecase struct {
	repo repositoryports.TransactionRepository
}

func NewConfirmTransactionUsecase(repo repositoryports.TransactionRepository) *ConfirmTransactionUsecase {
	return &ConfirmTransactionUsecase{repo: repo}
}

func (uc *ConfirmTransactionUsecase) Execute(ctx context.Context, id uuid.UUID) (*transactiondomain.Transaction, error) {
	tx, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := tx.Confirm(); err != nil {
		return nil, err
	}

	return uc.repo.Update(ctx, tx)
}
