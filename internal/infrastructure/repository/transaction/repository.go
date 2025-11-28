package transactionrepository

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/lehoangvuvt/go-ent-boilerplate/ent"
	"github.com/lehoangvuvt/go-ent-boilerplate/ent/transaction"
	transactiondomain "github.com/lehoangvuvt/go-ent-boilerplate/internal/domain/transaction"
	repositoryports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/repository"
)

type TransactionRepository struct {
	client *ent.Client
}

var _ repositoryports.TransactionRepository = (*TransactionRepository)(nil)

func NewTransactionRepository(client *ent.Client) *TransactionRepository {
	return &TransactionRepository{
		client: client,
	}
}

func (r *TransactionRepository) Create(ctx context.Context, dt *transactiondomain.Transaction) (*transactiondomain.Transaction, error) {
	builder := r.client.Transaction.Create()
	applyDomainToCreate(builder, dt)

	et, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return toDomain(et)
}

func (r *TransactionRepository) Update(ctx context.Context, dt *transactiondomain.Transaction) (*transactiondomain.Transaction, error) {
	builder := r.client.Transaction.UpdateOneID(dt.ID)
	applyDomainToUpdate(builder, dt)

	et, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return toDomain(et)
}

func (r *TransactionRepository) FindByID(ctx context.Context, id uuid.UUID) (*transactiondomain.Transaction, error) {
	et, err := r.client.Transaction.Get(ctx, id)
	if ent.IsNotFound(err) {
		return nil, repositoryports.ErrTransactionNotFound
	}
	if err != nil {
		return nil, err
	}
	return toDomain(et)
}

func (r *TransactionRepository) FindByUser(ctx context.Context, userID uuid.UUID) ([]*transactiondomain.Transaction, error) {
	ets, err := r.client.Transaction.
		Query().
		Where(transaction.UserIDEQ(userID)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return toDomainList(ets)
}

func (r *TransactionRepository) FindByStatus(ctx context.Context, status transactiondomain.TransactionStatus) ([]*transactiondomain.Transaction, error) {
	ets, err := r.client.Transaction.
		Query().
		Where(transaction.StatusEQ(transaction.Status(status))).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return toDomainList(ets)
}

func (r *TransactionRepository) FindByReferenceID(ctx context.Context, refID string) (*transactiondomain.Transaction, error) {
	predicate := func(s *sql.Selector) {
		t := s.Table() // current table name

		s.Where(sql.Or(
			sql.ExprP(fmt.Sprintf("%s.visa_details::text LIKE ?", t), "%"+refID+"%"),
			sql.ExprP(fmt.Sprintf("%s.banking_details::text LIKE ?", t), "%"+refID+"%"),
			sql.ExprP(fmt.Sprintf("%s.ewallet_details::text LIKE ?", t), "%"+refID+"%"),
			sql.ExprP(fmt.Sprintf("%s.qr_details::text LIKE ?", t), "%"+refID+"%"),
		))
	}

	et, err := r.client.Transaction.Query().
		Where(predicate).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, repositoryports.ErrTransactionNotFound
	}
	if err != nil {
		return nil, err
	}

	return toDomain(et)
}

func (r *TransactionRepository) List(ctx context.Context, f repositoryports.TransactionFilter) ([]*transactiondomain.Transaction, error) {
	q := r.client.Transaction.Query()

	if f.UserID != nil {
		q = q.Where(transaction.UserIDEQ(*f.UserID))
	}
	if f.Status != nil {
		q = q.Where(transaction.StatusEQ(transaction.Status(*f.Status)))
	}
	if f.Method != nil {
		q = q.Where(transaction.MethodEQ(transaction.Method(*f.Method)))
	}

	if f.Limit > 0 {
		q = q.Limit(f.Limit)
	}
	if f.Offset > 0 {
		q = q.Offset(f.Offset)
	}

	ets, err := q.All(ctx)
	if err != nil {
		return nil, err
	}

	return toDomainList(ets)
}
