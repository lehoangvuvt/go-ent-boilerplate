package transactiondomain

import (
	"time"

	"github.com/google/uuid"
	transactionmethoddomain "github.com/lehoangvuvt/go-ent-boilerplate/internal/domain/transaction/method"
)

func NewVisaTransaction(
	amount int64,
	currency string,
	userID uuid.UUID,
	visa *transactionmethoddomain.VisaDetails,
) *Transaction {
	return &Transaction{
		ID:        uuid.New(),
		Amount:    amount,
		Currency:  currency,
		UserID:    userID,
		Method:    transactionmethoddomain.MethodVisa,
		Visa:      visa,
		Status:    Pending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func NewBankingTransaction(
	amount int64,
	currency string,
	userID uuid.UUID,
	bank *transactionmethoddomain.BankingDetails,
) *Transaction {
	return &Transaction{
		ID:        uuid.New(),
		Amount:    amount,
		Currency:  currency,
		UserID:    userID,
		Method:    transactionmethoddomain.MethodBanking,
		Banking:   bank,
		Status:    Pending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
