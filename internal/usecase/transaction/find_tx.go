package transactionusecase

import (
	"context"

	"github.com/google/uuid"
	transactiondomain "github.com/lehoangvuvt/go-ent-boilerplate/internal/domain/transaction"
	repositoryports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/repository"
)

type FindTransactionByIDUsecase struct {
	repo repositoryports.TransactionRepository
}

func NewFindTransactionByIDUsecase(repo repositoryports.TransactionRepository) *FindTransactionByIDUsecase {
	return &FindTransactionByIDUsecase{repo: repo}
}

func (uc *FindTransactionByIDUsecase) Execute(ctx context.Context, id uuid.UUID) (*transactiondomain.Transaction, error) {
	return uc.repo.FindByID(ctx, id)
}
