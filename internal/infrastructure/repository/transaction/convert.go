package transactionrepository

import (
	"github.com/lehoangvuvt/go-ent-boilerplate/ent"
	"github.com/lehoangvuvt/go-ent-boilerplate/ent/transaction"
	transactiondomain "github.com/lehoangvuvt/go-ent-boilerplate/internal/domain/transaction"
	transactionmethoddomain "github.com/lehoangvuvt/go-ent-boilerplate/internal/domain/transaction/method"
)

func toDomain(et *ent.Transaction) (*transactiondomain.Transaction, error) {
	dt := &transactiondomain.Transaction{
		ID:        et.ID,
		Amount:    et.Amount,
		Currency:  et.Currency,
		UserID:    et.UserID,
		Method:    transactionmethoddomain.PaymentMethodType(et.Method),
		Status:    transactiondomain.TransactionStatus(et.Status),
		CreatedAt: et.CreatedAt,
		UpdatedAt: et.UpdatedAt,
	}

	if et.VisaDetails != nil {
		v := *et.VisaDetails
		dt.Visa = &v
	}

	if et.BankingDetails != nil {
		b := *et.BankingDetails
		dt.Banking = &b
	}

	if et.EwalletDetails != nil {
		e := *et.EwalletDetails
		dt.EWallet = &e
	}

	if et.QrDetails != nil {
		q := *et.QrDetails
		dt.QRPay = &q
	}

	return dt, nil
}

func toDomainList(list []*ent.Transaction) ([]*transactiondomain.Transaction, error) {
	out := make([]*transactiondomain.Transaction, 0, len(list))
	for _, et := range list {
		dt, err := toDomain(et)
		if err != nil {
			return nil, err
		}
		out = append(out, dt)
	}
	return out, nil
}

func applyDomainToCreate(b *ent.TransactionCreate, dt *transactiondomain.Transaction) {
	b.
		SetID(dt.ID).
		SetAmount(dt.Amount).
		SetCurrency(dt.Currency).
		SetUserID(dt.UserID).
		SetMethod(transaction.Method(dt.Method)).
		SetStatus(transaction.Status(dt.Status)).
		SetCreatedAt(dt.CreatedAt).
		SetUpdatedAt(dt.UpdatedAt)

	if dt.Visa != nil {
		b.SetVisaDetails(dt.Visa)
	}

	if dt.Banking != nil {
		b.SetBankingDetails(dt.Banking)
	}

	if dt.EWallet != nil {
		b.SetEwalletDetails(dt.EWallet)
	}

	if dt.QRPay != nil {
		b.SetQrDetails(dt.QRPay)
	}
}

func applyDomainToUpdate(b *ent.TransactionUpdateOne, dt *transactiondomain.Transaction) {
	b.
		SetAmount(dt.Amount).
		SetCurrency(dt.Currency).
		SetUserID(dt.UserID).
		SetMethod(transaction.Method(dt.Method)).
		SetStatus(transaction.Status(dt.Status)).
		SetUpdatedAt(dt.UpdatedAt)

	if dt.Visa != nil {
		b.SetVisaDetails(dt.Visa)
	} else {
		b.ClearVisaDetails()
	}

	if dt.Banking != nil {
		b.SetBankingDetails(dt.Banking)
	} else {
		b.ClearBankingDetails()
	}

	if dt.EWallet != nil {
		b.SetEwalletDetails(dt.EWallet)
	} else {
		b.ClearEwalletDetails()
	}

	if dt.QRPay != nil {
		b.SetQrDetails(dt.QRPay)
	} else {
		b.ClearQrDetails()
	}
}
