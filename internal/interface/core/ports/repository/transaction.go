package repositoryports

import (
	"context"

	"github.com/google/uuid"
	transactiondomain "github.com/lehoangvuvt/go-ent-boilerplate/internal/domain/transaction"
	transactionmethoddomain "github.com/lehoangvuvt/go-ent-boilerplate/internal/domain/transaction/method"
)

type TransactionFilter struct {
	UserID *uuid.UUID
	Status *transactiondomain.TransactionStatus
	Method *transactionmethoddomain.PaymentMethodType

	Limit  int
	Offset int
}

type TransactionRepository interface {
	Create(ctx context.Context, t *transactiondomain.Transaction) (*transactiondomain.Transaction, error)
	Update(ctx context.Context, t *transactiondomain.Transaction) (*transactiondomain.Transaction, error)
	FindByID(ctx context.Context, id uuid.UUID) (*transactiondomain.Transaction, error)
	FindByUser(ctx context.Context, userID uuid.UUID) ([]*transactiondomain.Transaction, error)
	FindByStatus(ctx context.Context, status transactiondomain.TransactionStatus) ([]*transactiondomain.Transaction, error)
	FindByReferenceID(ctx context.Context, refID string) (*transactiondomain.Transaction, error)
	List(ctx context.Context, filter TransactionFilter) ([]*transactiondomain.Transaction, error)
}
