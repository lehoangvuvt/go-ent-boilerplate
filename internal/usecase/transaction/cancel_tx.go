package transactionusecase

import (
	"context"

	"github.com/google/uuid"
	transactiondomain "github.com/lehoangvuvt/go-ent-boilerplate/internal/domain/transaction"
	repositoryports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/repository"
)

type CancelTransactionUsecase struct {
	repo repositoryports.TransactionRepository
}

func NewCancelTransactionUsecase(repo repositoryports.TransactionRepository) *CancelTransactionUsecase {
	return &CancelTransactionUsecase{repo: repo}
}

func (uc *CancelTransactionUsecase) Execute(ctx context.Context, id uuid.UUID) (*transactiondomain.Transaction, error) {
	tx, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := tx.Cancel(); err != nil {
		return nil, err
	}

	return uc.repo.Update(ctx, tx)
}
