package transactiondomain

import (
	"time"

	"github.com/google/uuid"
	transactionmethoddomain "github.com/lehoangvuvt/go-ent-boilerplate/internal/domain/transaction/method"
)

type TransactionStatus string

const (
	Pending   TransactionStatus = "pending"
	Completed TransactionStatus = "completed"
	Failed    TransactionStatus = "failed"
	Rejected  TransactionStatus = "rejected"
)

type Transaction struct {
	ID       uuid.UUID
	Amount   int64
	Currency string
	UserID   uuid.UUID

	Method transactionmethoddomain.PaymentMethodType

	Visa    *transactionmethoddomain.VisaDetails
	Banking *transactionmethoddomain.BankingDetails
	EWallet *transactionmethoddomain.EWalletDetails
	QRPay   *transactionmethoddomain.QRDetails

	Status    TransactionStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}
