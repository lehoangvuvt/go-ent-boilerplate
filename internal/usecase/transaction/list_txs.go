package transactionusecase

import (
	"context"

	transactiondomain "github.com/lehoangvuvt/go-ent-boilerplate/internal/domain/transaction"
	repositoryports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/repository"
)

type ListTransactionsUsecase struct {
	repo repositoryports.TransactionRepository
}

func NewListTransactionsUsecase(repo repositoryports.TransactionRepository) *ListTransactionsUsecase {
	return &ListTransactionsUsecase{repo: repo}
}

func (uc *ListTransactionsUsecase) Execute(ctx context.Context, filter repositoryports.TransactionFilter) ([]*transactiondomain.Transaction, error) {
	return uc.repo.List(ctx, filter)
}
